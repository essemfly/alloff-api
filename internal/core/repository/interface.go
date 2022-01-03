package repository

import "github.com/lessbutter/alloff-api/internal/core/domain"

type BrandsRepository interface {
	Get(ID string) (*domain.BrandDAO, error)
	GetByKeyname(keyname string) (*domain.BrandDAO, error)
	List(limit, offset int, filter, sortingOptions interface{}) ([]*domain.BrandDAO, int, error)
	Upsert(*domain.BrandDAO) (*domain.BrandDAO, error)
}

type ProductsRepository interface {
	Get(ID string) (*domain.ProductDAO, error)
	GetByMetaID(MetaID string) (*domain.ProductDAO, error)
	List(limit, offset int, filter, sortingOptions interface{}) ([]*domain.ProductDAO, int, error)
	Insert(*domain.ProductDAO) (*domain.ProductDAO, error)
	Upsert(*domain.ProductDAO) (*domain.ProductDAO, error)
}

type ProductMetaInfoRepository interface {
	GetByProductID(brandKeyname string, productID string) (*domain.ProductMetaInfoDAO, error)
	Insert(*domain.ProductMetaInfoDAO) (*domain.ProductMetaInfoDAO, error)
	Upsert(*domain.ProductMetaInfoDAO) (*domain.ProductMetaInfoDAO, error)
}

type ProductDiffsRepository interface {
	Insert(*domain.ProductDiffDAO) error
	List(filter interface{}) ([]*domain.ProductDiffDAO, error)
}

type ProductGroupsRepository interface {
}

type CrawlSourcesRepository interface {
	List(filter interface{}) ([]*domain.CrawlSourceDAO, int, error)
	Upsert(*domain.CrawlSourceDAO) (*domain.CrawlSourceDAO, error)
}

type CategoriesRepository interface {
	List(brandKeyname string) ([]*domain.CategoryDAO, error)
	Upsert(*domain.CategoryDAO) (*domain.CategoryDAO, error)
}

type AlloffCategoriesRepository interface {
	Get(ID string) (*domain.AlloffCategoryDAO, error)
	GetByKeyname(keyname string) (*domain.AlloffCategoryDAO, error)
	List(parentID *string) ([]*domain.AlloffCategoryDAO, error)
	Upsert(*domain.AlloffCategoryDAO) (*domain.AlloffCategoryDAO, error)
}

type ClassifyRulesRepository interface {
	Upsert(*domain.ClassifierDAO) (*domain.ClassifierDAO, error)
	GetByKeyname(brandKeyname, categoryKeyname string) (*domain.ClassifierDAO, error)
}

type CrawlRecordsRepository interface {
	GetLast() (*domain.CrawlRecordDAO, error)
	Insert(*domain.CrawlRecordDAO) error
}

type FeaturedsRepository interface {
}

type HomeItemsRepository interface {
}

type OrdersRepository interface {
}

type PaymentsRepository interface {
}

type UsersRepository interface {
}

type DevicesRepository interface {
}

type NotificationsRepository interface {
}

type AlimtalksRepository interface {
}

type LikeBrandsRepository interface {
}

type LikeProductsRepository interface {
}

type AlarmProductGroupsRepository interface {
}
