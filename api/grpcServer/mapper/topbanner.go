package mapper

import (
	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func TopbannerMapper(bannerDao *domain.TopBannerDAO) *grpcServer.TopBannerMessage {
	return &grpcServer.TopBannerMessage{
		BannerId:     bannerDao.ID.Hex(),
		BannerImage:  bannerDao.ImageUrl,
		ExhibitionId: bannerDao.ExhibitionID,
		Title:        bannerDao.Title,
		Subtitle:     bannerDao.SubTitle,
		IsLive:       bannerDao.IsLive,
		Weight:       int32(bannerDao.Weight),
	}
}
