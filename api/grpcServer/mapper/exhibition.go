package mapper

import (
	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func ExhibitionMapper(exDao *domain.ExhibitionDAO, brief bool) *grpcServer.ExhibitionMessage {
	pgs := []*grpcServer.ProductGroupMessage{}

	if !brief {
		for _, pg := range exDao.ProductGroups {
			pgs = append(pgs, ProductGroupMapper(pg))
		}
	}

	return &grpcServer.ExhibitionMessage{
		ExhibitionId:   exDao.ID.Hex(),
		BannerImage:    exDao.BannerImage,
		ThumbnailImage: exDao.ThumbnailImage,
		Title:          exDao.Title,
		Subtitle:       exDao.SubTitle,
		Description:    exDao.Description,
		StartTime:      exDao.StartTime.String(),
		FinishTime:     exDao.FinishTime.String(),
		Pgs:            pgs,
		IsLive:         exDao.IsLive,
	}
}
