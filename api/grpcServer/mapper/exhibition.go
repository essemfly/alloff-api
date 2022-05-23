package mapper

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
)

func ExhibitionMapper(exDao *domain.ExhibitionDAO, brief bool) *grpcServer.ExhibitionMessage {
	pgs := []*grpcServer.ProductGroupMessage{}

	return &grpcServer.ExhibitionMessage{
		ExhibitionId: exDao.ID.Hex(),
		// MetaInfos
		ExhibitionType: ExhibitionGroupTypeMapper(exDao.ExhibitionType),
		Title:          exDao.Title,
		Subtitle:       exDao.SubTitle,
		Description:    exDao.Description,
		// Tags
		BannerImage:    exDao.BannerImage,
		ThumbnailImage: exDao.ThumbnailImage,
		Pgs:            pgs,
		StartTime:      exDao.StartTime.String(),
		FinishTime:     exDao.FinishTime.String(),
		IsLive:         exDao.IsLive,
		// NumAlarms
	}
}

func ExhibitionGroupTypeMapper(groupType domain.ExhibitionType) grpcServer.ExhibitionType {
	switch groupType {
	case domain.EXHIBITION_TIMEDEAL:
		return grpcServer.ExhibitionType_EXHIBITION_TIMEDEAL
	case domain.EXHIBITION_NORMAL:
		return grpcServer.ExhibitionType_EXHIBITION_NORMAL
	case domain.EXHIBITION_GROUPDEAL:
		return grpcServer.ExhibitionType_EXHIBITION_GROUPDEAL
	}
	return grpcServer.ExhibitionType_EXHIBITION_NORMAL
}
