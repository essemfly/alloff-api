package exhibition

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type ExhibitionRequest struct {
	BannerImage     string
	ThumbnailImage  string
	Title           string
	SubTitle        string
	Description     string
	Tags            []string
	ProductGroupIDs []string
	ExhibitionType  domain.ExhibitionType
	StartTime       time.Time
	FinishTime      time.Time
}

// 현재 Mock에서만 쓰이고있네.
func AddExhibition(req *ExhibitionRequest) (*domain.ExhibitionDAO, error) {
	exDao := &domain.ExhibitionDAO{
		ID:             primitive.NewObjectID(),
		BannerImage:    req.BannerImage,
		ThumbnailImage: req.ThumbnailImage,
		Title:          req.Title,
		SubTitle:       req.SubTitle,
		Description:    req.Description,
		Tags:           req.Tags,
		StartTime:      req.StartTime,
		FinishTime:     req.FinishTime,
		IsLive:         false,
		ExhibitionType: req.ExhibitionType,
		CreatedAt:      time.Now(),
	}

	pgs := []*domain.ProductGroupDAO{}
	for _, pgID := range req.ProductGroupIDs {
		pg, err := ioc.Repo.ProductGroups.Get(pgID)
		if err != nil {
			log.Println("get product group failed: "+pgID, err)
			continue
		}
		productGroupType := pg.GroupType
		pg.StartTime = req.StartTime
		pg.FinishTime = req.FinishTime
		if pg.Brand != nil {
			productGroupType = domain.PRODUCT_GROUP_BRAND_TIMEDEAL
		}
		pg.GroupType = productGroupType
		pg.ExhibitionID = exDao.ID.Hex()
		newPg, err := ioc.Repo.ProductGroups.Upsert(pg)
		if err != nil {
			log.Println("update product group failed: "+pgID, err)
		}
		pgs = append(pgs, newPg)
	}

	exDao.ProductGroups = pgs

	newExDao, err := ioc.Repo.Exhibitions.Upsert(exDao)
	if err != nil {
		config.Logger.Error("Exhibition create error", zap.Error(err))
	}

	return newExDao, err
}
