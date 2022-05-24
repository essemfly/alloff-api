package omnious

import (
	"fmt"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func mapPostResponseToResult(pr *PostResponse) (*OmniousResult, error) {
	if len(pr.Data.Objects) > 1 {
		return nil, fmt.Errorf("ERR500:image has two more items")
	}
	if len(pr.Data.Objects[0].Tags) == 0 {
		return nil, fmt.Errorf("ERR501:omnious tagger did not retrieves data")
	}
	omniousData := pr.Data.Objects[0].Tags[0]

	taggingResult := domain.TaggingResultDAO{
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

	res := &OmniousResult{
		Category:         omniousData.Category,
		TaggingResultDAO: taggingResult,
	}

	return res, nil
}

func MapOmniousCategoryToCategoryClassifier(omniousData *OmniousResult) *domain.AlloffCategoryDAO {
	catMap := map[string]string{
		"탑":       "1_top",
		"블라우스":    "1_top",
		"캐주얼 상의":  "1_top",
		"니트웨어":    "1_top",
		"셔츠":      "1_top",
		"베스트":     "1_outer",
		"코트":      "1_outer",
		"재킷":      "1_outer",
		"점퍼":      "1_outer",
		"패딩":      "1_outer",
		"청바지":     "1_bottom",
		"팬츠":      "1_bottom",
		"스커트":     "1_skirt",
		"드레스":     "1_onePiece",
		"점프수트":    "1_onePiece",
		"잡화":      "1_accessory",
		"모자":      "1_accessory",
		"가방":      "1_bags",
		"지갑":      "1_bags",
		"부츠/워커":   "1_shoes",
		"정장구두":    "1_shoes",
		"운동화":     "1_shoes",
		"로퍼":      "1_shoes",
		"슬리퍼":     "1_shoes",
		"뮬":       "1_shoes",
		"샌들":      "1_shoes",
		"펌프스":     "1_shoes",
		"스포츠화":    "1_shoes",
		"목걸이":     "1_jewelry",
		"귀걸이/피어싱": "1_jewelry",
		"팔찌/발찌":   "1_jewelry",
		"반지":      "1_jewelry",
		"브로치":     "1_jewelry",
		"팬던트":     "1_jewelry",
		"수영복":     "1_jewelry",
	}

	// Item 에서 파자마자상의 혹은 잠옷바지가 나오는 경우에는 특별히 라운지/언더웨어로 정한다.
	omniousItem := omniousData.Item.Name
	if omniousItem == "파자마상의" || omniousItem == "잠옷바지" {
		return &domain.AlloffCategoryDAO{
			KeyName: "1_underwear",
			Name:    "라운지/언더웨어",
		}
	}

	// 나머지는 정해진 정책에 따라 정한다.
	omniousCat := omniousData.Category.Name
	alloffCat, err := ioc.Repo.AlloffCategories.GetByKeyname(catMap[omniousCat])
	if err != nil {
		return nil
	}
	return &domain.AlloffCategoryDAO{
		KeyName: alloffCat.KeyName,
		Name:    alloffCat.Name,
	}
}
