package ioc

import (
	"github.com/lessbutter/alloff-api/internal/core/repository"
)

type iocRepo struct {
	Brands             repository.BrandsRepository
	Products           repository.ProductsRepository
	CrawlSources       repository.CrawlSourcesRepository
	CrawlRecords       repository.CrawlRecordsRepository
	Categories         repository.CategoriesRepository
	AlloffCategories   repository.AlloffCategoriesRepository
	ClassifyRules      repository.ClassifyRulesRepository
	ProductDiffs       repository.ProductDiffsRepository
	ProductGroups      repository.ProductGroupsRepository
	Featureds          repository.FeaturedsRepository
	HomeItems          repository.HomeItemsRepository
	Orders             repository.OrdersRepository
	Payments           repository.PaymentsRepository
	Users              repository.UsersRepository
	Devices            repository.DevicesRepository
	Notifications      repository.NotificationsRepository
	Alimtalks          repository.AlimtalksRepository
	LikeBrands         repository.LikeBrandsRepository
	LikeProducts       repository.LikeProductsRepository
	AlarmProductGroups repository.AlarmProductGroupsRepository
}

var Repo iocRepo
