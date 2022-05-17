package repository

import (
	"github.com/lessbutter/alloff-api/internal/core/dto"
	"time"

	"github.com/lessbutter/alloff-api/internal/core/domain"
)

type BrandsRepository interface {
	Get(ID string) (*domain.BrandDAO, error)
	GetByKeyname(keyname string) (*domain.BrandDAO, error)
	List(offset, limit int, onlyPopular, excludeHide bool, sortingOptions interface{}) ([]*domain.BrandDAO, int, error)
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
	Update(*domain.ProductDiffDAO) (*domain.ProductDiffDAO, error)
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
	List(offset, limit int, groupType *domain.ProductGroupType, keyword string) ([]*domain.ProductGroupDAO, int, error)
	Upsert(*domain.ProductGroupDAO) (*domain.ProductGroupDAO, error)
}

type ExhibitionsRepository interface {
	Get(ID string) (*domain.ExhibitionDAO, error)
	List(offset, limit int, onlyLive bool, exhibitionType domain.ExhibitionType, query string) ([]*domain.ExhibitionDAO, int, error)
	Upsert(*domain.ExhibitionDAO) (*domain.ExhibitionDAO, error)
}

type OrdersRepository interface {
	Get(ID int) (*domain.OrderDAO, error)
	GetByAlloffID(ID string) (*domain.OrderDAO, error)
	ListAllPaid() ([]*domain.OrderDAO, error)
	List(userID string, onlyPaid bool) ([]*domain.OrderDAO, error)
	Insert(*domain.OrderDAO) (*domain.OrderDAO, error)
	Update(*domain.OrderDAO) (*domain.OrderDAO, error)
}

type OrderItemsRepository interface {
	Get(ID int) (*domain.OrderItemDAO, error)
	GetByCode(code string) (*domain.OrderItemDAO, error)
	ListByProductGroupID(pgID string) ([]*domain.OrderItemDAO, int, error)
	ListByOrderID(orderID int) ([]*domain.OrderItemDAO, error)
	ListAllCanceled() ([]*domain.OrderItemDAO, error)
	Update(*domain.OrderItemDAO) (*domain.OrderItemDAO, error)
}
type PaymentsRepository interface {
	GetByOrderIDAndAmount(orderID string, amount int) (*domain.PaymentDAO, error)
	GetByImpUID(impUID string) (*domain.PaymentDAO, error)
	ListHolding() ([]*domain.PaymentDAO, error)
	Insert(*domain.PaymentDAO) (*domain.PaymentDAO, error)
	Update(*domain.PaymentDAO) (*domain.PaymentDAO, error)
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
	MakeRemoved(deviceID string) error
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
	ListProductsLike(productId string) ([]*domain.LikeProductDAO, error)
	Update(*domain.LikeProductDAO) (*domain.LikeProductDAO, error)
}

type RefundItemsRepository interface {
	Insert(*domain.RefundItemDAO) (*domain.RefundItemDAO, error)
}

type NotificationsRepository interface {
	Insert(*domain.NotificationDAO) (*domain.NotificationDAO, error)
	Get(notiID string) ([]*domain.NotificationDAO, error)
	List(offset, limit int, notiTypes []domain.NotificationType, onlyReady bool) ([]*domain.NotificationDAO, error)
	Update(*domain.NotificationDAO) (*domain.NotificationDAO, error)
}

type HomeTabItemsRepository interface {
	Insert(*domain.HomeTabItemDAO) (*domain.HomeTabItemDAO, error)
	Get(itemID string) (*domain.HomeTabItemDAO, error)
	List(offset, limit int, onlyLive bool) ([]*domain.HomeTabItemDAO, int, error)
	Update(*domain.HomeTabItemDAO) (*domain.HomeTabItemDAO, error)
}

type TopBannersRepository interface {
	Insert(*domain.TopBannerDAO) (*domain.TopBannerDAO, error)
	Get(itemID string) (*domain.TopBannerDAO, error)
	List(offset, limit int, onlyLive bool) ([]*domain.TopBannerDAO, int, error)
	Update(*domain.TopBannerDAO) (*domain.TopBannerDAO, error)
}

type BestProductsRepository interface {
	Insert(*domain.BestProductDAO) (*domain.BestProductDAO, error)
	GetLatest(alloffCategoryID string) (*domain.BestProductDAO, error)
}

type BestBrandRepository interface {
	Insert(dao *domain.BestBrandDAO) (*domain.BestBrandDAO, error)
	GetLatest() (*domain.BestBrandDAO, error)
}

type OrderCountsRepository interface {
	Get(exhibitionID string) (int, error)
	Push(exhibitionID string) (int, error)
	Cancel(exhibitionID string) (int, error)
}

type AlloffSizeRepository interface {
	Get(alloffSizeID string) (*domain.AlloffSizeDAO, error)
	Upsert(dao *domain.AlloffSizeDAO) (*domain.AlloffSizeDAO, error)
	List(offset, limit int) ([]*domain.AlloffSizeDAO, int, error)
}

type AccessLogRepository interface {
	Index(*domain.AccessLogDAO) (int, error)
	List(limit int, from, to time.Time, order string) (*dto.AccessLogDTO, error)
	GetLatest(limit int) (*dto.AccessLogDTO, error)
}

type ProductLogRepository interface {
	Index(*domain.ProductDAO, domain.LogType) (int, error)
	GetRank(limit int, from time.Time, to time.Time) (*dto.DocumentCountDTO, error)
	GetRankByCatId(limit int, from time.Time, to time.Time, catId string) (*dto.DocumentCountDTO, error)
}

type BrandLogRepository interface {
	Index(*domain.BrandDAO) (int, error)
	GetRank(limit int, from time.Time, to time.Time) (*dto.DocumentCountDTO, error)
}

type SearchLogRepository interface {
	Index(string) (int, error)
}
