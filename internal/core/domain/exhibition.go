package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExhibitionType string

const (
	EXHIBITION_TIMEDEAL  ExhibitionType = "TIMEDEAL"
	EXHIBITION_NORMAL    ExhibitionType = "NORMAL"
	EXHIBITION_GROUPDEAL ExhibitionType = "GROUPDEAL"
	EXHIBITION_ALL       ExhibitionType = "ALL"
)

type ExhibitionStatus string

const (
	EXHIBITION_LIVE       ExhibitionStatus = "LIVE"
	EXHIBITION_NOTOPEN    ExhibitionStatus = "NOT_OPEN"
	EXHIBITION_CLOSED     ExhibitionStatus = "CLOSED"
	EXHIBITION_STATUS_ALL ExhibitionStatus = "ALL"
)

type GroupdealStatus string

const (
	GROUPDEAL_PENDING GroupdealStatus = "PENDING"
	GROUPDEAL_OPEN    GroupdealStatus = "OPEN"
	GROUPDEAL_CLOSED  GroupdealStatus = "CLOSED"
)

type ExhibitionBanner struct {
	ImgUrl         string
	Title          string
	Subtitle       string
	ProductGroupId string
}

type ExhibitionDAO struct {
	ID             primitive.ObjectID `bson:"_id, omitempty"`
	ProductTypes   []AlloffProductType
	ExhibitionType ExhibitionType
	Title          string
	SubTitle       string
	Description    string
	Tags           []string
	BannerImage    string
	ThumbnailImage string
	ProductGroups  []*ProductGroupDAO
	StartTime      time.Time
	FinishTime     time.Time
	IsLive         bool
	NumAlarms      int
	MaxDiscounts   int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (exDao *ExhibitionDAO) IsSales() bool {
	now := time.Now()
	if now.After(exDao.StartTime) && now.Before(exDao.FinishTime) {
		return true
	}
	return false
}
