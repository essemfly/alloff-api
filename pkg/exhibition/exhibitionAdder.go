package exhibition

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

type ExhibitionRequest struct {
	BannerImage     string
	ThumbnailImage  string
	Title           string
	Description     string
	ProductGroupIDs []string
	StartTime       time.Time
	FinishTime      time.Time
}

// 현재 Mock에서만 쓰이고있네.
func AddExhibition(req *ExhibitionRequest) (*domain.ExhibitionDAO, error) {
	pgDaos := []*domain.ProductGroupDAO{}

	for _, pgID := range req.ProductGroupIDs {
		pgDao, err := ioc.Repo.ProductGroups.Get(pgID)
		if err != nil {
			log.Println("Add exhibition not found pgID: "+pgID, err)
			continue
		}
		pgDao.StartTime = req.StartTime
		pgDao.FinishTime = req.FinishTime
		newPgDao, err := ioc.Repo.ProductGroups.Upsert(pgDao)
		if err != nil {
			log.Println("update product group failed: "+pgID, err)
		}
		pgDaos = append(pgDaos, newPgDao)
	}

	exhibition := &domain.ExhibitionDAO{
		BannerImage:    req.BannerImage,
		ThumbnailImage: req.ThumbnailImage,
		Title:          req.Title,
		Description:    req.Description,
		ProductGroups:  pgDaos,
		StartTime:      req.StartTime,
		FinishTime:     req.FinishTime,
	}

	return exhibition, nil
}
