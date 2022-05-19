package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapReference(ref *domain.ReferenceTarget) *model.ItemReference {
	return &model.ItemReference{
		Path:    ref.Path,
		Params:  ref.Params,
		Options: ref.Options,
	}
}

func MapHomeTabItem(item *domain.HomeTabItemDAO) *model.HomeTabItem {
	pds := []*model.Product{}
	brands := []*model.Brand{}
	exhibitions := []*model.Exhibition{}

	for _, pdDao := range item.Products {
		pds = append(pds, MapProductDaoToProduct(pdDao))
	}

	for _, brandDao := range item.Brands {
		brands = append(brands, MapBrandDaoToBrand(brandDao, false))
	}

	for _, exhibitionDao := range item.Exhibitions {
		exhibitions = append(exhibitions, MapExhibition(exhibitionDao, true))
	}

	return &model.HomeTabItem{
		ID:           item.ID.Hex(),
		Title:        item.Title,
		Description:  item.Description,
		Tags:         item.Tags,
		BackImageURL: item.BackImageUrl,
		ItemType:     model.HomeTabItemTypeEnum(item.Type),
		Products:     pds,
		Brands:       brands,
		Exhibitions:  exhibitions,
		Reference:    MapReference(item.Reference),
	}
}

func MapTopBanner(banner *domain.TopBannerDAO) *model.TopBanner {
	return &model.TopBanner{
		ID:             banner.ID.Hex(),
		ImageURL:       banner.ImageUrl,
		ExhibitionID:   banner.ExhibitionID,
		ExhibitionType: MapExhibitionType(domain.ExhibitionType(banner.ExhibitionType)),
		Title:          banner.Title,
		SubTitle:       banner.SubTitle,
	}
}
