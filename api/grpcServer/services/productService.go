package services

import (
	"context"
	"errors"

	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
	"go.uber.org/zap"
)

type ProductService struct {
	grpcServer.ProductServer
}

func (s *ProductService) GetProduct(ctx context.Context, req *grpcServer.GetProductRequest) (*grpcServer.GetProductResponse, error) {
	pdInfoDao, err := ioc.Repo.ProductMetaInfos.Get(req.AlloffProductId)
	if err != nil {
		return nil, err
	}

	return &grpcServer.GetProductResponse{
		Product: mapper.ProductInfoMapper(pdInfoDao),
	}, nil
}

func (s *ProductService) ListProducts(ctx context.Context, req *grpcServer.ListProductsRequest) (*grpcServer.ListProductsResponse, error) {
	moduleName := ""
	if req.ModuleName != nil {
		moduleName = *req.ModuleName
	}

	alloffCategoryID := ""
	if req.Query.AlloffCategoryId != nil {
		alloffCategoryID = *req.Query.AlloffCategoryId
	}

	searchKeyword := ""
	if req.Query.SearchQuery != nil {
		searchKeyword = *req.Query.SearchQuery
	}

	classifiedType := domain.NO_MATTER_CLASSIFIED
	if req.Query.IsClassifiedDone != nil {
		if *req.Query.IsClassifiedDone {
			classifiedType = domain.CLASSIFIED_DONE
		} else {
			classifiedType = domain.NOT_CLASSIFIED
		}
	}

	var priceSorting domain.PriceSortingType
	priceRanges := []domain.PriceRangeType{}
	if req.Query.Options != nil {
		priceRanges, priceSorting = mapper.ProductSortingAndRangesMapper(req.Query.Options)
	}

	query := productinfo.ProductInfoListInput{
		Offset:                int(req.Offset),
		Limit:                 int(req.Limit),
		BrandID:               "",
		AlloffCategoryID:      alloffCategoryID,
		Keyword:               searchKeyword,
		Modulename:            moduleName,
		IncludeClassifiedType: classifiedType,
		PriceRanges:           priceRanges,
		PriceSorting:          priceSorting,
	}

	products, cnt, err := productinfo.ListProductInfos(query)
	if err != nil {
		return nil, err
	}

	pds := []*grpcServer.ProductMessage{}

	for _, pd := range products {
		pds = append(pds, mapper.ProductInfoMapper(pd))
	}

	ret := &grpcServer.ListProductsResponse{
		Offset:      req.Offset,
		Limit:       req.Limit,
		TotalCounts: int32(cnt),
		Products:    pds,
		ListQuery:   req.Query,
	}

	return ret, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, req *grpcServer.CreateProductRequest) (*grpcServer.CreateProductResponse, error) {
	discountedPrice := int(req.DiscountedPrice)
	originalPrice := discountedPrice
	if req.OriginalPrice != nil {
		originalPrice = int(*req.OriginalPrice)
	}

	moduleName := "manual"
	if req.ModuleName != nil {
		moduleName = *req.ModuleName
	}

	// TODO: ProductTypes를 바꾸는 작업이 되어야한다.
	productTypes := []domain.AlloffProductType{
		domain.Female,
		domain.Female,
		domain.Kids,
		domain.Sports,
	}
	for _, reqPdType := range req.ProductTypes {
		if reqPdType == grpcServer.ProductType_FEMALE {
			productTypes = append(productTypes, domain.Female)
		} else if reqPdType == grpcServer.ProductType_MALE {
			productTypes = append(productTypes, domain.Male)
		} else if reqPdType == grpcServer.ProductType_KIDS {
			productTypes = append(productTypes, domain.Kids)
		}
	}

	productID := ""
	if req.ProductId != nil {
		productID = *req.ProductId
	} else {
		productID = utils.CreateShortUUID()
	}

	productUrl := ""
	if req.ProductUrl != nil {
		productUrl = *req.ProductUrl
	}

	var descInfos, pdInfos map[string]string
	if req.DescriptionInfos != nil {
		descInfos = req.DescriptionInfos
	}
	if req.ProductInfos != nil {
		pdInfos = req.ProductInfos
	}
	thumbnailImage := ""
	if req.ThumbnailImage != nil && *req.ThumbnailImage != "" {
		thumbnailImage = *req.ThumbnailImage
	}

	brand, _ := ioc.Repo.Brands.GetByKeyname(req.BrandKeyName)
	alloffcat, _ := ioc.Repo.AlloffCategories.Get(*req.AlloffCategoryId)
	invDaos := []*domain.InventoryDAO{}
	for _, inv := range req.Inventory {
		invDaos = append(invDaos, &domain.InventoryDAO{
			Quantity: int(inv.Quantity),
			Size:     inv.Size,
		})
	}

	addRequest := &productinfo.AddMetaInfoRequest{
		AlloffName:           req.AlloffName,
		ProductID:            productID,
		ProductUrl:           productUrl,
		ProductType:          productTypes,
		OriginalPrice:        float32(originalPrice),
		DiscountedPrice:      float32(discountedPrice),
		CurrencyType:         domain.CurrencyKRW,
		Brand:                brand,
		Source:               &domain.CrawlSourceDAO{CrawlModuleName: "manual"},
		AlloffCategory:       alloffcat,
		Images:               req.Images,
		ThumbnailImage:       thumbnailImage,
		Colors:               []string{},
		Sizes:                []string{},
		Description:          req.Description,
		DescriptionImages:    req.DescriptionImages,
		DescriptionInfos:     descInfos,
		Information:          pdInfos,
		IsForeignDelivery:    req.IsForeignDelivery,
		EarliestDeliveryDays: int(req.EarliestDeliveryDays),
		LatestDeliveryDays:   int(req.LatestDeliveryDays),
		IsRefundPossible:     req.IsRefundPossible,
		RefundFee:            int(req.RefundFee),
		Inventory:            invDaos,
		ModuleName:           moduleName,
		IsInventoryMapped:    true,
	}

	pdInfoDao, err := productinfo.AddProductInfo(addRequest)
	if err != nil {
		return nil, err
	}

	pdMessage := mapper.ProductInfoMapper(pdInfoDao)

	return &grpcServer.CreateProductResponse{
		Product: pdMessage,
	}, nil
}

func (s *ProductService) EditProduct(ctx context.Context, req *grpcServer.EditProductRequest) (*grpcServer.EditProductResponse, error) {
	pdInfoDao, err := ioc.Repo.ProductMetaInfos.Get(req.AlloffProductId)
	if err != nil {
		return nil, err
	}

	updatedRequest := productinfo.LoadMetaInfoRequest(pdInfoDao)

	if req.ModuleName != "" && req.ModuleName != "manual" {
		if pdInfoDao.Source.CrawlModuleName != req.ModuleName {
			return nil, errors.New("not authorized product for this module" + req.ModuleName)
		}
	}

	if req.AlloffName != nil {
		updatedRequest.AlloffName = *req.AlloffName
	}

	if req.IsForeignDelivery != nil {
		if *req.IsForeignDelivery {
			updatedRequest.IsForeignDelivery = true
		} else {
			updatedRequest.IsForeignDelivery = false
		}
	}

	if req.OriginalPrice != nil {
		updatedRequest.OriginalPrice = float32(*req.OriginalPrice)
	}

	if req.DiscountedPrice != nil {
		updatedRequest.DiscountedPrice = float32(*req.DiscountedPrice)
	}

	if req.BrandKeyName != nil && req.BrandKeyName != &updatedRequest.Brand.KeyName {
		brand, err := ioc.Repo.Brands.GetByKeyname(*req.BrandKeyName)
		if err != nil {
			return nil, err
		}
		updatedRequest.Brand = brand
	}

	// TODO: ProductTypes를 바꾸는 작업이 되어야한다.
	if req.ProductTypes != nil && len(req.ProductTypes) > 0 {
		allofftypes := []domain.AlloffProductType{
			domain.Female,
			domain.Female,
			domain.Kids,
			domain.Sports,
		}
		for _, reqPdType := range req.ProductTypes {
			if reqPdType == grpcServer.ProductType_FEMALE {
				allofftypes = append(allofftypes, domain.Female)
			} else if reqPdType == grpcServer.ProductType_MALE {
				allofftypes = append(allofftypes, domain.Male)
			} else if reqPdType == grpcServer.ProductType_KIDS {
				allofftypes = append(allofftypes, domain.Kids)
			}
		}
		updatedRequest.ProductType = allofftypes
	}

	if req.Inventory != nil {
		invDaos := []*domain.InventoryDAO{}
		for _, inv := range req.Inventory {
			invDaos = append(invDaos, &domain.InventoryDAO{
				Quantity: int(inv.Quantity),
				Size:     inv.Size,
			})
		}
		updatedRequest.Inventory = invDaos
	}

	if req.Description != nil {
		updatedRequest.Description = req.Description
	}

	if req.DescriptionInfos != nil {
		updatedRequest.DescriptionInfos = req.DescriptionInfos
	}

	if req.ProductInfos != nil {
		updatedRequest.Information = req.ProductInfos
	}

	if req.EarliestDeliveryDays != nil {
		updatedRequest.EarliestDeliveryDays = int(*req.EarliestDeliveryDays)
	}

	if req.LatestDeliveryDays != nil {
		updatedRequest.LatestDeliveryDays = int(*req.LatestDeliveryDays)
	}

	if req.IsRefundPossible != nil {
		updatedRequest.IsRefundPossible = *req.IsRefundPossible
	}

	if req.Images != nil {
		updatedRequest.Images = req.Images
	}

	if req.DescriptionImages != nil {
		updatedRequest.DescriptionImages = req.DescriptionImages
	}

	if req.IsRemoved != nil {
		updatedRequest.IsRemoved = *req.IsRemoved
	}

	if req.AlloffCategoryId != nil {
		alloffcat, err := ioc.Repo.AlloffCategories.Get(*req.AlloffCategoryId)
		if err != nil {
			config.Logger.Error("err occured on build product alloff category : alloffcat ID"+*req.AlloffCategoryId, zap.Error(err))
		}
		updatedRequest.AlloffCategory = alloffcat
	}

	if req.ProductId != nil {
		updatedRequest.ProductID = *req.ProductId
	}

	if req.ProductUrl != nil {
		updatedRequest.ProductUrl = *req.ProductUrl
	}

	if req.RefundFee != nil {
		updatedRequest.RefundFee = int(*req.RefundFee)
	}

	if req.IsSoldout != nil {
		updatedRequest.IsSoldout = *req.IsSoldout
	}

	if req.ThumbnailImage != nil {
		updatedRequest.ThumbnailImage = *req.ThumbnailImage
	}

	newPdInfoDao, err := productinfo.UpdateProductInfo(pdInfoDao, updatedRequest)
	if err != nil {
		return nil, err
	}

	pdMessage := mapper.ProductInfoMapper(newPdInfoDao)
	return &grpcServer.EditProductResponse{
		Product: pdMessage,
	}, nil
}
