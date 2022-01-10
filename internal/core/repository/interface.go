package repository

import (
	"time"

	"github.com/lessbutter/alloff-api/internal/core/domain"
)

type BrandsRepository interface {
	Get(ID string) (*domain.BrandDAO, error)
	GetByKeyname(keyname string) (*domain.BrandDAO, error)
	List(offset, limit int, filter, sortingOptions interface{}) ([]*domain.BrandDAO, int, error)
	Upsert(*domain.BrandDAO) (*domain.BrandDAO, error)
}

type ProductsRepository interface {
	Get(ID string) (*domain.ProductDAO, error)
	GetByMetaID(MetaID string) (*domain.ProductDAO, error)
	List(offset, limit int, filter, sortingOptions interface{}) ([]*domain.ProductDAO, int, error)
	ListDistinctBrands(alloffCategoryID string) ([]*domain.BrandDAO, error)
	Insert(*domain.ProductDAO) (*domain.ProductDAO, error)
	Upsert(*domain.ProductDAO) (*domain.ProductDAO, error)
	CountNewProducts([]string) int
	MakeOutdateProducts([]string, time.Time) int
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
	GetByName(name string) (*domain.AlloffCategoryDAO, error)
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
	Insert(*domain.FeaturedDAO) error
	List() ([]*domain.FeaturedDAO, error)
}

type HomeItemsRepository interface {
	Insert(*domain.HomeItemDAO) error
	Update(*domain.HomeItemDAO) error
	List() ([]*domain.HomeItemDAO, error)
}

type ProductGroupsRepository interface {
	Get(ID string) (*domain.ProductGroupDAO, error)
	List() ([]*domain.ProductGroupDAO, error)
	Upsert(*domain.ProductGroupDAO) (*domain.ProductGroupDAO, error)
}

type ExhibitionsRepository interface {
	Get(ID string) (*domain.ExhibitionDAO, error)
	List() ([]*domain.ExhibitionDAO, error)
	Upsert(*domain.ExhibitionDAO) (*domain.ExhibitionDAO, error)
}

type OrdersRepository interface {
	Get(ID int) (*domain.OrderDAO, error)
	GetByAlloffID(ID string) (*domain.OrderDAO, error)
	List(userID string) ([]*domain.OrderDAO, error)
	Insert(*domain.OrderDAO) (*domain.OrderDAO, error)
	Update(*domain.OrderDAO) (*domain.OrderDAO, error)
}

type PaymentsRepository interface {
	GetByOrderIDAndAmount(orderID string, amount int) (*domain.PaymentDAO, error)
	GetByImpUID(impUID string) (*domain.PaymentDAO, error)
	Insert(*domain.PaymentDAO) (*domain.PaymentDAO, error)
}

type UsersRepository interface {
	Get(ID string) (*domain.UserDAO, error)
	GetByMobile(mobile string) (*domain.UserDAO, error)
	Insert(*domain.UserDAO) (*domain.UserDAO, error)
	Update(*domain.UserDAO) (*domain.UserDAO, error)
}

type DevicesRepository interface {
	GetByDeviceID(deviceID string) (*domain.DeviceDAO, error)
	ListAllowedByUser(userID string) ([]*domain.DeviceDAO, error)
	ListAllowed() ([]*domain.DeviceDAO, error)
	UpdateDevices(deviceID string, allowNotification bool, userID *string) error
}

type AlimtalksRepository interface {
	GetByDetail(userID, templateCode, referenceID string) (*domain.AlimtalkDAO, error)
	Insert(*domain.AlimtalkDAO) (*domain.AlimtalkDAO, error)
	Update(*domain.AlimtalkDAO) (*domain.AlimtalkDAO, error)
	// TO BE MODIFIED
}

type LikeBrandsRepository interface {
	Like(userID, brandID string) (bool, error)
	List(userID string) (*domain.LikeBrandDAO, error)
}

type LikeProductsRepository interface {
	Like(userID, productID string) (bool, error)
	List(userID string) ([]*domain.LikeProductDAO, error)
}

type NotificationsRepository interface {
	// NOT URGENT
}

type AlarmProductGroupsRepository interface {
	// To BE REMOVED
}
