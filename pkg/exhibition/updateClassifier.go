package exhibition

import (
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/product"
	"log"
)

// UpdateExhibitionClassifier : 해당 딜(Exhibition)에 해당하는 상품들의 classifier 정보를 해당 딜에 업데이트 해준다.
func UpdateExhibitionClassifier(exhibitionDao *domain.ExhibitionDAO) {
	filter := product.ProductListInput{
		Offset:       0,
		Limit:        10000,
		ExhibitionID: exhibitionDao.ID.Hex(),
	}
	pds, _, err := product.Listing(filter)
	if err != nil {
		log.Println("err occurred on get product list : ", err)
	}

	classifiers := []domain.AlloffClassifier{}
	firstCats := []domain.CategoryClassifier{}
	secondCats := []domain.CategoryClassifier{}

	for _, pd := range pds {
		// clear pds which has no classifier
		if pd.ProductClassifier == nil {
			continue
		}

		// add classifier
		if len(classifiers) == 0 {
			for _, productClassifier := range pd.ProductClassifier.Classifier {
				classifiers = append(classifiers, productClassifier)
			}
		} else {
			for _, productClassifier := range pd.ProductClassifier.Classifier {
				contains := false
				for _, classifier := range classifiers {
					if productClassifier == classifier {
						contains = true
					}
					if !contains {
						classifiers = append(classifiers, productClassifier)
					}
				}
			}
		}

		// add firstCat
		if len(firstCats) == 0 {
			firstCats = append(firstCats, pd.ProductClassifier.First)
		} else {
			firstCatContains := false
			for _, firstCat := range firstCats {
				if pd.ProductClassifier.First.KeyName == firstCat.KeyName {
					firstCatContains = true
				}
			}
			if !firstCatContains {
				firstCats = append(firstCats, pd.ProductClassifier.First)
			}
		}

		// add secondCat
		if len(secondCats) == 0 {
			secondCats = append(secondCats, pd.ProductClassifier.Second)
		} else {
			secondCatContains := false
			for _, secondCat := range secondCats {
				if pd.ProductClassifier.Second.KeyName == secondCat.KeyName {
					secondCatContains = true
				}
			}
			if !secondCatContains {
				secondCats = append(secondCats, pd.ProductClassifier.Second)
			}
		}
	}
	exhibitionClassifier := domain.ExhibitionClassifier{
		Classifier: classifiers,
		First:      firstCats,
		Second:     secondCats,
	}
	exhibitionDao.Classifier = exhibitionClassifier
	_, err = ioc.Repo.Exhibitions.Upsert(exhibitionDao)
	if err != nil {
		log.Println("err occurred on upsert exhibition : ", err)
	}
}
