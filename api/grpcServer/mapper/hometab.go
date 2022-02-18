package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func ReferenceMapper(ref *domain.ReferenceTarget) *grpcServer.HomeTabItemReferenceMessage {
	options := []grpcServer.SortingOptions{}
	for _, option := range ref.Options {
		options = append(options, SortingOptionMapper(option))
	}

	return &grpcServer.HomeTabItemReferenceMessage{
		Path:    ref.Path,
		Params:  ref.Params,
		Options: options,
	}
}

func SortingOptionMapper(option model.SortingType) grpcServer.SortingOptions {
	switch option {
	case model.SortingTypeDiscount0_30:
		return grpcServer.SortingOptions_DISCOUNT_0_30
	case model.SortingTypeDiscount30_50:
		return grpcServer.SortingOptions_DISCOUNT_30_50
	case model.SortingTypeDiscount50_70:
		return grpcServer.SortingOptions_DISCOUNT_50_70
	case model.SortingTypeDiscount70_100:
		return grpcServer.SortingOptions_DISCOUNT_50_70
	case model.SortingTypeDiscountrateAscending:
		return grpcServer.SortingOptions_DISCOUNTRATE_ASCENDING
	case model.SortingTypeDiscountrateDescending:
		return grpcServer.SortingOptions_DISCOUNTRATE_DESCENDING
	case model.SortingTypePriceAscending:
		return grpcServer.SortingOptions_DISCOUNTRATE_ASCENDING
	case model.SortingTypePriceDescending:
		return grpcServer.SortingOptions_DISCOUNTRATE_DESCENDING
	}
	return 0
}

func ItemTypeMapper(itemType domain.HomeTabItemType) grpcServer.ItemType {
	switch itemType {
	case domain.HOMETAB_ITEM_BRANDS:
		return grpcServer.ItemType_HOMETAB_ITEM_BRANDS
	case domain.HOMETAB_ITEM_BRAND_EXHIBITION:
		return grpcServer.ItemType_HOMETAB_ITEM_BRAND_EXHIBITION
	case domain.HOMETAB_ITEM_EXHIBITION:
		return grpcServer.ItemType_HOMETAB_ITEM_EXHIBITION
	case domain.HOMETAB_ITEM_EXHIBITIONS:
		return grpcServer.ItemType_HOMETAB_ITEM_EXHIBITIONS
	case domain.HOMETAB_ITEM_PRODUCTS_BRANDS:
		return grpcServer.ItemType_HOMETAB_ITEM_PRODUCTS_BRANDS
	case domain.HOMETAB_ITEM_PRODUCTS_CATEGORIES:
		return grpcServer.ItemType_HOMETAB_ITEM_PRODUCTS_CATEGORIES
	}
	return 0
}

func HomeTabItemMapper(item *domain.HomeTabItemDAO) *grpcServer.HomeTabItemMessage {
	pds := []*grpcServer.ProductMessage{}
	brands := []*grpcServer.BrandMessage{}
	exhibitions := []*grpcServer.ExhibitionMessage{}

	for _, pdDao := range item.Products {
		pds = append(pds, ProductMapper(pdDao))
	}

	for _, brandDao := range item.Brands {
		brands = append(brands, BrandMapper(brandDao))
	}

	for _, exhibitionDao := range item.Exhibitions {
		exhibitions = append(exhibitions, ExhibitionMapper(exhibitionDao))
	}

	return &grpcServer.HomeTabItemMessage{
		ItemId:       item.ID.Hex(),
		Title:        item.Title,
		Description:  item.Description,
		Tags:         item.Tags,
		ItemType:     ItemTypeMapper(item.Type),
		Products:     pds,
		Brands:       brands,
		Exhibitions:  exhibitions,
		BackImageUrl: item.BackImageUrl,
		StartTime:    item.StartedAt.String(),
		FinishTime:   item.FinishedAt.String(),
		Reference:    ReferenceMapper(item.Reference),
	}
}
