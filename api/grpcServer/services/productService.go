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

	product := mapper.ProductInfoMapper(pdInfoDao)

	if req.ProductGroupId != nil {
		pdDao, err := ioc.Repo.Products.GetByMetaID(req.AlloffProductId, "", *req.ProductGroupId)
		if err != nil {
			return nil, err
		}
		product.DealProduct_Id = pdDao.ID.Hex()
	}

	return &grpcServer.GetProductResponse{
		Product: product,
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

	brandID := ""
	if req.Query.BrandId != nil {
		brandID = *req.Query.BrandId
	}

	searchKeyword := ""
	if req.Query.SearchQuery != nil {
		searchKeyword = *req.Query.SearchQuery
	}

	productUrl := ""
	if req.Query.ProductUrl != nil {
		productUrl = *req.Query.ProductUrl
	}

	alloffSizeIds := req.Query.AlloffSizeIds

	productTypes := mapper.ProductTypeReverseMapper(req.Query.ProductTypes)

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
		BrandID:               brandID,
		AlloffCategoryID:      alloffCategoryID,
		AlloffSizeIDs:         alloffSizeIds,
		ProductTypes:          productTypes,
		Keyword:               searchKeyword,
		ProductUrl:            productUrl,
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

	productTypes := mapper.ProductTypeReverseMapper(req.ProductTypes)

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

	descriptionImages := []string{}
	descriptionImages = append(descriptionImages, req.Images...)
	descriptionImages = append(descriptionImages, req.DescriptionImages...)

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
		DescriptionImages:    descriptionImages,
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

	if req.ProductTypes != nil && len(req.ProductTypes) > 0 {
		updatedRequest.ProductType = mapper.ProductTypeReverseMapper(req.ProductTypes)
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

	if len(req.Description) > 0 {
		updatedRequest.Description = req.Description
	}

	updatedRequest.DescriptionInfos = req.DescriptionInfos

	updatedRequest.Information = req.ProductInfos

	if req.EarliestDeliveryDays != nil {
		updatedRequest.EarliestDeliveryDays = int(*req.EarliestDeliveryDays)
	}

	if req.LatestDeliveryDays != nil {
		updatedRequest.LatestDeliveryDays = int(*req.LatestDeliveryDays)
	}

	if req.IsRefundPossible != nil {
		updatedRequest.IsRefundPossible = *req.IsRefundPossible
	}

	if len(req.Images) > 0 {
		updatedRequest.Images = req.Images
	}

	if len(req.DescriptionImages) > 0 {
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

	newPdInfoDao, err := productinfo.UpdateProductInfo(pdInfoDao, updatedRequest, "GRPC")
	if err != nil {
		return nil, err
	}

	pdMessage := mapper.ProductInfoMapper(newPdInfoDao)
	return &grpcServer.EditProductResponse{
		Product: pdMessage,
	}, nil
}

func (s *ProductService) CacheProductImages(ctx context.Context, req *grpcServer.ProductsUpdateRequest) (*grpcServer.ProductsUpdateResponse, error) {
	out := new(grpcServer.ProductsUpdateResponse)
	return out, nil
}

func (s *ProductService) CrawlProductInventories(ctx context.Context, req *grpcServer.ProductsUpdateRequest) (*grpcServer.ProductsUpdateResponse, error) {
	out := new(grpcServer.ProductsUpdateResponse)
	return out, nil
}

func (s *ProductService) CrawlProductPrices(ctx context.Context, req *grpcServer.ProductsPriceUpdateRequest) (*grpcServer.ProductsUpdateResponse, error) {
	out := new(grpcServer.ProductsUpdateResponse)
	return out, nil
}

func (s *ProductService) TranslateProduct(ctx context.Context, req *grpcServer.ProductsUpdateRequest) (*grpcServer.ProductsUpdateResponse, error) {
	out := new(grpcServer.ProductsUpdateResponse)
	return out, nil
}
