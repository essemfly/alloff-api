package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapTopBanner(banner *domain.TopBannerDAO) *model.TopBanner {
	return &model.TopBanner{
		ID:           banner.ID.Hex(),
		ImageURL:     banner.ImageUrl,
		ExhibitionID: banner.ExhibitionID,
		Title:        banner.Title,
		SubTitle:     banner.SubTitle,
	}
}
