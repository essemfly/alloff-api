package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapProductGroupDao(pgDao *domain.ProductGroupDAO) *model.ProductGroup {
	var pds []*model.Product
	for _, productInPg := range pgDao.Products {
		pdDao, _ := ioc.Repo.Products.Get(productInPg.ProductID.Hex())
		pd := MapProductDaoToProduct(pdDao)

		if pd != nil {
			pds = append(pds, pd)
		}
	}

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
		ShortTitle:     exDao.ShortTitle,
		ProductGroups:  pgs,
	}
}
