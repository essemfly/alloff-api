package domain

import (
	"time"

	"github.com/lessbutter/alloff-api/api/server/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BrandDAO struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	KorName               string
	EngName               string
	KeyName               string
	Description           string
	LogoImgUrl            string
	Category              []*CategoryDAO
	SizeGuide             []SizeGuideDAO
	Created               time.Time
	Modulename            string
	Onpopular             bool
	MaxDiscountRate       int
	IsOpen                bool
	IsHide                bool
	InMaintenance         bool
	NumNewProductsIn3days int
}

type SizeGuideDAO struct {
	Label  string
	ImgUrl string
}

type LikeBrandDAO struct {
	Userid  string
	Brands  []*BrandDAO
	Created time.Time
	Updated time.Time
}

func (brDao *BrandDAO) ToDTO(includeCategory bool) *model.Brand {
	var cats []*model.Category
	for _, catDao := range brDao.Category {
		cats = append(cats, catDao.ToDTO())
	}

	sizes := []*model.SizeGuide{}
	for _, guide := range brDao.SizeGuide {
		sizes = append(sizes, &model.SizeGuide{
			Label:  guide.Label,
			ImgURL: guide.ImgUrl,
		})
	}

	return &model.Brand{
		ID:              brDao.ID.Hex(),
		EngName:         brDao.EngName,
		KorName:         brDao.KorName,
		KeyName:         brDao.KeyName,
		LogoImgURL:      brDao.LogoImgUrl,
		OnPopular:       brDao.Onpopular,
		Description:     brDao.Description,
		MaxDiscountRate: brDao.MaxDiscountRate,
		Categories:      cats,
		IsOpen:          brDao.IsOpen,
		InMaintenance:   brDao.InMaintenance,
		NumNewProducts:  brDao.NumNewProductsIn3days,
		SizeGuide:       sizes,
	}
}
