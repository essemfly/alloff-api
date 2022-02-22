package classifier

import (
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func ClassifyProducts(alloffCategoryID string) *domain.ProductAlloffCategoryDAO {
	alloffCat, err := ioc.Repo.AlloffCategories.Get(alloffCategoryID)
	if err == nil {
		if alloffCat.Level == 1 {
			return &domain.ProductAlloffCategoryDAO{
				First:   alloffCat,
				Second:  nil,
				Done:    true,
				Touched: false,
			}
		} else if alloffCat.Level == 2 {
			parentCat, err := ioc.Repo.AlloffCategories.Get(alloffCat.ParentId.Hex())
			if err != nil {
				return &domain.ProductAlloffCategoryDAO{
					First:   parentCat,
					Second:  alloffCat,
					Done:    true,
					Touched: false,
				}
			}
		}
	}

	return &domain.ProductAlloffCategoryDAO{
		First:   nil,
		Second:  nil,
		Done:    false,
		Touched: false,
	}
}
