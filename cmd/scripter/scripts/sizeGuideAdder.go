package scripts

import (
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/broker"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/product"
	"go.uber.org/zap"
)

func AddSizeGuide(sizeGuide []domain.SizeGuideDAO) {
	exhibitions, _, _ := ioc.Repo.Exhibitions.List(0, 1000, false, domain.EXHIBITION_STATUS_ALL, domain.EXHIBITION_TIMEDEAL, "")

	liveBrands := []string{}
	for _, exhibition := range exhibitions {
		productListInput := product.ProductListInput{
			Offset:       0,
			Limit:        1000,
			ExhibitionID: exhibition.ID.Hex(),
		}

		pds, _, err := product.ListProducts(productListInput)
		if err != nil {
			config.Logger.Error("error on get products list of product groups : ", zap.Error(err))
		}

		for _, pd := range pds {
			liveBrands = append(liveBrands, pd.ProductInfo.Brand.KeyName)
		}
	}

	liveBrands = utils.RemoveDuplicateString(liveBrands)
	for _, bdKey := range liveBrands {
		bdDao, err := ioc.Repo.Brands.GetByKeyname(bdKey)
		if err != nil {
			config.Logger.Error("error occurred on get brand by keyname : ", zap.Error(err))
		}
		bdDao.SizeGuide = sizeGuide
		bdDao, err = ioc.Repo.Brands.Upsert(bdDao)
		if err != nil {
			config.Logger.Error("error occurred on upsert brand : ", zap.Error(err))
		}
		broker.BrandSyncer(bdDao.KeyName)
	}
}
