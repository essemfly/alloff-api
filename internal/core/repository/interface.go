package repository

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

type BrandsRepository interface {
	Get(keyname string) (*domain.BrandDAO, error)
	List(alloffCategoryID *string) ([]*domain.BrandDAO, error)
	Upsert(*domain.BrandDAO) error
}

type ProductsRepository interface {
	Get(id string) (*domain.ProductDAO, error)
	GetByProductID(productID string) (*domain.ProductDAO, error)
	List(offset, limit int, brandID, categoryID, alloffCategoryID, sorting string, priceRanges []string) ([]*domain.ProductDAO, int, error)
	Upsert(*domain.ProductDAO) error
	ListDistinctBrands(alloffCategoryID string) ([]*domain.BrandDAO, error)
}

type CrawlSourcesRepository interface {
	AddSource(*domain.CrawlSourceDAO) error
	ListSourcesByModule(moduleName string) ([]*domain.CrawlSourceDAO, error)
}
