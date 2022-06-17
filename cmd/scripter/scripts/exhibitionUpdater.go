package scripts

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/product"
	"go.uber.org/zap"
)

func ExhibitionStatusUpdater() {
	offset, limit := 0, 100 // Set 100 for live exhibitions
	onlyLive := true
	exDaos, cnt, err := ioc.Repo.Exhibitions.List(offset, limit, onlyLive, domain.EXHIBITION_LIVE, domain.EXHIBITION_TIMEDEAL, "")
	if err != nil {
		config.Logger.Error("err on listing exhibitions", zap.Error(err))
	}
	log.Println("total exhibitions", cnt)
	for _, exDao := range exDaos {
		// Exhibition Syncer를 태울까 했는데, Exhibition Syncer에는 close하는 것은 안들어가 있다.
		if exDao.FinishTime.Before(time.Now()) {
			exDao.IsLive = false
			ioc.Repo.Exhibitions.Upsert(exDao)
			productListInput := product.ProductListInput{
				Offset:       0,
				Limit:        1000,
				ExhibitionID: exDao.ID.Hex(),
			}
			pds, cnt, err := product.ListProducts(productListInput)
			log.Println("total products# to removed", exDao.ID, cnt)
			if err != nil {
				config.Logger.Error("exhibition syncer error", zap.Error(err))
			}
			for _, pd := range pds {
				pd.IsNotSale = true
				_, err := ioc.Repo.Products.Upsert(pd)
				if err != nil {
					config.Logger.Error("exhibition syncer error", zap.Error(err))
				}
			}
		}
	}
}
