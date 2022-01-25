package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapHomeitem(item *domain.HomeItemDAO) *model.HomeItem {
	comItems := []*model.CommunityItem{}
	for _, comItem := range item.CommunityItems {
		comItems = append(comItems, MapHomeCommItem(comItem))
	}

	brandItems := []*model.BrandItem{}
	for _, brandItem := range item.Brands {
		brandItems = append(brandItems, MapHomeBrandItem(brandItem))
	}

	pdItems := []*model.Product{}
	for _, pdItem := range item.Products {
		pdItems = append(pdItems, MapProductDaoToProduct(pdItem))
	}

	grouopItems := []*model.ProductGroup{}
	for _, groupItem := range item.ProductGroups {
		grouopItems = append(grouopItems, MapProductGroupDao(groupItem))
	}

	return &model.HomeItem{
		ID:             item.ID.Hex(),
		Priority:       item.Priority,
		Title:          item.Title,
		ItemType:       item.ItemType,
		TargetID:       item.TargetID,
		Sorting:        item.Sorting,
		Images:         item.Images,
		CommunityItems: comItems,
		Brands:         brandItems,
		Products:       pdItems,
		ProductGroups:  grouopItems,
	}
}

func MapHomeBrandItem(item *domain.HomeBrandItemDAO) *model.BrandItem {
	return &model.BrandItem{
		ImgURL: item.ImgUrl,
		Brand:  MapBrandDaoToBrand(item.Brand, false),
	}
}

func MapHomeCommItem(item *domain.HomeCommunityItemDAO) *model.CommunityItem {
	return &model.CommunityItem{
		Name:       item.Name,
		Target:     item.Target,
		TargetType: item.TargetType,
		ImgURL:     item.ImgURL,
	}
}

func MapFeatured(featured *domain.FeaturedDAO) *model.FeaturedItem {
	newFeature := model.FeaturedItem{
		ID:    featured.ID.Hex(),
		Order: featured.Order,
		Img:   featured.Img,
	}

	if featured.Brand != nil {
		newFeature.Brand = MapBrandDaoToBrand(featured.Brand, false)
	}

	if featured.Category != nil {
		newFeature.Category = MapCategoryDaoToCategory(featured.Category)
	}

	return &newFeature
}
