package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapExhibition(exDao *domain.ExhibitionDAO, brief bool) *model.Exhibition {
	brands := []*model.Brand{}
	pds := []*model.Product{}

	for _, pdDao := range exDao.ChiefProducts {
		pds = append(pds, MapProduct(pdDao))
	}
	for _, brandDao := range exDao.Brands {
		brands = append(brands, MapBrandDaoToBrand(brandDao, false))
	}

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
		Brands:         brands,
		ChiefProducts:  pds,
		NumProducts:    exDao.NumProducts,
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

func MapOrderedAlloffCats(alloffCats []*model.AlloffCategory) []*model.AlloffCategory {
	ordered := [11]*model.AlloffCategory{
		{Name: "아우터"},
		{Name: "상의"},
		{Name: "원피스/세트"},
		{Name: "스커트"},
		{Name: "바지"},
		{Name: "신발"},
		{Name: "가방"},
		{Name: "패션잡화"},
		{Name: "라운지/언더웨어"},
		{Name: "쥬얼리"},
		{Name: "비치웨어"},
	}

	for i, orderedCat := range ordered {
		for _, cat := range alloffCats {
			if cat.Name == orderedCat.Name {
				orderedCat = cat
				ordered[i] = orderedCat
			}
		}
	}

	result := []*model.AlloffCategory{}
	for _, orderedCat := range ordered {
		if orderedCat.ID != "" {
			result = append(result, orderedCat)
		}
	}

	return result
}
