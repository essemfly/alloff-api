package scripts

import (
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/dto"
	"github.com/lessbutter/alloff-api/internal/pkg/omnious"
	"log"
)

func AddProductClassifier(pds []*domain.ProductDAO) {
	log.Println(":: product classifier add start ::")
	for idx, pd := range pds {
		log.Printf("now add #%v/%v of the products \n", idx+1, len(pds))

		firstCat := domain.CategoryClassifier{}
		secondCat := domain.CategoryClassifier{}
		omniousTaggingResult := dto.TaggingResult{}
		image := ""

		// if product has no cached image
		if len(pd.Images) == 0 {
			if len(pd.ProductInfo.Images) != 0 {
				image = pd.ProductInfo.Images[0]
			} else if len(pd.SalesInstruction.Description.Images) != 0 {
				image = pd.SalesInstruction.Description.Images[0]
			} else {
				continue
			}
		} else {
			image = pd.Images[0]
		}

		// get omnious data
		data, err := omnious.GetOmniousData(image)
		if err != nil {
			log.Printf("[%s] err occurred on get omnious data : %v\n", pd.ID.Hex(), err)
			continue
		}

		// if product has no alloff category (래밸 1 카테고리만 살리고 first가 lv2 인 카테고리는 버린다.)
		if pd.AlloffCategories.First != nil && pd.AlloffCategories.First.Level == 1 {
			firstCat = domain.CategoryClassifier{
				Name:    pd.AlloffCategories.First.Name,
				KeyName: pd.AlloffCategories.First.KeyName,
			}
		} else {
			cat := omnious.MapOmniousCategoryToCategoryClassifier(data.Category.Name)
			firstCat = *cat
		}

		// get second category from omnious data
		secondCat.Name = data.Category.Name
		secondCat.KeyName = data.Category.ID
		omniousTaggingResult = data.TaggingResult

		classifier := &domain.ProductClassifierDAO{
			Classifier:    []domain.AlloffClassifier{domain.Female}, // 일단 상품들은 여성상품으로
			First:         firstCat,
			Second:        secondCat,
			TaggingResult: omniousTaggingResult,
		}
		pd.ProductClassifier = classifier
		pd.IsProductClassified = true

		_, err = ioc.Repo.Products.Upsert(pd)
		if err != nil {
			log.Println("err occurred on upsert pd")
			continue
		}
	}
	log.Println(":: product classifier end ::")
}
