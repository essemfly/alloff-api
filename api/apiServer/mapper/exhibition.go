package mapper

import (
	"time"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapExhibition(exDao *domain.ExhibitionDAO, brief bool) *model.Exhibition {
	info := MapExhibitionMetaInfo(exDao.MetaInfos)
	return &model.Exhibition{
		ID:             exDao.ID.Hex(),
		ExhibitionType: MapExhibitionType(exDao.ExhibitionType),
		Title:          exDao.Title,
		SubTitle:       exDao.SubTitle,
		Description:    exDao.Description,
		Tags:           exDao.Tags,
		BannerImage:    exDao.BannerImage,
		ThumbnailImage: exDao.ThumbnailImage,
		StartTime:      exDao.StartTime.Add(9 * time.Hour).String(),
		FinishTime:     exDao.FinishTime.Add(9 * time.Hour).String(),
		NumAlarms:      exDao.NumAlarms,
		MetaInfos:      info,
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

func MapExhibitionMetaInfo(info *domain.ExhibitionMetaInfoDAO) *model.ExhibitionInfo {
	brands := []*model.Brand{}
	for _, brandDao := range info.Brands {
		brands = append(brands, MapBrandDaoToBrand(brandDao, false))
	}

	cats := []*model.AlloffCategory{}
	for _, catDao := range info.AlloffCategories {
		cats = append(cats, MapAlloffCatDaoToAlloffCat(catDao))
	}

	invs := []*model.AlloffInventory{}
	for _, invDao := range info.AlloffInventories {
		invs = append(invs, &model.AlloffInventory{
			AlloffSize: &model.AlloffSize{
				SizeName:       invDao.AlloffSize.SizeName,
				AlloffCategory: MapAlloffCatDaoToAlloffCat(invDao.AlloffSize.AlloffCategory),
			},
			Quantity: invDao.Quantity,
		})
	}

	productTypes := MapProductTypes(info.ProductTypes)

	return &model.ExhibitionInfo{
		ProductTypes:      productTypes,
		Brands:            brands,
		AlloffCategories:  cats,
		AlloffInventories: invs,
		MaxDisctounts:     info.MaxDiscounts,
	}
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
		} else if productType == domain.Sports {
			productTypes = append(productTypes, model.AlloffProductTypeSports)
		}
	}

	return productTypes
}
