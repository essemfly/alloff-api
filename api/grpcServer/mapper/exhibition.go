package mapper

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/exhibition"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
)

func ExhibitionMapper(exDao *domain.ExhibitionDAO, brief bool) *grpcServer.ExhibitionMessage {
	pgs := []*grpcServer.ProductGroupMessage{}

	if !brief {
		for _, pg := range exDao.ProductGroups {
			pgs = append(pgs, ProductGroupMapper(pg))
		}
	}
	sales := 0
	if exDao.ExhibitionType == domain.EXHIBITION_GROUPDEAL {
		sales = exhibition.GetCurrentSales(exDao)
	}

	return &grpcServer.ExhibitionMessage{
		ExhibitionId:      exDao.ID.Hex(),
		BannerImage:       exDao.BannerImage,
		ThumbnailImage:    exDao.ThumbnailImage,
		Title:             exDao.Title,
		Subtitle:          exDao.SubTitle,
		Description:       exDao.Description,
		RecruitStartTime:  exDao.RecruitStartTime.String(),
		StartTime:         exDao.StartTime.String(),
		FinishTime:        exDao.FinishTime.String(),
		Pgs:               pgs,
		IsLive:            exDao.IsLive,
		ExhibitionType:    ExhibitionGroupTypeMapper(exDao.ExhibitionType),
		TargetSales:       int32(exDao.TargetSales),
		CurrentSales:      int32(sales),
		Banners:           bannersMapper(exDao.Banners),
		TotalParticipants: int32(exDao.TotalParticipants),
		NumUsersRequired:  int32(exDao.NumUsersRequired),
		TotalGroups:       int32(exDao.TotalGroups),
		AllowOldUser:      exDao.AllowOldUser,
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

func bannersMapper(banners []domain.ExhibitionBanner) []*grpcServer.ExhibitionBannerMessage {
	bannerMsgs := []*grpcServer.ExhibitionBannerMessage{}
	for _, banner := range banners {
		bannerMsg := &grpcServer.ExhibitionBannerMessage{
			ImgUrl:         banner.ImgUrl,
			Title:          banner.Title,
			Subtitle:       banner.Subtitle,
			ProductGroupId: banner.ProductGroupId,
		}
		bannerMsgs = append(bannerMsgs, bannerMsg)
	}
	return bannerMsgs
}
