package scripts

import (
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/dto"
	"github.com/lessbutter/alloff-api/internal/pkg/omnious"
	"log"
)

func AddProductClassifier(pds []*domain.ProductDAO, withOmnious bool) {
	log.Println(":: product classifier add start ::")
	for idx, pd := range pds {
		log.Printf("now add #%v/%v of the products \n", idx+1, len(pds))

		firstCat := domain.CategoryClassifier{
			Name:    pd.AlloffCategories.First.Name,
			KeyName: pd.AlloffCategories.First.KeyName,
		}
		secondCat := domain.CategoryClassifier{}
		omniousTaggingResult := dto.TaggingResult{}
		if withOmnious {
			data, err := omnious.GetOmniousData(pd.ProductInfo.Images[0])
			if err != nil {
				log.Println("err occurred on get omnious data")
			}
			secondCat.Name = data.Category.Name
			secondCat.KeyName = data.Category.ID
			omniousTaggingResult = data.TaggingResult
		}
		classifier := &domain.ProductClassifierDAO{
			Classifier:    []domain.AlloffClassifier{domain.Female}, // 일단 상품들은 여성상품으로
			First:         firstCat,
			Second:        secondCat,
			TaggingResult: omniousTaggingResult,
		}
		pd.ProductClassifier = classifier

		_, err := ioc.Repo.Products.Upsert(pd)
		if err != nil {
			log.Println("err occurred on upsert pd")
		}
	}
	log.Println(":: product classifier end start ::")
}
