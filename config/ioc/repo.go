package ioc

import (
	"github.com/lessbutter/alloff-api/internal/core/repository"
)

type iocRepo struct {
	Brands           repository.BrandsRepository
	Products         repository.ProductsRepository
	ProductMetaInfos repository.ProductMetaInfoRepository
	ProductGroups    repository.ProductGroupsRepository
	CrawlSources     repository.CrawlSourcesRepository
	CrawlRecords     repository.CrawlRecordsRepository
	Categories       repository.CategoriesRepository
	AlloffCategories repository.AlloffCategoriesRepository
	ClassifyRules    repository.ClassifyRulesRepository
	Orders           repository.OrdersRepository
	OrderItems       repository.OrderItemsRepository
	Payments         repository.PaymentsRepository
	Users            repository.UsersRepository
	Devices          repository.DevicesRepository
	Notifications    repository.NotificationsRepository
	Alimtalks        repository.AlimtalksRepository
	LikeBrands       repository.LikeBrandsRepository
	Exhibitions      repository.ExhibitionsRepository
	Refunds          repository.RefundItemsRepository
	OrderCounts      repository.OrderCountsRepository
	AlloffSizes      repository.AlloffSizeRepository
}

var Repo iocRepo
