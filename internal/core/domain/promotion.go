package domain

import (
	"time"

	"github.com/lessbutter/alloff-api/api/server/model"
)

type HomeItemDAO struct {
	ID             string                  `bson:"_id,omitempty"`
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

type HomeBrandItemDAO struct {
	ImgUrl string
	Brand  *BrandDAO
}

type HomeCommunityItemDAO struct {
	Name       string                  `json:"name"`
	Target     string                  `json:"target"`
	TargetType model.CommunityItemType `json:"targetType"`
	ImgURL     string                  `json:"imgUrl"`
}

type FeaturedDAO struct {
	ID           string       `json:"id"`
	IdentifyName string       `json:"identifyname"`
	Order        int          `json:"order"`
	Brand        *BrandDAO    `json:"brand"`
	Img          string       `json:"img"`
	Category     *CategoryDAO `json:"category"`
	EndDate      time.Time
}
