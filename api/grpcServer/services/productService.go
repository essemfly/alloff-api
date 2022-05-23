package services

import (
	"context"
	"errors"
	"log"

	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/broker"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/product"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
)

type ProductService struct {
	grpcServer.ProductServer
}

func (s *ProductService) GetProduct(ctx context.Context, req *grpcServer.GetProductRequest) (*grpcServer.GetProductResponse, error) {
	pdDao, err := ioc.Repo.Products.Get(req.AlloffProductId)
	if err != nil {
		return nil, err
	}

	return &grpcServer.GetProductResponse{
		Product: mapper.ProductMapper(pdDao),
	}, nil
}

func (s *ProductService) ListProducts(ctx context.Context, req *grpcServer.ListProductsRequest) (*grpcServer.ListProductsResponse, error) {
	moduleName := ""
	if req.ModuleName != nil {
		moduleName = *req.ModuleName
	}

	categoryID := ""
	if req.Query.CategoryId != nil {
		categoryID = *req.Query.CategoryId
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

	query := product.ProductListInput{
		Offset:                    int(req.Offset),
		Limit:                     int(req.Limit),
		BrandID:                   "",
		CategoryID:                categoryID,
		AlloffCategoryID:          alloffCategoryID,
		Keyword:                   searchKeyword,
		Modulename:                moduleName,
		IncludeClassifiedType:     classifiedType,
		IncludeSpecialProductType: product.ALL_PRODUCTS,
		PriceRanges:               priceRanges,
		PriceSorting:              priceSorting,
	}

	products, cnt, err := product.Listing(query)
	if err != nil {
		return nil, err
	}

	pds := []*grpcServer.ProductMessage{}

	for _, pd := range products {
		pds = append(pds, mapper.ProductMapper(pd))
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
	invDaos := []domain.InventoryDAO{}
	for _, inv := range req.Inventory {
		invDaos = append(invDaos, domain.InventoryDAO{
			Size:     inv.Size,
			Quantity: int(inv.Quantity),
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
		OriginalPrice:        originalPrice,
		DiscountedPrice:      discountedPrice,
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
	}

	pdInfoDao, err := product.AddProductInfo(addRequest)
	if err != nil {
		return nil, err
	}

	pdMessage := mapper.ProductMapper(pdInfoDao)

	return &grpcServer.CreateProductResponse{
		Product: pdMessage,
	}, nil
}

func (s *ProductService) EditProduct(ctx context.Context, req *grpcServer.EditProductRequest) (*grpcServer.EditProductResponse, error) {

	pdDao, err := ioc.Repo.Products.Get(req.AlloffProductId)
	if err != nil {
		return nil, err
	}

	if req.ModuleName != "" && req.ModuleName != "manual" {
		if pdDao.ProductInfo.Source.CrawlModuleName != req.ModuleName {
			return nil, errors.New("not authorized product for this module" + req.ModuleName)
		}
	}

	if req.AlloffName != nil {
		pdDao.AlloffName = *req.AlloffName
	}

	if req.IsForeignDelivery != nil {
		if *req.IsForeignDelivery {
			pdDao.ProductInfo.Source.IsForeignDelivery = true
		} else {
			pdDao.ProductInfo.Source.IsForeignDelivery = false
		}
	}

	if req.OriginalPrice != nil {
		pdDao.ProductInfo.Price.OriginalPrice = float32(*req.OriginalPrice)
		pdDao.OriginalPrice = int(*req.OriginalPrice)
	}

	if req.DiscountedPrice != nil {
		pdDao.DiscountedPrice = int(*req.DiscountedPrice)
	}

	alloffPriceDiscountRate := pdDao.DiscountRate

	if req.BrandKeyName != nil {
		brand, err := ioc.Repo.Brands.GetByKeyname(*req.BrandKeyName)
		if err != nil {
			return nil, err
		}
		pdDao.ProductInfo.Brand = brand
	}

	if req.Inventory != nil {
		invDaos := []*domain.InventoryDAO{}
		for _, inv := range req.Inventory {
			invDaos = append(invDaos, &domain.InventoryDAO{
				Size:     inv.Size,
				Quantity: int(inv.Quantity),
			})
		}
		pdDao.Inventory = invDaos
	}

	if req.Description != nil {
		pdDao.ProductInfo.SalesInstruction.Description.Texts = req.Description
	}

	if req.DescriptionInfos != nil {
		pdDao.ProductInfo.SalesInstruction.Description.Infos = req.DescriptionInfos
	}

	if req.ProductInfos != nil {
		pdDao.ProductInfo.SalesInstruction.Information = req.ProductInfos
	}

	if req.EarliestDeliveryDays != nil {
		pdDao.ProductInfo.SalesInstruction.DeliveryDescription.EarliestDeliveryDays = int(*req.EarliestDeliveryDays)
	}

	if req.LatestDeliveryDays != nil {
		pdDao.ProductInfo.SalesInstruction.DeliveryDescription.LatestDeliveryDays = int(*req.LatestDeliveryDays)
	}

	if req.IsRefundPossible != nil {
		pdDao.ProductInfo.SalesInstruction.CancelDescription.RefundAvailable = *req.IsRefundPossible
		pdDao.ProductInfo.SalesInstruction.CancelDescription.ChangeAvailable = *req.IsRefundPossible
	}

	if req.Images != nil {
		pdDao.ProductInfo.Images = req.Images
		pdDao.ProductInfo.CachedImages = req.Images
	}

	if req.DescriptionImages != nil {
		pdDao.ProductInfo.SalesInstruction.Description.Images = req.DescriptionImages
	}

	if req.IsRemoved != nil {
		pdDao.IsRemoved = *req.IsRemoved
	}

	if req.AlloffCategoryId != nil {
		// productCatDao := classifier.ClassifyProducts(*req.AlloffCategoryId)
		// pdDao.UpdateAlloffCategory(productCatDao)
		pdDao.ProductInfo.AlloffCategory.Touched = true
	}

	if req.ProductId != nil {
		pdDao.ProductInfo.ProductID = *req.ProductId
	}

	if req.ProductUrl != nil {
		pdDao.ProductInfo.ProductUrl = *req.ProductUrl
	}

	if req.RefundFee != nil {
		pdDao.ProductInfo.SalesInstruction.CancelDescription.ChangeFee = int(*req.RefundFee)
		pdDao.ProductInfo.SalesInstruction.CancelDescription.RefundFee = int(*req.RefundFee)
	}

	if req.IsSoldout != nil {
		pdDao.IsSoldout = *req.IsSoldout
	}

	if !pdDao.IsSoldout {
		pdDao.CheckSoldout()
	}

	if alloffPriceDiscountRate > pdDao.ProductInfo.Brand.MaxDiscountRate {
		pdDao.ProductInfo.Brand.MaxDiscountRate = alloffPriceDiscountRate
	}

	if req.ThumbnailImage != nil {
		pdDao.ThumbnailImage = *req.ThumbnailImage
	}

	newPdDao, err := ioc.Repo.Products.Upsert(pdDao)
	if err != nil {
		return nil, err
	}

	if newPdDao.ProductGroupID != "" {
		pg, err := ioc.Repo.ProductGroups.Get(newPdDao.ProductGroupID)
		if err != nil {
			log.Println("err found in product group update", err)
		} else {
			go broker.ProductGroupSyncer(pg)
			if pg.ExhibitionID != "" {
				exDao, err := ioc.Repo.Exhibitions.Get(pg.ExhibitionID)
				if err != nil {
					log.Println("exhibbition find error", err)
				} else {
					go broker.ExhibitionSyncer(exDao)
				}
			}
		}
	}

	pdMessage := mapper.ProductMapper(newPdDao)
	return &grpcServer.EditProductResponse{
		Product: pdMessage,
	}, nil
}
