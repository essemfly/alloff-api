package mapper

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
)

func BrandMapper(brand *domain.BrandDAO) *grpcServer.BrandMessage {
	return &grpcServer.BrandMessage{
		BrandId:       brand.ID.Hex(),
		Korname:       brand.KorName,
		Keyname:       brand.KeyName,
		Engname:       brand.EngName,
		Description:   brand.Description,
		LogoImageUrl:  brand.LogoImgUrl,
		BackImageUrl:  brand.BackImgUrl,
		IsPopular:     brand.Onpopular,
		IsOpen:        brand.IsOpen,
		InMaintenance: brand.InMaintenance,
		IsHide:        brand.IsHide,
		SizeGuide:     SizeGuideMapper(brand.SizeGuide),
	}
}

func SizeGuideMapper(sizeGuide []domain.SizeGuideDAO) []*grpcServer.SizeGuideMessage {
	guides := []*grpcServer.SizeGuideMessage{}
	for _, guide := range sizeGuide {
		guides = append(guides, &grpcServer.SizeGuideMessage{
			Label:    guide.Label,
			ImageUrl: guide.ImgUrl,
		})
	}
	return guides
}
