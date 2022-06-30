package scripts

import (
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/product"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
	"go.uber.org/zap"
	"log"
)

func UpdateExhibitionHistory() {
	onlyLive := true
	exs, _, err := ioc.Repo.Exhibitions.List(0, 100, &onlyLive, domain.EXHIBITION_STATUS_ALL, domain.EXHIBITION_TIMEDEAL, "")
	if err != nil {
		config.Logger.Error("error on get live exhibitions : ", zap.Error(err))
		return
	}

	for _, ex := range exs {
		exId := ex.ID.Hex()
		input := product.ProductListInput{
			Offset:       0,
			Limit:        1000,
			ExhibitionID: exId,
		}

		pds, cnt, err := product.ListProducts(input)
		if err != nil {
			log.Println("error occurred on listing products", err)
			return
		}

		log.Printf("exhibition : %s has %v pds", exId, cnt)

		for _, pd := range pds {
			pdInfo, err := ioc.Repo.ProductMetaInfos.Get(pd.ProductInfo.ID.Hex())
			if err != nil {
				log.Println("error occurred on get productinfo of product ", pd.ID.Hex(), "error : ", err)
			}
			pdInfo.ExhibitionHistory = &domain.ExhibitionHistoryDAO{
				ExhibitionID: exId,
				Title:        ex.Title,
				StartTime:    ex.StartTime,
				FinishTime:   ex.FinishTime,
			}
			newPd, err := productinfo.Update(pdInfo)
			if newPd.ExhibitionHistory.ExhibitionID != exId {
				log.Println("something wrong with : ", newPd.ID.Hex())
			}
		}
	}
}
