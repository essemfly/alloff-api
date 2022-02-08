package ioc

import (
	"github.com/lessbutter/alloff-api/internal/core/repository"
)

type iocRepo struct {
	Brands           repository.BrandsRepository
	Products         repository.ProductsRepository
	ProductMetaInfos repository.ProductMetaInfoRepository
	ProductDiffs     repository.ProductDiffsRepository
	ProductGroups    repository.ProductGroupsRepository
	CrawlSources     repository.CrawlSourcesRepository
	CrawlRecords     repository.CrawlRecordsRepository
	Categories       repository.CategoriesRepository
	AlloffCategories repository.AlloffCategoriesRepository
	ClassifyRules    repository.ClassifyRulesRepository
	Featureds        repository.FeaturedsRepository
	HomeItems        repository.HomeItemsRepository
	Orders           repository.OrdersRepository
	OrderItems       repository.OrderItemsRepository
	Payments         repository.PaymentsRepository
	Users            repository.UsersRepository
	Devices          repository.DevicesRepository
	Notifications    repository.NotificationsRepository
	Alimtalks        repository.AlimtalksRepository
	LikeBrands       repository.LikeBrandsRepository
	LikeProducts     repository.LikeProductsRepository
	Exhibitions      repository.ExhibitionsRepository
	Refunds          repository.RefundItemsRepository
	HomeTabItems     repository.HomeTabItemsRepository
}

var Repo iocRepo
