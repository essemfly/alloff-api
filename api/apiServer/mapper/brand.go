package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapBrandDaoToBrand(brDao *domain.BrandDAO, includeCategory bool) *model.Brand {
	var cats []*model.Category
	if includeCategory {
		for _, catDao := range brDao.Category {
			cats = append(cats, MapCategoryDaoToCategory(catDao))
		}
	}

	sizes := []*model.SizeGuide{}
	for _, guide := range brDao.SizeGuide {
		sizes = append(sizes, &model.SizeGuide{
			Label:  guide.Label,
			ImgURL: guide.ImgUrl,
		})
	}

	return &model.Brand{
		ID:              brDao.ID.Hex(),
		EngName:         brDao.EngName,
		KorName:         brDao.KorName,
		KeyName:         brDao.KeyName,
		LogoImgURL:      brDao.LogoImgUrl,
		OnPopular:       brDao.Onpopular,
		Description:     brDao.Description,
		MaxDiscountRate: brDao.MaxDiscountRate,
		Categories:      cats,
		IsOpen:          brDao.IsOpen,
		InMaintenance:   brDao.InMaintenance,
		NumNewProducts:  brDao.NumNewProductsIn3days,
		SizeGuide:       sizes,
	}
}
