package productgroup

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

func AddExhibition(req *ExhibitionRequest) (*domain.ExhibitionDAO, error) {
	pgDaos := []*domain.ProductGroupDAO{}

	for _, pgID := range req.ProductGroupIDs {
		pgDao, err := ioc.Repo.ProductGroups.Get(pgID)
		if err != nil {
			log.Println("Add exhibition not found pgID: "+pgID, err)
			continue
		}
		pgDaos = append(pgDaos, pgDao)
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
