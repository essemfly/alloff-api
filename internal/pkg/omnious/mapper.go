package omnious

import (
	"fmt"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/dto"
)

func mapPostResponseToResult(pr *dto.PostResponse) (*dto.OmniousResult, error) {
	if len(pr.Data.Objects) > 1 {
		return nil, fmt.Errorf("ERR500:image has two more items")
	}
	omniousData := pr.Data.Objects[0].Tags[0]

	taggingResult := dto.TaggingResult{
		Item:         omniousData.Item,
		Colors:       omniousData.Colors,
		ColorDetails: omniousData.ColorDetails,
		Prints:       omniousData.Prints,
		Looks:        omniousData.Looks,
		Textures:     omniousData.Textures,
		Details:      omniousData.Details,
		Length:       omniousData.Length,
		SleeveLength: omniousData.SleeveLength,
		NeckLine:     omniousData.NeckLine,
		Fit:          omniousData.Fit,
		Shape:        omniousData.Shape,
	}

	res := &dto.OmniousResult{
		Category:      omniousData.Category,
		TaggingResult: taggingResult,
	}

	return res, nil
}

func MapOmniousCategoryToCategoryClassifier(omniousCat string) *domain.CategoryClassifier {
	catMap := map[string]string{
		"베스트":     "1_outer",
		"코트":      "1_outer",
		"재킷":      "1_outer",
		"점퍼":      "1_outer",
		"패딩":      "1_outer",
		"청바지":     "1_bottom",
		"팬츠":      "1_bottom",
		"탑":       "1_top",
		"블라우스":    "1_top",
		"캐주얼상의":   "1_top",
		"니트웨어":    "1_top",
		"셔츠":      "1_top",
		"드레스":     "1_onePiece",
		"점프수트":    "1_onePiece",
		"가방":      "1_bags",
		"여행가방":    "1_bags",
		"클러치":     "1_bags",
		"스커트":     "1_skirt",
		"부츠":      "1_shoes",
		"정장":      "1_shoes",
		"운동화":     "1_shoes",
		"로퍼":      "1_shoes",
		"슬리퍼":     "1_shoes",
		"뮬":       "1_shoes",
		"샌들":      "1_shoes",
		"펌프스":     "1_shoes",
		"모자":      "1_accessory",
		"목걸이":     "1_accessory",
		"귀걸이/피어싱": "1_accessory",
		"팔찌/발찌":   "1_accessory",
		"반지":      "1_accessory",
		"브로치":     "1_accessory",
		"팬던트":     "1_accessory",
		"잡화":      "1_accessory",
	}
	alloffCat, err := ioc.Repo.AlloffCategories.GetByKeyname(catMap[omniousCat])
	if err != nil {
		return nil
	}
	return &domain.CategoryClassifier{
		KeyName: alloffCat.KeyName,
		Name:    alloffCat.Name,
	}
}
