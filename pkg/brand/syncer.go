package brand

import (
	"log"

	"github.com/lessbutter/alloff-api/config"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
	"go.uber.org/zap"

	"github.com/lessbutter/alloff-api/config/ioc"
)

func BrandSyncer(brandKeyname string) {
	offset, limit := 0, 20000

	newBrand, _ := ioc.Repo.Brands.GetByKeyname(brandKeyname)

	query := productinfo.ProductInfoListInput{
		Offset:  offset,
		Limit:   limit,
		BrandID: newBrand.ID.Hex(),
	}

	products, _, err := productinfo.ListProductInfos(query)
	if err != nil {
		config.Logger.Error("error occurred on listing product infos : ", zap.Error(err))
	}

	for _, pd := range products {
		pd.Brand = newBrand
		_, err := productinfo.Update(pd)
		if err != nil {
			log.Println("err occured", err)
		}
	}
}
