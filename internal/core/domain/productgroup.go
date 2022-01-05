package domain

import (
	"time"

	"github.com/lessbutter/alloff-api/api/server/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductGroupDAO struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	Title       string             `json:"title"`
	ShortTitle  string             `json:"shorttitle"`
	Instruction []string           `json:"instruction"`
	ImgUrl      string             `json:"imgurl"`
	NumAlarms   int
	Products    []*ProductPriorityDAO
	StartTime   time.Time
	FinishTime  time.Time
	Created     time.Time
	Updated     time.Time
}

func (pgDao *ProductGroupDAO) ToDTO() *model.ProductGroup {
	pds := []*model.Product{}
	for _, pd := range pgDao.Products {
		pds = append(pds, pd.Product.ToDTO())
	}

	pg := &model.ProductGroup{
		ID:          pgDao.ID.Hex(),
		Title:       pgDao.Title,
		ShortTitle:  pgDao.ShortTitle,
		Instruction: pgDao.Instruction,
		ImgURL:      pgDao.ImgUrl,
		NumAlarms:   pgDao.NumAlarms,
		Products:    pds,
		StartTime:   pgDao.StartTime.String(),
		FinishTime:  pgDao.FinishTime.String(),
		SetAlarm:    false,
	}

	return pg
}

type ProductPriorityDAO struct {
	Priority int
	Product  *ProductDAO
}

type ExhibitionDAO struct {
	ID             primitive.ObjectID `bons:"_id, omitempty"`
	BannerImage    string
	ThumbnailImage string
	Title          string
	ShortTitle     string
	ProductGroups  []*ProductGroupDAO
}

func (exDao *ExhibitionDAO) ToDTO() *model.Exhibition {
	pgs := []*model.ProductGroup{}

	for _, pg := range exDao.ProductGroups {
		pgs = append(pgs, pg.ToDTO())
	}

	return &model.Exhibition{
		ID:             exDao.ID.Hex(),
		BannerImage:    exDao.BannerImage,
		ThumbnailImage: exDao.ThumbnailImage,
		Title:          exDao.Title,
		ShortTitle:     exDao.ShortTitle,
		ProductGroups:  pgs,
	}
}
