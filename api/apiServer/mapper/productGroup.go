package mapper

import (
	"sort"
	"time"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapProductGroupDao(pgDao *domain.ProductGroupDAO) *model.ProductGroup {
	var pds []*model.Product
	var soldouts []*model.Product

	sort.Slice(pgDao.Products, func(i, j int) bool {
		return pgDao.Products[i].Priority < pgDao.Products[j].Priority
	})

	for _, productInPg := range pgDao.Products {
		if productInPg.Product.Removed {
			continue
		}

		pd := MapProductDaoToProduct(productInPg.Product)
		if pd.Soldout {
			soldouts = append(soldouts, pd)
		} else {
			pds = append(pds, pd)
		}
	}

	pds = append(pds, soldouts...)

	pg := &model.ProductGroup{
		ID:          pgDao.ID.Hex(),
		Title:       pgDao.Title,
		ShortTitle:  pgDao.ShortTitle,
		Instruction: pgDao.Instruction,
		ImgURL:      pgDao.ImgUrl,
		NumAlarms:   pgDao.NumAlarms,
		Products:    pds,
		StartTime:   pgDao.StartTime.Add(9 * time.Hour).String(),
		FinishTime:  pgDao.FinishTime.Add(9 * time.Hour).String(),
		SetAlarm:    false,
	}
	return pg
}

func MapExhibition(exDao *domain.ExhibitionDAO, brief bool) *model.Exhibition {
	pgs := []*model.ProductGroup{}

	if !brief {
		for _, pg := range exDao.ProductGroups {
			pgs = append(pgs, MapProductGroupDao(pg))
		}
	}

	return &model.Exhibition{
		ID:             exDao.ID.Hex(),
		BannerImage:    exDao.BannerImage,
		ThumbnailImage: exDao.ThumbnailImage,
		Title:          exDao.Title,
		SubTitle:       exDao.SubTitle,
		Description:    exDao.Description,
		ProductGroups:  pgs,
		StartTime:      exDao.StartTime.Add(9 * time.Hour).String(),
		FinishTime:     exDao.FinishTime.Add(9 * time.Hour).String(),
		ExhibitionType: MapExhibitionType(exDao.ExhibitionType),
		TargetSales:    exDao.TargetSales,
		CurrentSales:   exDao.GetCurrentSales(),
	}
}

func MapExhibitionType(enum domain.ExhibitionType) model.ExhibitionType {
	switch enum {
	case domain.EXHIBITION_GROUPDEAL:
		return model.ExhibitionTypeGroupdeal
	case domain.EXHIBITION_NORMAL:
		return model.ExhibitionTypeNormal
	case domain.EXHIBITION_TIMEDEAL:
		return model.ExhibitionTypeTimedeal
	}
	return model.ExhibitionTypeNormal
}
