package scripts

import (
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/product"
	"log"
)

func UpdateExhibitionHistory(exId string) {
	input := product.ProductListInput{
		Offset:       0,
		Limit:        1000,
		ExhibitionID: exId,
	}
	ex, err := ioc.Repo.Exhibitions.Get(exId)
	if err != nil {
		log.Println("error occurred on get target exhibition : ", err)
		return
	}

	pds, _, err := product.ListProducts(input)
	if err != nil {
		log.Println("error occurred on listing products", err)
		return
	}

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
	}

}
