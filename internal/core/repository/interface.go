package repository

import (
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
	GetByMetaID(metaID, exhibitionID string) (*domain.ProductDAO, error)
	ListByMetaID(metaID string) ([]*domain.ProductDAO, error)
	List(offset, limit int, filter, sortingOptions interface{}) ([]*domain.ProductDAO, int, error)
	ListDistinctBrands(alloffCategoryID string) ([]*domain.BrandDAO, error)
	ListDistinctInfos(filter interface{}) ([]*domain.BrandCountsData, []*domain.AlloffCategoryDAO, []*domain.AlloffSizeDAO)
	ListInfos(filter interface{}) (brands []*domain.BrandCountsData, cats []*domain.AlloffCategoryDAO, sizes []*domain.AlloffSizeDAO)
	Insert(*domain.ProductDAO) (*domain.ProductDAO, error)
	Upsert(*domain.ProductDAO) (*domain.ProductDAO, error)
	CountNewProducts([]string) int
	MakeOutdateProducts([]string, time.Time) int
}

type ProductMetaInfoRepository interface {
	Get(ID string) (*domain.ProductMetaInfoDAO, error)
	GetByProductID(brandKeyname string, productID string) (*domain.ProductMetaInfoDAO, error)
	List(offset, limit int, filter, sortingOptions interface{}) ([]*domain.ProductMetaInfoDAO, int, error)
	Insert(*domain.ProductMetaInfoDAO) (*domain.ProductMetaInfoDAO, error)
	Upsert(*domain.ProductMetaInfoDAO) (*domain.ProductMetaInfoDAO, error)
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

type ProductGroupsRepository interface {
	Get(ID string) (*domain.ProductGroupDAO, error)
	List(offset, limit int, groupType *domain.ProductGroupType, keyword string) ([]*domain.ProductGroupDAO, int, error)
	Upsert(*domain.ProductGroupDAO) (*domain.ProductGroupDAO, error)
}

type ExhibitionsRepository interface {
	Get(ID string) (*domain.ExhibitionDAO, error)
	List(offset, limit int, onlyLive bool, exhibitionStatus domain.ExhibitionStatus, exhibitionType domain.ExhibitionType, query string) ([]*domain.ExhibitionDAO, int, error)
	ListGroupDeals(offset, limit int, onlyLive bool, exhibitionStatus domain.GroupdealStatus) ([]*domain.ExhibitionDAO, int, error)
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
	ListByExhibitionID(exhibitionID string) ([]*domain.OrderItemDAO, error)
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

type RefundItemsRepository interface {
	Insert(*domain.RefundItemDAO) (*domain.RefundItemDAO, error)
}

type NotificationsRepository interface {
	Insert(*domain.NotificationDAO) (*domain.NotificationDAO, error)
	Get(notiID string) ([]*domain.NotificationDAO, error)
	List(offset, limit int, notiTypes []domain.NotificationType, onlyReady bool) ([]*domain.NotificationDAO, error)
	Update(*domain.NotificationDAO) (*domain.NotificationDAO, error)
}

type AlloffSizeRepository interface {
	Get(alloffSizeID string) (*domain.AlloffSizeDAO, error)
	Upsert(dao *domain.AlloffSizeDAO) (*domain.AlloffSizeDAO, error)
	List(offset, limit int) ([]*domain.AlloffSizeDAO, int, error)
}

type CartsRepository interface {
	Get(cartID string) (*domain.Basket, error)
	Upsert(cartDao *domain.Basket) (*domain.Basket, error)
}

type SizeMappingPolicyRepository interface {
	Insert(dao *domain.SizeMappingPolicyDAO) (*domain.SizeMappingPolicyDAO, error)
	Get(id string) (*domain.SizeMappingPolicyDAO, error)
	GetByDetail(size string, productTypes []domain.AlloffProductType, alloffCategoryID string) (*domain.SizeMappingPolicyDAO, error)
	List() ([]*domain.SizeMappingPolicyDAO, error)
	ListByDetail(size string, productTypes []domain.AlloffProductType, alloffCategpryID string) ([]*domain.SizeMappingPolicyDAO, error)
	Update(dao *domain.SizeMappingPolicyDAO) (*domain.SizeMappingPolicyDAO, error)
}
