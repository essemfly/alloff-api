package services

import (
	"context"
	"errors"

	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/product"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
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

	classifiedType := product.NO_MATTER_CLASSIFIED
	if req.Query.IsClassifiedDone != nil {
		if *req.Query.IsClassifiedDone {
			classifiedType = product.CLASSIFIED_DONE
		} else {
			classifiedType = product.NOT_CLASSIFIED
		}
	}

	var priceSorting product.PriceSortingType
	priceRanges := []product.PriceRangeType{}
	if req.Query.Options != nil {
		priceRanges, priceSorting = mapper.ProductSortingAndRangesMapper(req.Query.Options)
	}

	query := product.ProductInfoListInput{
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

	products, cnt, err := product.ListProductInfos(query)
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
	invDaos := []*domain.AlloffInventoryDAO{}
	for _, inv := range req.Inventory {
		size, _ := ioc.Repo.AlloffSizes.Get(inv.AlloffSize.AlloffSizeId)
		invDaos = append(invDaos, &domain.AlloffInventoryDAO{
			AlloffSize: *size,
			Quantity:   int(inv.Quantity),
		})
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

	addRequest := &product.AddMetaInfoRequest{
		AlloffName:           req.AlloffName,
		ProductID:            productID,
		ProductUrl:           productUrl,
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
		Inventory:            nil,
		AlloffInventory:      invDaos,
		ModuleName:           moduleName,
		IsInventoryMapped:    true,
	}

	pdInfoDao, err := product.AddProductInfo(addRequest)
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

	if req.ModuleName != "" && req.ModuleName != "manual" {
		if pdInfoDao.Source.CrawlModuleName != req.ModuleName {
			return nil, errors.New("not authorized product for this module" + req.ModuleName)
		}
	}

	if req.AlloffName != nil {
		pdInfoDao.AlloffName = *req.AlloffName
	}

	if req.IsForeignDelivery != nil {
		if *req.IsForeignDelivery {
			pdInfoDao.Source.IsForeignDelivery = true
		} else {
			pdInfoDao.Source.IsForeignDelivery = false
		}
	}

	if req.OriginalPrice != nil {
		pdInfoDao.Price.OriginalPrice = int(*req.OriginalPrice)
	}

	if req.DiscountedPrice != nil {
		pdInfoDao.Price.CurrentPrice = int(*req.DiscountedPrice)
	}

	alloffPriceDiscountRate := pdInfoDao.Price.DiscountRate

	if req.BrandKeyName != nil {
		brand, err := ioc.Repo.Brands.GetByKeyname(*req.BrandKeyName)
		if err != nil {
			return nil, err
		}
		pdInfoDao.Brand = brand
	}

	// TODO: AlloffInventory를 바꿔주는 작업이 되어야한다.
	if req.Inventory != nil {
		invDaos := []*domain.AlloffInventoryDAO{}
		for _, inv := range req.Inventory {
			size, _ := ioc.Repo.AlloffSizes.Get(inv.AlloffSize.AlloffSizeId)
			invDaos = append(invDaos, &domain.AlloffInventoryDAO{
				AlloffSize: *size,
				Quantity:   int(inv.Quantity),
			})
		}
		pdInfoDao.AlloffInventory = invDaos
	}

	if req.Description != nil {
		pdInfoDao.SalesInstruction.Description.Texts = req.Description
	}

	if req.DescriptionInfos != nil {
		pdInfoDao.SalesInstruction.Description.Infos = req.DescriptionInfos
	}

	if req.ProductInfos != nil {
		pdInfoDao.SalesInstruction.Information = req.ProductInfos
	}

	if req.EarliestDeliveryDays != nil {
		pdInfoDao.SalesInstruction.DeliveryDescription.EarliestDeliveryDays = int(*req.EarliestDeliveryDays)
	}

	if req.LatestDeliveryDays != nil {
		pdInfoDao.SalesInstruction.DeliveryDescription.LatestDeliveryDays = int(*req.LatestDeliveryDays)
	}

	if req.IsRefundPossible != nil {
		pdInfoDao.SalesInstruction.CancelDescription.RefundAvailable = *req.IsRefundPossible
		pdInfoDao.SalesInstruction.CancelDescription.ChangeAvailable = *req.IsRefundPossible
	}

	if req.Images != nil {
		pdInfoDao.Images = req.Images
		pdInfoDao.CachedImages = req.Images
	}

	if req.DescriptionImages != nil {
		pdInfoDao.SalesInstruction.Description.Images = req.DescriptionImages
	}

	if req.IsRemoved != nil {
		pdInfoDao.IsRemoved = *req.IsRemoved
	}

	// TODO: Update Alloff Category should be modified
	if req.AlloffCategoryId != nil {
		// productCatDao := classifier.ClassifyProducts(*req.AlloffCategoryId)
		// pdInfoDao.UpdateAlloffCategory(productCatDao)
		pdInfoDao.AlloffCategory.Touched = true
	}

	if req.ProductId != nil {
		pdInfoDao.ProductID = *req.ProductId
	}

	if req.ProductUrl != nil {
		pdInfoDao.ProductUrl = *req.ProductUrl
	}

	if req.RefundFee != nil {
		pdInfoDao.SalesInstruction.CancelDescription.ChangeFee = int(*req.RefundFee)
		pdInfoDao.SalesInstruction.CancelDescription.RefundFee = int(*req.RefundFee)
	}

	if req.IsSoldout != nil {
		pdInfoDao.IsSoldout = *req.IsSoldout
	}

	if !pdInfoDao.IsSoldout {
		pdInfoDao.CheckSoldout()
	}

	if alloffPriceDiscountRate > pdInfoDao.Brand.MaxDiscountRate {
		pdInfoDao.Brand.MaxDiscountRate = alloffPriceDiscountRate
	}

	if req.ThumbnailImage != nil {
		pdInfoDao.ThumbnailImage = *req.ThumbnailImage
	}

	newPdInfoDao, err := ioc.Repo.ProductMetaInfos.Upsert(pdInfoDao)
	if err != nil {
		return nil, err
	}

	// TODO: pdInfo에 맞는 pd들을 업데이트 시키는 작업이 필요하다.
	// if newPdDao.ProductGroupID != "" {
	// 	pg, err := ioc.Repo.ProductGroups.Get(newPdDao.ProductGroupID)
	// 	if err != nil {
	// 		log.Println("err found in product group update", err)
	// 	} else {
	// 		go broker.ProductGroupSyncer(pg)
	// 		if pg.ExhibitionID != "" {
	// 			exDao, err := ioc.Repo.Exhibitions.Get(pg.ExhibitionID)
	// 			if err != nil {
	// 				log.Println("exhibbition find error", err)
	// 			} else {
	// 				go broker.ExhibitionSyncer(exDao)
	// 			}
	// 		}
	// 	}
	// }

	pdMessage := mapper.ProductInfoMapper(newPdInfoDao)
	return &grpcServer.EditProductResponse{
		Product: pdMessage,
	}, nil
}
