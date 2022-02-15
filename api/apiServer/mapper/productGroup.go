package mapper

import (
	"sort"

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
		StartTime:   pgDao.StartTime.String(),
		FinishTime:  pgDao.FinishTime.String(),
		SetAlarm:    false,
	}
	return pg
}

func MapExhibition(exDao *domain.ExhibitionDAO) *model.Exhibition {
	pgs := []*model.ProductGroup{}

	for _, pg := range exDao.ProductGroups {
		pgs = append(pgs, MapProductGroupDao(pg))
	}

	return &model.Exhibition{
		ID:             exDao.ID.Hex(),
		BannerImage:    exDao.BannerImage,
		ThumbnailImage: exDao.ThumbnailImage,
		Title:          exDao.Title,
		SubTitle:       exDao.SubTitle,
		Description:    exDao.Description,
		ProductGroups:  pgs,
		StartTime:      exDao.StartTime.String(),
		FinishTime:     exDao.FinishTime.String(),
	}
}
