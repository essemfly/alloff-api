package alloffcategory

import (
	"fmt"
	"log"
	"strings"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/omnious"
	"go.uber.org/zap"

	"github.com/lessbutter/alloff-api/config/ioc"
)

func InferAlloffCategory(pdInfo *domain.ProductMetaInfoDAO) (*domain.ProductAlloffCategoryDAO, error) {
	alloffCat, err := categoryClassifier(pdInfo)
	if err != nil {
		return nil, err
	}
	if alloffCat != nil {
		return alloffCat, nil
	}

	// if pdInfo.AlloffCategory == nil || !pdInfo.AlloffCategory.Done {
	// 	return omniousClassifier(pdInfo)
	// }

	return pdInfo.AlloffCategory, nil
}

func categoryClassifier(pdInfo *domain.ProductMetaInfoDAO) (*domain.ProductAlloffCategoryDAO, error) {
	//  1. brand + category 보고 분류1 실행
	//  만약 티셔츠, 원피스, 스커트인경우에는 카테고리2 분류 필요 없음
	//  2. product name보고 분류2실행

	// 향후 두개 합쳐서 Keyname으로 query하면 좋겠다, 하지만 지금은 keyname이 틀리는 게 있을 확률이 높으니
	res := &domain.ProductAlloffCategoryDAO{}
	possibleCat2 := []string{}
	classifier, err := ioc.Repo.ClassifyRules.GetByKeyname(pdInfo.Brand.KeyName, pdInfo.Category.Name)
	if err != nil {
		log.Println("Brand key Category Key find no rules:", pdInfo.Brand.KeyName, pdInfo.Category.Name, " find using ommnious")
		alloffCat, err := omniousClassifier(pdInfo)
		if err != nil {
			config.Logger.Error("error occurred on get omnious data : ", zap.Error(err))
			return nil, err
		}
		return alloffCat, nil
	}

	for k, v := range classifier.HeuristicRules {
		loweredName := strings.ToLower(pdInfo.OriginalName)
		if strings.Contains(loweredName, k) {
			possibleCat2 = append(possibleCat2, v)
		}
	}

	if classifier.AlloffCategory1 == nil {
		if len(possibleCat2) == 0 {
			return nil, fmt.Errorf("Classifying ")
		} else if len(possibleCat2) == 1 {
			cat2, err := ioc.Repo.AlloffCategories.GetByName(possibleCat2[0])
			if err != nil {
				log.Println("error: no matched possible cat2", err)
			}
			if cat2.Level == 2 {
				cat1, _ := ioc.Repo.AlloffCategories.Get(cat2.ParentId.Hex())
				res.First = cat1
				res.Done = true
				return res, nil
			} else {
				res.First = cat2
				res.Done = true
				return res, nil
			}
		} else {
			return nil, fmt.Errorf("Classifying ")
		}
	}

	if classifier.AlloffCategory1.Name == "가방" || classifier.AlloffCategory1.Name == "원피스/세트" || classifier.AlloffCategory1.Name == "신발" || classifier.AlloffCategory1.Name == "스커트" || classifier.AlloffCategory1.Name == "라운지/언더웨어" || classifier.AlloffCategory1.Name == "패션잡화" {
		res.First = classifier.AlloffCategory1
		res.Done = true
		return res, nil
	}

	// *** Temp ignore and deactivate ProductAlloffCategory.Second ***

	//if classifier.AlloffCategory2 != nil {
	//	return classifier.AlloffCategory1, classifier.AlloffCategory2, true
	//}
	//
	//if len(possibleCat2) == 0 {
	//	return classifier.AlloffCategory1, nil, false
	//} else if len(possibleCat2) == 1 {
	//	cat2, _ := ioc.Repo.AlloffCategories.GetByName(possibleCat2[0])
	//	return classifier.AlloffCategory1, cat2, true
	//} else {
	//	log.Println("Possible Cats", possibleCat2)
	//	// for _, possibleCat := range possibleCat2 {
	//	// 	if possibleCat == "티셔츠" {
	//	// 		cat2, _ := ioc.Repo.AlloffCategories.GetByName(possibleCat)
	//	// 		return classifier.AlloffCategory1, cat2, true
	//	// 	}
	//	// }
	//}

	res.First = classifier.AlloffCategory1
	res.Done = true

	return res, nil
}

func omniousClassifier(pdInfo *domain.ProductMetaInfoDAO) (*domain.ProductAlloffCategoryDAO, error) {
	res := &domain.ProductAlloffCategoryDAO{}
	omniousTargetImg := ""
	// 이미지를 썸네일 > 캐시된 이미지 > 이미지 순으로 가져온다
	if pdInfo.ThumbnailImage != "" {
		omniousTargetImg = pdInfo.ThumbnailImage
	} else if len(pdInfo.CachedImages) > 0 {
		omniousTargetImg = pdInfo.CachedImages[0]
	} else if len(pdInfo.Images) > 0 {
		omniousTargetImg = pdInfo.Images[0]
	} else {
		return nil, fmt.Errorf("ERR502:no images on product")
	}

	data, err := omnious.GetOmniousData(omniousTargetImg)
	if err != nil {
		config.Logger.Error("error occurred on getting omnious data", zap.Error(err))
		return nil, err
	}

	category1 := omnious.MapOmniousCategoryToCategoryClassifier(data)
	res.First = category1
	res.Done = true
	return res, nil
}
