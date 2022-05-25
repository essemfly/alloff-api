package classifier

import (
	"fmt"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/omnious"
	"go.uber.org/zap"
	"log"
	"strings"

	"github.com/lessbutter/alloff-api/config/ioc"
)

func SetProductAlloffCategory(pdInfo *domain.ProductMetaInfoDAO) *domain.ProductMetaInfoDAO {
	if pdInfo.AlloffCategory == nil {
		alloffCat := GetAlloffCategory(pdInfo)
		pdInfo.AlloffCategory = alloffCat
		pdInfo.IsCategoryClassified = true
		return pdInfo
	} else if pdInfo.AlloffCategory == nil || !pdInfo.AlloffCategory.Done {
		alloffCat := GetOmniousClassifier(pdInfo)
		pdInfo.AlloffCategory = alloffCat
		pdInfo.IsCategoryClassified = true
		return pdInfo
	}
	return nil
}

func GetAlloffCategory(product *domain.ProductMetaInfoDAO) *domain.ProductAlloffCategoryDAO {
	category1, category2, done := categoryClassifier(product)
	return &domain.ProductAlloffCategoryDAO{
		First:   category1,
		Second:  category2,
		Done:    done,
		Touched: false,
	}
}

func categoryClassifier(pdInfo *domain.ProductMetaInfoDAO) (category1, category2 *domain.AlloffCategoryDAO, done bool) {
	//  1. brand + category 보고 분류1 실행
	//  만약 티셔츠, 원피스, 스커트인경우에는 카테고리2 분류 필요 없음
	//  2. product name보고 분류2실행

	// 향후 두개 합쳐서 Keyname으로 query하면 좋겠다, 하지만 지금은 keyname이 틀리는 게 있을 확률이 높으니
	possibleCat2 := []string{}
	classifier, err := ioc.Repo.ClassifyRules.GetByKeyname(pdInfo.Brand.KeyName, pdInfo.Category.Name)
	if err != nil {
		log.Println("Brand key Category Key find no rules:", pdInfo.Brand.KeyName, pdInfo.Category.Name)
		return
	}

	for k, v := range classifier.HeuristicRules {
		loweredName := strings.ToLower(pdInfo.OriginalName)
		if strings.Contains(loweredName, k) {
			possibleCat2 = append(possibleCat2, v)
		}
	}

	if classifier.AlloffCategory1 == nil {
		if len(possibleCat2) == 0 {
			return nil, nil, false
		} else if len(possibleCat2) == 1 {
			cat2, err := ioc.Repo.AlloffCategories.GetByName(possibleCat2[0])
			if err != nil {
				log.Println("error: no matched possible cat2", err)
			}
			if cat2.Level == 2 {
				cat1, _ := ioc.Repo.AlloffCategories.Get(cat2.ParentId.Hex())
				return cat1, cat2, true
			} else {
				return cat2, nil, true
			}

		} else {
			return nil, nil, false
		}
	}

	if classifier.AlloffCategory1.Name == "가방" || classifier.AlloffCategory1.Name == "원피스/세트" || classifier.AlloffCategory1.Name == "신발" || classifier.AlloffCategory1.Name == "스커트" || classifier.AlloffCategory1.Name == "라운지/언더웨어" || classifier.AlloffCategory1.Name == "패션잡화" {
		return classifier.AlloffCategory1, nil, true
	}

	if classifier.AlloffCategory2 != nil {
		return classifier.AlloffCategory1, classifier.AlloffCategory2, true
	}

	if len(possibleCat2) == 0 {
		return classifier.AlloffCategory1, nil, false
	} else if len(possibleCat2) == 1 {
		cat2, _ := ioc.Repo.AlloffCategories.GetByName(possibleCat2[0])
		return classifier.AlloffCategory1, cat2, true
	} else {
		log.Println("Possible Cats", possibleCat2)
		// for _, possibleCat := range possibleCat2 {
		// 	if possibleCat == "티셔츠" {
		// 		cat2, _ := ioc.Repo.AlloffCategories.GetByName(possibleCat)
		// 		return classifier.AlloffCategory1, cat2, true
		// 	}
		// }
	}

	return classifier.AlloffCategory1, nil, false
}

func GetOmniousClassifier(pdInfo *domain.ProductMetaInfoDAO) *domain.ProductAlloffCategoryDAO {
	category1, category2, taggingResult, done, err := omniousClassifier(pdInfo)
	if err != nil {
		return &domain.ProductAlloffCategoryDAO{}
	}

	return &domain.ProductAlloffCategoryDAO{
		TaggingResults: taggingResult,
		First:          category1,
		Second:         category2,
		Done:           done,
		Touched:        false,
	}
}

func omniousClassifier(pdInfo *domain.ProductMetaInfoDAO) (category1, category2 *domain.AlloffCategoryDAO, taggingResult *domain.TaggingResultDAO, done bool, err error) {
	category1 = &domain.AlloffCategoryDAO{}
	category2 = &domain.AlloffCategoryDAO{}
	taggingResult = &domain.TaggingResultDAO{}
	done = false
	err = nil

	omniousTargetImg := ""

	// 이미지를 썸네일 > 캐시된 이미지 > 이미지 순으로 가져온다
	if pdInfo.ThumbnailImage != "" {
		omniousTargetImg = pdInfo.ThumbnailImage
	} else if len(pdInfo.CachedImages) > 0 {
		omniousTargetImg = pdInfo.CachedImages[0]
	} else if len(pdInfo.Images) > 0 {
		omniousTargetImg = pdInfo.Images[0]
	} else {
		return nil, nil, nil, false, fmt.Errorf("ERR502:no images on product")
	}

	data, err := omnious.GetOmniousData(omniousTargetImg)
	if err != nil {
		config.Logger.Error("error occurred on getting omnious data", zap.Error(err))
	}

	category1 = omnious.MapOmniousCategoryToCategoryClassifier(data)

	return
}
