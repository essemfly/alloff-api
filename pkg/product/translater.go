package product

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/translater"
	"golang.org/x/text/language"
)

func TranslateProductInfo(worker chan bool, done chan bool, pd *domain.ProductDAO) (*domain.ProductDAO, error) {
	titleInKorean, err := translater.TranslateText(language.Korean.String(), pd.AlloffName)
	if err != nil {
		log.Println("err", err)
		return nil, err
	}
	informationKorean := map[string]string{}
	for key, value := range pd.ProductInfo.Information {
		keyKorean, err := translater.TranslateText(language.Korean.String(), key)
		if err != nil {
			log.Println("info translate key err", err)
			return nil, err
		}
		valueKorean, err := translater.TranslateText(language.Korean.String(), value)
		if err != nil {
			log.Println("info translate value err", err)
			return nil, err
		}
		informationKorean[keyKorean] = valueKorean
	}

	// inventoryKorean := []domain.InventoryDAO{}
	// for _, inv := range pd.Inventory {
	// 	sizeKorean, err := translater.TranslateText(language.Korean.String(), inv.Size)
	// 	if err != nil {
	// 		log.Println("inventory korean err", err)
	// 	}
	// 	inventoryKorean = append(inventoryKorean, domain.InventoryDAO{
	// 		Size:     sizeKorean,
	// 		Quantity: inv.Quantity,
	// 	})
	// }

	pd.AlloffName = titleInKorean
	pd.ProductInfo.Information = informationKorean
	pd.IsTranslateRequired = false
	// pd.Inventory = inventoryKorean
	newPd, err := ioc.Repo.Products.Upsert(pd)
	if err != nil {
		log.Println("err", err)
		return nil, err
	}

	<-worker
	done <- true

	return newPd, nil
}
