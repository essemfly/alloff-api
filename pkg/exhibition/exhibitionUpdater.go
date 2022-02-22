package exhibition

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func FindExhibitionInProductGroup(productGroupID string) (*domain.ExhibitionDAO, error) {
	offset, limit := 0, 1000 // Dev code
	onlyLive := false
	exhibitions, _, err := ioc.Repo.Exhibitions.List(offset, limit, onlyLive)
	if err != nil {
		log.Println("Exhibition err in find exhibition from group", err)
	}

	for _, ex := range exhibitions {
		for _, pg := range ex.ProductGroups {
			if pg.ID.Hex() == productGroupID {
				return ex, nil
			}
		}
	}
	return nil, nil
}

func UpdateExhibition(exhibition *domain.ExhibitionDAO) (*domain.ExhibitionDAO, error) {
	newPgs := []*domain.ProductGroupDAO{}
	for _, pg := range exhibition.ProductGroups {
		pgDao, err := ioc.Repo.ProductGroups.Get(pg.ID.Hex())
		if err != nil {
			log.Println("Update exhibition not found pgID: "+pg.ID.Hex(), err)
			continue
		}
		newPgs = append(newPgs, pgDao)
	}

	exhibition.ProductGroups = newPgs
	newExhibition, err := ioc.Repo.Exhibitions.Upsert(exhibition)
	if err != nil {
		log.Println("Upsert error in exhibitionID: "+exhibition.ID.Hex(), err)
	}

	return newExhibition, nil
}
