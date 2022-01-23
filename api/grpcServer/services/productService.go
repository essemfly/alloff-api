package services

import (
	"context"

	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
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

func (s *ProductService) PutProduct(ctx context.Context, req *grpcServer.PutProductRequest) (*grpcServer.PutProductResponse, error) {
	pdDao, err := ioc.Repo.Products.Get(req.AlloffProductId)
	if err != nil {
		return nil, err
	}

	pdDao.SpecialPrice = int(req.SpecialPrice)
	newPdDao, err := ioc.Repo.Products.Upsert(pdDao)
	if err != nil {
		return nil, err
	}
	return &grpcServer.PutProductResponse{
		Product: mapper.ProductMapper(newPdDao),
	}, nil
}

func (s *ProductService) ListProducts(ctx context.Context, req *grpcServer.ListProductsRequest) (*grpcServer.ListProductsResponse, error) {
	brandID := ""
	if req.Query.BrandId != nil {
		brandID = *req.Query.BrandId
	}
	categoryID := ""
	if req.Query.CategoryId != nil {
		categoryID = *req.Query.CategoryId
	}

	products, cnt, err := product.ProductsListing(int(req.Offset), int(req.Limit), brandID, categoryID, "", nil)
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
	}

	return ret, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, req *grpcServer.CreateProductRequest) (*grpcServer.CreateProductResponse, error) {
	specialPrice := int(req.SpecialPrice)
	originalPrice := specialPrice
	if req.OriginalPrice != nil {
		originalPrice = int(*req.OriginalPrice)
	}
	discountedPrice := specialPrice
	if req.DiscountedPrice != nil {
		discountedPrice = int(*req.DiscountedPrice)
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
	addRequest := &product.ProductManuelAddRequest{
		AlloffName:           req.AlloffName,
		IsForeignDelivery:    req.IsForeignDelivery,
		ProductID:            productID,
		OriginalPrice:        originalPrice,
		DiscountedPrice:      discountedPrice,
		SpecialPrice:         int(req.SpecialPrice),
		BrandKeyName:         req.BrandKeyName,
		Inventory:            invDaos,
		Description:          req.Description,
		EarliestDeliveryDays: int(req.EarliestDeliveryDays),
		LatestDeliveryDays:   int(req.LatestDeliveryDays),
		IsRefundPossible:     req.IsRefundPossible,
		RefundFee:            int(req.RefundFee),
		Images:               req.Images,
		DescriptionImages:    req.DescriptionImages,
	}

	pdDao, err := product.AddProductInManuel(addRequest)
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

	if req.IsRefundPossible != nil {
		pdDao.SalesInstruction.CancelDescription.RefundAvailable = *req.IsRefundPossible
		pdDao.SalesInstruction.CancelDescription.ChangeAvailable = *req.IsRefundPossible
	}

	if req.Images != nil {
		pdDao.ProductInfo.Images = req.Images
	}

	if req.DescriptionImages != nil {
		pdDao.SalesInstruction.Description.Images = req.DescriptionImages
	}

	if req.IsRemoved != nil {
		pdDao.Removed = *req.IsRemoved
	}

	newPdDao, err := ioc.Repo.Products.Upsert(pdDao)
	if err != nil {
		return nil, err
	}

	pdMessage := mapper.ProductMapper(newPdDao)
	return &grpcServer.EditProductResponse{
		Product: pdMessage,
	}, nil
}
