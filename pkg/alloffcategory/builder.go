package alloffcategory

import (
	"errors"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func BuildProductAlloffCategory(alloffCat *domain.AlloffCategoryDAO, touched bool) (*domain.ProductAlloffCategoryDAO, error) {
	if alloffCat == nil {
		return nil, errors.New("cannot find alloff category nil")
	}

	if alloffCat.Level == 1 {
		return &domain.ProductAlloffCategoryDAO{
			TaggingResults: &domain.TaggingResultDAO{},
			First:          alloffCat,
			Second:         nil,
			Done:           true,
			Touched:        touched,
		}, nil
	}

	parentCat, err := ioc.Repo.AlloffCategories.Get(alloffCat.ParentId.Hex())
	if err != nil {
		return nil, err
	}

	return &domain.ProductAlloffCategoryDAO{
		TaggingResults: &domain.TaggingResultDAO{},
		First:          parentCat,
		Second:         alloffCat,
		Done:           true,
		Touched:        touched,
	}, nil
}
