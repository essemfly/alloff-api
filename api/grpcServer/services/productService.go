package services

import (
	"context"
	"errors"
	"log"

	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/broker"
	"github.com/lessbutter/alloff-api/pkg/classifier"
	"github.com/lessbutter/alloff-api/pkg/product"
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
	brandID := ""
	if req.Query.BrandId != nil {
		brandID = *req.Query.BrandId
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
	products, cnt, err := product.ProductsSearchListing(int(req.Offset), int(req.Limit), moduleName, brandID, categoryID, alloffCategoryID, searchKeyword, "", nil)
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
	specialPrice := discountedPrice
	if req.SpecialPrice != nil {
		specialPrice = int(*req.SpecialPrice)
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
	}
	alloffCatID := ""
	if req.AlloffCategoryId != nil {
		alloffCatID = *req.AlloffCategoryId
	}

	addRequest := &product.ProductManualAddRequest{
		AlloffName:           req.AlloffName,
		IsForeignDelivery:    req.IsForeignDelivery,
		ProductID:            productID,
		OriginalPrice:        originalPrice,
		DiscountedPrice:      discountedPrice,
		SpecialPrice:         specialPrice,
		BrandKeyName:         req.BrandKeyName,
		Inventory:            invDaos,
		Description:          req.Description,
		EarliestDeliveryDays: int(req.EarliestDeliveryDays),
		LatestDeliveryDays:   int(req.LatestDeliveryDays),
		IsRefundPossible:     req.IsRefundPossible,
		RefundFee:            int(req.RefundFee),
		Images:               req.Images,
		DescriptionImages:    req.DescriptionImages,
		ModuleName:           moduleName,
		AlloffCategoryID:     alloffCatID,
	}

	pdDao, err := product.AddProductManually(addRequest)
	if err != nil {
		return nil, err
	}

	pdMessage := mapper.ProductMapper(pdDao)

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

	if req.SpecialPrice != nil {
		pdDao.SpecialPrice = int(*req.SpecialPrice)
	}

	if req.BrandKeyName != nil {
		brand, err := ioc.Repo.Brands.GetByKeyname(*req.BrandKeyName)
		if err != nil {
			return nil, err
		}
		pdDao.ProductInfo.Brand = brand
	}

	if req.Inventory != nil {
		invDaos := []domain.InventoryDAO{}
		for _, inv := range req.Inventory {
			invDaos = append(invDaos, domain.InventoryDAO{
				Size:     inv.Size,
				Quantity: int(inv.Quantity),
			})
		}
		pdDao.Inventory = invDaos
	}

	if req.Description != nil {
		pdDao.SalesInstruction.Description.Texts = req.Description
	}

	if req.EarliestDeliveryDays != nil {
		pdDao.SalesInstruction.DeliveryDescription.EarliestDeliveryDays = int(*req.EarliestDeliveryDays)
	}

	if req.LatestDeliveryDays != nil {
		pdDao.SalesInstruction.DeliveryDescription.LatestDeliveryDays = int(*req.LatestDeliveryDays)
	}

	if req.IsRefundPossible != nil {
		pdDao.SalesInstruction.CancelDescription.RefundAvailable = *req.IsRefundPossible
		pdDao.SalesInstruction.CancelDescription.ChangeAvailable = *req.IsRefundPossible
	}

	if req.Images != nil {
		pdDao.ProductInfo.Images = req.Images
		pdDao.Images = req.Images
	}

	if req.DescriptionImages != nil {
		pdDao.SalesInstruction.Description.Images = req.DescriptionImages
	}

	if req.IsRemoved != nil {
		pdDao.Removed = *req.IsRemoved
	}

	if req.AlloffCategoryId != nil {
		productCatDao := classifier.ClassifyProducts(*req.AlloffCategoryId)
		pdDao.UpdateAlloffCategory(productCatDao)
	}

	if req.ProductId != nil {
		pdDao.ProductInfo.ProductID = *req.ProductId
	}

	if req.RefundFee != nil {
		pdDao.SalesInstruction.CancelDescription.ChangeFee = int(*req.RefundFee)
		pdDao.SalesInstruction.CancelDescription.RefundFee = int(*req.RefundFee)
	}

	pdDao.CheckSoldout()

	newPdDao, err := ioc.Repo.Products.Upsert(pdDao)
	if err != nil {
		return nil, err
	}

	if newPdDao.ProductGroupId != "" {
		pg, err := ioc.Repo.ProductGroups.Get(newPdDao.ProductGroupId)
		if err != nil {
			log.Println("err found in product group update", err)
		} else {
			broker.ProductGroupSyncer(pg)
			if pg.ExhibitionID != "" {
				exDao, err := ioc.Repo.Exhibitions.Get(pg.ExhibitionID)
				if err != nil {
					log.Println("exhibbition find error", err)
				} else {
					broker.ExhibitionSyncer(exDao)
				}
			}
		}
	}

	pdMessage := mapper.ProductMapper(newPdDao)
	return &grpcServer.EditProductResponse{
		Product: pdMessage,
	}, nil
}
