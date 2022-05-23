package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"time"
)

func MapExhibition(exDao *domain.ExhibitionDAO, brief bool) *model.Exhibition {
	pgs := []*model.ProductGroup{}

	if !brief {
		for _, pg := range exDao.ProductGroups {
			pgs = append(pgs, MapProductGroup(pg))
		}
	}

	numProducts := 0
	for _, pg := range exDao.ProductGroups {
		numProducts += len(pg.Products)
	}

	return &model.Exhibition{
		ID:             exDao.ID.Hex(),
		ExhibitionType: MapExhibitionType(exDao.ExhibitionType),
		Title:          exDao.Title,
		SubTitle:       exDao.SubTitle,
		Description:    exDao.Description,
		Tags:           exDao.Tags,
		BannerImage:    exDao.BannerImage,
		ThumbnailImage: exDao.ThumbnailImage,
		ProductGroups:  pgs,
		StartTime:      exDao.StartTime.Add(9 * time.Hour).String(),
		FinishTime:     exDao.FinishTime.Add(9 * time.Hour).String(),
		NumAlarms:      exDao.NumAlarms,
		MetaInfos:      nil,
	}
}

func MapExhibitionType(exhibitionType domain.ExhibitionType) model.ExhibitionType {
	var res model.ExhibitionType
	switch exhibitionType {
	case domain.EXHIBITION_NORMAL:
		res = model.ExhibitionTypeNormal
	case domain.EXHIBITION_TIMEDEAL:
		res = model.ExhibitionTypeTimedeal
	case domain.EXHIBITION_GROUPDEAL:
		res = model.ExhibitionTypeGroupdeal
	}
	return res
}
