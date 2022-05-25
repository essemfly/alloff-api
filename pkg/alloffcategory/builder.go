package alloffcategory

import (
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func BuildProductAlloffCategory(alloffCategoryID string, touched bool) (*domain.ProductAlloffCategoryDAO, error) {
	alloffCat, err := ioc.Repo.AlloffCategories.Get(alloffCategoryID)
	if err != nil {
		return &domain.ProductAlloffCategoryDAO{}, err
	}

	if alloffCat.Level == 1 {
		return &domain.ProductAlloffCategoryDAO{
			First:   alloffCat,
			Second:  nil,
			Done:    true,
			Touched: touched,
		}, nil
	}

	parentCat, err := ioc.Repo.AlloffCategories.Get(alloffCat.ParentId.Hex())
	if err != nil {
		return nil, err
	}

	return &domain.ProductAlloffCategoryDAO{
		First:   parentCat,
		Second:  alloffCat,
		Done:    true,
		Touched: touched,
	}, nil
}
