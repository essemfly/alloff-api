package crawler

import "github.com/lessbutter/alloff-api/internal/core/domain"

func GetAlloffCategory(pd *domain.ProductDAO) *domain.ProductAlloffCategoryDAO {
	return &domain.ProductAlloffCategoryDAO{
		First:   &domain.AlloffCategoryDAO{},
		Second:  &domain.AlloffCategoryDAO{},
		Done:    false,
		Touched: false,
	}
}
