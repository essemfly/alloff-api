package classifier

import (
	"fmt"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/omnious"
	"go.uber.org/zap"
)

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
