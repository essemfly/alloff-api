package services

import (
	"context"
	"time"

	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

type BrandService struct {
	grpcServer.BrandServer
}

func (s *BrandService) CreateBrand(ctx context.Context, req *grpcServer.CreateBrandRequest) (*grpcServer.CreateBrandResponse, error) {

	sizeGuideDaos := []domain.SizeGuideDAO{}
	for _, guide := range req.Sizeguide {
		sizeGuideDaos = append(sizeGuideDaos, domain.SizeGuideDAO{
			Label:  guide.Label,
			ImgUrl: guide.ImageUrl,
		})
	}

	brandDao := &domain.BrandDAO{
		KorName:       req.Korname,
		EngName:       req.Engname,
		KeyName:       req.Keyname,
		Description:   req.Description,
		LogoImgUrl:    req.LogoImageUrl,
		Onpopular:     false,
		Created:       time.Now(),
		IsOpen:        false,
		Modulename:    "manuel",
		InMaintenance: false,
		SizeGuide:     sizeGuideDaos,
	}
	newBrand, err := ioc.Repo.Brands.Upsert(brandDao)
	if err != nil {
		return nil, err
	}

	return &grpcServer.CreateBrandResponse{
		Brand: mapper.BrandMapper(newBrand),
	}, nil
}

func (s *BrandService) ListBrand(ctx context.Context, req *grpcServer.ListBrandRequest) (*grpcServer.ListBrandResponse, error) {
	offset, limit := 0, 1000
	brandDaos, _, err := ioc.Repo.Brands.List(offset, limit, nil, nil)
	if err != nil {
		return nil, err
	}

	brands := []*grpcServer.BrandMessage{}
	for _, brand := range brandDaos {
		brands = append(brands, mapper.BrandMapper(brand))
	}

	return &grpcServer.ListBrandResponse{
		Brands: brands,
	}, nil
}
