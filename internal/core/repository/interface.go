package repository

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

type BrandsRepository interface {
	Get(ID string) (*domain.BrandDAO, error)
	GetByKeyname(keyname string) (*domain.BrandDAO, error)
	List(limit, offset int, filter, sortingOptions interface{}) ([]*domain.BrandDAO, int, error)
	Upsert(*domain.BrandDAO) (*domain.BrandDAO, error)
}

type ProductsRepository interface {
	Get(ID string) (*domain.ProductDAO, error)
	GetByProductID(brandKeyname string, productID string) (*domain.ProductDAO, error)
	List(limit, offset int, filter, sortingOptions interface{}) ([]*domain.ProductDAO, int, error)
	Upsert(product *domain.ProductDAO) (*domain.ProductDAO, error)
}

type CrawlSourcesRepository interface {
	AddSource(*domain.CrawlSourceDAO) error
	ListSourcesByModule(moduleName string) ([]*domain.CrawlSourceDAO, error)
}
