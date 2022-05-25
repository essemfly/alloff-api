package productinfo

import (
	"log"

	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/translater"
	"golang.org/x/text/language"
)

func TranslateProductInfo(pdInfo *domain.ProductMetaInfoDAO) (*domain.ProductMetaInfoDAO, error) {
	titleInKorean, err := translater.TranslateText(language.Korean.String(), pdInfo.AlloffName)
	if err != nil {
		log.Println("err", err)
		return nil, err
	}
	informationKorean := map[string]string{}
	for key, value := range pdInfo.SalesInstruction.Information {
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

	pdInfo.AlloffName = titleInKorean
	pdInfo.SalesInstruction.Information = informationKorean
	pdInfo.IsTranslateRequired = false

	newPdInfo, err := Update(pdInfo)
	if err != nil {
		log.Println("err", err)
		return nil, err
	}

	return newPdInfo, nil
}
