package mapper

import (
	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func BrandMapper(brand *domain.BrandDAO) *grpcServer.BrandMessage {
	return &grpcServer.BrandMessage{
		Korname:       brand.KorName,
		Keyname:       brand.KeyName,
		LogoImageUrl:  brand.LogoImgUrl,
		IsPopular:     brand.Onpopular,
		IsOpen:        brand.IsOpen,
		InMaintenance: brand.InMaintenance,
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
