package domain

import (
	"time"

	"github.com/lessbutter/alloff-api/api/server/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HomeItemDAO struct {
	ID             primitive.ObjectID      `bson:"_id,omitempty"`
	Priority       int                     `json:"priority"`
	Title          string                  `json:"title"`
	ItemType       model.HomeItemType      `json:"itemType"`
	TargetID       string                  `json:"target"`
	Sorting        []model.SortingType     `json:"sorting"`
	Images         []string                `json:"images"`
	CommunityItems []*HomeCommunityItemDAO `json:"communityItems"`
	Brands         []*HomeBrandItemDAO     `json:"brands"`
	Products       []*ProductDAO           `json:"products"`
	ProductGroups  []*ProductGroupDAO
	Removed        bool
}

func (item *HomeItemDAO) ToDTO() *model.HomeItem {
	comItems := []*model.CommunityItem{}
	for _, comItem := range item.CommunityItems {
		comItems = append(comItems, comItem.ToDTO())
	}

	brandItems := []*model.BrandItem{}
	for _, brandItem := range item.Brands {
		brandItems = append(brandItems, brandItem.ToDTO())
	}

	pdItems := []*model.Product{}
	for _, pdItem := range item.Products {
		pdItems = append(pdItems, pdItem.ToDTO())
	}

	grouopItems := []*model.ProductGroup{}
	for _, groupItem := range item.ProductGroups {
		grouopItems = append(grouopItems, groupItem.ToDTO())
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

type HomeBrandItemDAO struct {
	ImgUrl string
	Brand  *BrandDAO
}

func (item *HomeBrandItemDAO) ToDTO() *model.BrandItem {
	return &model.BrandItem{
		ImgURL: item.ImgUrl,
		Brand:  item.Brand.ToDTO(false),
	}
}

type HomeCommunityItemDAO struct {
	Name       string                  `json:"name"`
	Target     string                  `json:"target"`
	TargetType model.CommunityItemType `json:"targetType"`
	ImgURL     string                  `json:"imgUrl"`
}

func (item *HomeCommunityItemDAO) ToDTO() *model.CommunityItem {
	return &model.CommunityItem{
		Name:       item.Name,
		Target:     item.Target,
		TargetType: item.TargetType,
		ImgURL:     item.ImgURL,
	}
}

type FeaturedDAO struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	IdentifyName string             `json:"identifyname"`
	Order        int                `json:"order"`
	Brand        *BrandDAO          `json:"brand"`
	Img          string             `json:"img"`
	Category     *CategoryDAO       `json:"category"`
	EndDate      time.Time
}

func (featured *FeaturedDAO) ToDTO() *model.FeaturedItem {
	newFeature := model.FeaturedItem{
		ID:    featured.ID.Hex(),
		Order: featured.Order,
		Img:   featured.Img,
	}

	if featured.Brand != nil {
		newFeature.Brand = featured.Brand.ToDTO(false)
	}

	if featured.Category != nil {
		newFeature.Category = featured.Category.ToDTO()
	}

	return &newFeature
}
