package exhibition

import (
	"log"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/product"
	"go.uber.org/zap"
)

// Exhibition안에 들어있는 섹션들에 대해서, 색션들의 PG상태를 업데이트 시켜줌
// Exhibition이 갖고 있는 PG와 PG들이 갖고 있는 EXID들 사이에서 맞춰줘야하는데
// Ex가 갖고 있는 ProductGROUPS가 우선임
func ExhibitionSyncer(exDao *domain.ExhibitionDAO) (*domain.ExhibitionDAO, error) {
	pdTypes := []domain.AlloffProductType{}
	newPgs := []*domain.ProductGroupDAO{}
	maxDiscountRates := 0

	for _, pg := range exDao.ProductGroups {
		pgDao, err := ioc.Repo.ProductGroups.Get(pg.ID.Hex())
		if err != nil {
			log.Println("Update exhibition not found pgID: "+pg.ID.Hex(), err)
			continue
		}

		pgType := domain.PRODUCT_GROUP_EXHIBITION
		if exDao.ExhibitionType == domain.EXHIBITION_TIMEDEAL {
			pgType = domain.PRODUCT_GROUP_TIMEDEAL
		} else if exDao.ExhibitionType == domain.EXHIBITION_GROUPDEAL {
			pgType = domain.PRODUCT_GROUP_GROUPDEAL
		}

		if pg.Brand != nil {
			pgType = domain.PRODUCT_GROUP_BRAND_TIMEDEAL
		}
		pgDao.StartTime = exDao.StartTime
		pgDao.FinishTime = exDao.FinishTime
		pgDao.ExhibitionID = exDao.ID.Hex()
		pgDao.GroupType = pgType
		updatedPgDao, err := ioc.Repo.ProductGroups.Upsert(pgDao)
		if err != nil {
			log.Println("product group update failed", pgDao.ID.Hex())
		}
		newPgs = append(newPgs, updatedPgDao)

		productListInput := product.ProductListInput{
			Offset:         0,
			Limit:          1000,
			ProductGroupID: pgDao.ID.Hex(),
		}
		pds, _, err := product.ListProducts(productListInput)
		if err != nil {
			config.Logger.Error("exhibition syncer error", zap.Error(err))
		}
		for _, pd := range pds {
			if maxDiscountRates < pd.ProductInfo.Price.DiscountRate {
				maxDiscountRates = pd.ProductInfo.Price.DiscountRate
			}
			pd.ExhibitionID = pgDao.ExhibitionID
			pd.ExhibitionStartTime = pgDao.StartTime
			pd.ExhibitionFinishTime = pgDao.FinishTime
			_, err := ioc.Repo.Products.Upsert(pd)
			if err != nil {
				config.Logger.Error("exhibition syncer error", zap.Error(err))
			}
			pdTypes = append(pdTypes, pd.ProductInfo.ProductType...)
		}
		pdTypes = removeDuplicateType(pdTypes)
	}

	exDao.ProductGroups = newPgs
	exDao.ProductTypes = pdTypes
	exDao.MaxDiscounts = maxDiscountRates

	newExDao, err := ioc.Repo.Exhibitions.Upsert(exDao)
	if err != nil {
		config.Logger.Error("failed in upsert exhibition", zap.Error(err))
	}

	return newExDao, err
}

func ProductGroupSyncer(pgDao *domain.ProductGroupDAO) error {
	productListInput := product.ProductListInput{
		Offset:         0,
		Limit:          1000,
		ProductGroupID: pgDao.ID.Hex(),
	}

	pds, _, err := product.ListProducts(productListInput)
	if err != nil {
		config.Logger.Error("exhibition syncer error", zap.Error(err))
		return err
	}
	for _, pd := range pds {
		pd.ExhibitionID = ""
		pd.ExhibitionFinishTime = pgDao.StartTime
		pd.ExhibitionStartTime = pgDao.StartTime
		_, err := ioc.Repo.Products.Upsert(pd)
		if err != nil {
			config.Logger.Error("exhibition syncer error", zap.Error(err))
			return err
		}
	}
	return nil
}

func removeDuplicateType(strSlice []domain.AlloffProductType) []domain.AlloffProductType {
	allKeys := make(map[domain.AlloffProductType]bool)
	list := []domain.AlloffProductType{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
