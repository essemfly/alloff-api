package productgroup

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

type ExhibitionRequest struct {
	BannerImage     string
	ThumbnailImage  string
	Title           string
	Description     string
	ProductGroupIDs []string
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
	}

	return exhibition, nil
}
