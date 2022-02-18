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
	for _, guide := range req.SizeGuide {
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
		BackImgUrl:    req.BackImageUrl,
		Onpopular:     false,
		Created:       time.Now(),
		IsOpen:        false,
		Modulename:    "manual",
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

func (s *BrandService) EditBrand(ctx context.Context, req *grpcServer.EditBrandRequest) (*grpcServer.EditBrandResponse, error) {
	brandDao, err := ioc.Repo.Brands.GetByKeyname(req.Keyname)
	if err != nil {
		return nil, err
	}
	if req.Korname != nil {
		brandDao.KorName = *req.Korname
	}
	if req.Engname != nil {
		brandDao.EngName = *req.Engname
	}
	if req.LogoImageUrl != nil {
		brandDao.LogoImgUrl = *req.LogoImageUrl
	}
	if req.BackImageUrl != nil {
		brandDao.BackImgUrl = *req.BackImageUrl
	}
	if req.Description != nil {
		brandDao.Description = *req.Description
	}
	if req.IsPopular != nil {
		brandDao.Onpopular = *req.IsPopular
	}
	if req.IsOpen != nil {
		brandDao.IsOpen = *req.IsOpen
	}
	if req.InMaintenance != nil {
		brandDao.InMaintenance = *req.InMaintenance
	}
	if req.SizeGuide != nil {
		guides := []domain.SizeGuideDAO{}

		for _, guide := range req.SizeGuide {
			guides = append(guides, domain.SizeGuideDAO{
				Label:  guide.Label,
				ImgUrl: guide.ImageUrl,
			})
		}

		brandDao.SizeGuide = guides
	}

	newBrandDao, err := ioc.Repo.Brands.Upsert(brandDao)
	if err != nil {
		return nil, err
	}

	return &grpcServer.EditBrandResponse{
		Brand: mapper.BrandMapper(newBrandDao),
	}, nil
}
