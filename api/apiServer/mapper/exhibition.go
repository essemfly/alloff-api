package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapExhibition(exDao *domain.ExhibitionDAO, brief bool) *model.Exhibition {
	return &model.Exhibition{
		ID:             exDao.ID.Hex(),
		ProductTypes:   MapProductTypes(exDao.ProductTypes),
		ExhibitionType: MapExhibitionType(exDao.ExhibitionType),
		Title:          exDao.Title,
		SubTitle:       exDao.SubTitle,
		Description:    exDao.Description,
		Tags:           exDao.Tags,
		BannerImage:    exDao.BannerImage,
		ThumbnailImage: exDao.ThumbnailImage,
		StartTime:      exDao.StartTime.String(),
		FinishTime:     exDao.FinishTime.String(),
		NumAlarms:      exDao.NumAlarms,
		MaxDiscounts:   exDao.MaxDiscounts,
		UserAlarmOn:    false,
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

func MapProductTypes(types []domain.AlloffProductType) []model.AlloffProductType {
	productTypes := []model.AlloffProductType{}

	for _, productType := range types {
		if productType == domain.Male {
			productTypes = append(productTypes, model.AlloffProductTypeMale)
		} else if productType == domain.Kids {
			productTypes = append(productTypes, model.AlloffProductTypeKids)
		} else if productType == domain.Female {
			productTypes = append(productTypes, model.AlloffProductTypeFemale)
		} else if productType == domain.Girl {
			productTypes = append(productTypes, model.AlloffProductTypeGirl)
		} else if productType == domain.Boy {
			productTypes = append(productTypes, model.AlloffProductTypeBoy)
		}
	}

	return productTypes
}
