package services

import (
	"context"
	"time"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/pkg/broker"
	"github.com/lessbutter/alloff-api/pkg/hometab"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/golang"
)

type HomeTabService struct {
	grpcServer.HomeTabItemServer
}

func (s *HomeTabService) GetHomeTabItem(ctx context.Context, req *grpcServer.GetHomeTabItemRequest) (*grpcServer.GetHomeTabItemResponse, error) {
	itemDao, err := ioc.Repo.HomeTabItems.Get(req.ItemId)
	if err != nil {
		return nil, err
	}

	return &grpcServer.GetHomeTabItemResponse{
		Item: mapper.HomeTabItemMapper(itemDao),
	}, nil
}

func (s *HomeTabService) ListHomeTabItems(ctx context.Context, req *grpcServer.ListHomeTabItemsRequest) (*grpcServer.ListHomeTabItemsResponse, error) {
	onlyLive := false
	itemDaos, cnt, err := ioc.Repo.HomeTabItems.List(int(req.Offset), int(req.Limit), onlyLive)
	if err != nil {
		return nil, err
	}

	items := []*grpcServer.HomeTabItemMessage{}
	for _, itemDao := range itemDaos {
		items = append(items, mapper.HomeTabItemMapper(itemDao))
	}

	return &grpcServer.ListHomeTabItemsResponse{
		Offset:      req.Offset,
		Limit:       req.Limit,
		TotalCounts: int32(cnt),
		Items:       items,
	}, nil
}

func (s *HomeTabService) EditHomeTabItem(ctx context.Context, req *grpcServer.EditHomeTabItemRequest) (*grpcServer.EditHomeTabItemResponse, error) {
	itemDao, err := ioc.Repo.HomeTabItems.Get(req.HometabId)
	if err != nil {
		return nil, err
	}

	layout := "2006-01-02T15:04:05Z07:00"

	if req.Title != nil {
		itemDao.Title = *req.Title
	}
	if req.Description != nil {
		itemDao.Description = *req.Description
	}
	if req.Tags != nil {
		itemDao.Tags = req.Tags
	}
	if req.BackImageUrl != nil {
		itemDao.BackImageUrl = *req.BackImageUrl
	}
	if req.StartTime != nil {
		startTimeObj, _ := time.Parse(layout, *req.StartTime)
		itemDao.StartedAt = startTimeObj
	}
	if req.FinishTime != nil {
		finishTimeObj, _ := time.Parse(layout, *req.FinishTime)
		itemDao.FinishedAt = finishTimeObj
	}
	if req.Weight != nil {
		itemDao.Weight = int(*req.Weight)
	}
	if req.IsLive != nil {
		itemDao.IsLive = *req.IsLive
	}

	// TO BE SPECIFIED LATER
	// Contents

	newItemDao, err := ioc.Repo.HomeTabItems.Update(itemDao)
	if err != nil {
		return nil, err
	}

	go broker.HomeTabSyncer()

	return &grpcServer.EditHomeTabItemResponse{
		Item: mapper.HomeTabItemMapper(newItemDao),
	}, nil
}

func (s *HomeTabService) CreateHomeTabItem(ctx context.Context, req *grpcServer.CreateHomeTabItemRequest) (*grpcServer.CreateHomeTabItemResponse, error) {
	layout := "2006-01-02T15:04:05Z07:00"

	startTimeObj, _ := time.Parse(layout, req.StartTime)
	finishTimeObj, _ := time.Parse(layout, req.FinishTime)

	backImageUrl := ""
	if req.BackImageUrl != nil {
		backImageUrl = *req.BackImageUrl
	}
	addItemRequest := &hometab.HomeTabItemRequest{
		Title:        req.Title,
		Description:  req.Description,
		Tags:         req.Tags,
		BackImageUrl: backImageUrl,
		StartedAt:    startTimeObj,
		FinishedAt:   finishTimeObj,
	}

	itemType := req.Contents.GetItemType()
	switch itemType {
	case grpcServer.ItemType_HOMETAB_ITEM_BRANDS:
		addItemRequest.Requester = &hometab.BrandsItemRequest{
			BrandKeynames: req.Contents.BrandKeynames,
		}
	case grpcServer.ItemType_HOMETAB_ITEM_EXHIBITION_A:
		if len(req.Contents.ExhibitionIds) == 0 {
			break
		}
		addItemRequest.Requester = &hometab.BrandExhibitionItemRequest{
			ExhibitionID: req.Contents.ExhibitionIds[0],
			ProductIDs:   req.Contents.ProductIds,
		}
	case grpcServer.ItemType_HOMETAB_ITEM_EXHIBITION:
		if len(req.Contents.ExhibitionIds) == 0 {
			break
		}
		addItemRequest.Requester = &hometab.ExhibitionItemRequest{
			ExhibitionID: req.Contents.ExhibitionIds[0],
			ProductIDs:   req.Contents.ProductIds,
		}
	case grpcServer.ItemType_HOMETAB_ITEM_EXHIBITIONS:
		addItemRequest.Requester = &hometab.ExhibitionsItemRequest{
			ExhibitionIDs: req.Contents.ExhibitionIds,
		}
	case grpcServer.ItemType_HOMETAB_ITEM_PRODUCTS_CATEGORIES:
		sortingOptions := []model.SortingType{}
		for _, option := range req.Contents.Options {
			sortingOptions = append(sortingOptions, model.SortingType(option.Descriptor().FullName()))
		}
		addItemRequest.Requester = &hometab.AlloffCategoryItemRequest{
			AlloffCategoryID: *req.Contents.AlloffcategoryId,
			SortingOptions:   sortingOptions, // TO BE SPECIFIED
		}
	case grpcServer.ItemType_HOMETAB_ITEM_PRODUCTS_BRANDS:
		sortingOptions := []model.SortingType{}
		for _, option := range req.Contents.Options {
			sortingOptions = append(sortingOptions, model.SortingType(option.Descriptor().FullName()))
		}
		addItemRequest.Requester = &hometab.BrandProductsItemRequest{
			BrandKeyname:   req.Contents.BrandKeynames[0],
			SortingOptions: sortingOptions,
		}
	}

	newTabDao, err := hometab.AddHometabItem(addItemRequest)
	if err != nil {
		return nil, err
	}

	return &grpcServer.CreateHomeTabItemResponse{
		Item: mapper.HomeTabItemMapper(newTabDao),
	}, nil
}
