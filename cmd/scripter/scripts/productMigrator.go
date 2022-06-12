package scripts

import (
	"encoding/json"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"log"
)

func MigrateFromPGUrl(pgUrl string) {
	resp, err := utils.MakeRequest(pgUrl, utils.REQUEST_GET, map[string]string{}, "")
	if err != nil {
		log.Panic(err)
	}

	totalProducts := 0
	errors := []map[string]string{}

	log.Println(":::::::::::::::::::::::::::PG:::::::::::::::::::::::::::")
	log.Println("now on : ", pgUrl)
	log.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::")

	body, err := ioutil.ReadAll(resp.Body)

	var result Resp
	if err = json.Unmarshal(body, &result); err != nil {
		log.Panic(err)
	}

	for _, oldPd := range result.Products {
		log.Println("now add : ", oldPd.Product.AlloffName)
		pd := oldPd.Product

		alloffCat, err := mapAlloffCatByName(pd.AlloffCategoryName)
		if err == mongo.ErrNoDocuments {
			alloffCat = &domain.AlloffCategoryDAO{}
			errors = append(errors, map[string]string{"타입": "카테고리를 찾을 수 없음", "상품": pd.AlloffProductID})
		}

		color := ""
		if val, ok := pd.DescriptionInfos["색상"]; ok {
			color = val
		}

		sizes := []string{}
		for _, inv := range pd.Inventory {
			sizes = append(sizes, inv.Size)
		}

		invDaos := []*domain.InventoryDAO{}
		for _, inv := range pd.Inventory {
			invDaos = append(invDaos, &domain.InventoryDAO{
				Quantity: int(inv.Quantity),
				Size:     inv.Size,
			})
		}

		// 브랜드 korName이 이상하게 입력되어있는건 checkBrandKorName 에서 거른다.
		brandKorName := checkBrandKorName(pd.BrandKorName)

		brand, err := ioc.Repo.Brands.GetByKorname(brandKorName)
		if err != nil {
			errors = append(errors, map[string]string{"type": "can not find brands from korname", "detail": pd.AlloffProductID})
		}

		request := &productinfo.AddMetaInfoRequest{
			AlloffName:           pd.AlloffName,
			ProductID:            pd.ProductURL + pd.AlloffName,
			ProductUrl:           pd.ProductURL,
			ProductType:          []domain.AlloffProductType{domain.Female},
			OriginalPrice:        float32(pd.OriginalPrice),
			DiscountedPrice:      float32(pd.DiscountedPrice),
			CurrencyType:         domain.CurrencyKRW,
			Brand:                brand,
			Source:               &domain.CrawlSourceDAO{CrawlModuleName: "manual"},
			AlloffCategory:       alloffCat,
			Images:               pd.Images,
			ThumbnailImage:       pd.ThumbnailImage,
			Colors:               []string{color},
			Sizes:                sizes,
			Inventory:            invDaos,
			Description:          pd.Description,
			DescriptionImages:    pd.DescriptionImages,
			DescriptionInfos:     pd.DescriptionInfos,
			DescriptionRawInfos:  pd.DescriptionInfos,
			Information:          pd.DescriptionInfos,
			RawInformation:       pd.DescriptionInfos,
			IsForeignDelivery:    pd.IsForeignDelivery,
			EarliestDeliveryDays: int(pd.EarliestDeliveryDays),
			LatestDeliveryDays:   int(pd.LatestDeliveryDays),
			IsRefundPossible:     pd.IsRefundPossible,
			RefundFee:            int(pd.RefundFee),
			ModuleName:           "manual",
			IsTranslateRequired:  false,
			IsInventoryMapped:    false,
			IsRemoved:            false,
			IsSoldout:            false,
		}
		_, err = productinfo.AddProductInfo(request)
		if err != nil {
			errors = append(errors, map[string]string{"타입": "상품을 넣을 수 없음", "상품": oldPd.Product.AlloffProductID})
			continue
		} else {
			totalProducts += 1
			log.Println(oldPd.Product.AlloffName, ": added")
		}
	}

	for _, error := range errors {
		log.Println(error)
	}

	log.Printf("총 상품 >> [%v] 개 입력완료", totalProducts)
}

type Resp struct {
	Brand struct {
		Keyname string `json:"keyname"`
	} `json:"brand"`
	Products []struct {
		Product struct {
			AlloffCategoryName   string            `json:"alloff_category_name"`
			AlloffName           string            `json:"alloff_name"`
			AlloffProductID      string            `json:"alloff_product_id"`
			BrandKorName         string            `json:"brand_kor_name"`
			Description          []string          `json:"description"`
			DescriptionImages    []string          `json:"description_images"`
			DescriptionInfos     map[string]string `json:"description_infos"`
			DiscountedPrice      int64             `json:"discounted_price"`
			EarliestDeliveryDays int64             `json:"earliest_delivery_days"`
			Images               []string          `json:"images"`
			Inventory            []struct {
				Quantity int64  `json:"quantity"`
				Size     string `json:"size"`
			} `json:"inventory"`
			IsForeignDelivery  bool   `json:"is_foreign_delivery"`
			IsRefundPossible   bool   `json:"is_refund_possible"`
			LatestDeliveryDays int64  `json:"latest_delivery_days"`
			OriginalPrice      int64  `json:"original_price"`
			ProductURL         string `json:"product_url"`
			RefundFee          int64  `json:"refund_fee"`
			ThumbnailImage     string `json:"thumbnail_image"`
		} `json:"product"`
	} `json:"products"`
}

func mapAlloffCatByName(catName string) (*domain.AlloffCategoryDAO, error) {
	lv1CatName := ""
	switch catName {
	case "코트", "점퍼", "자켓", "베스트", "패딩", "야상":
		lv1CatName = "아우터"
	case "슬랙스", "데님", "팬츠":
		lv1CatName = "바지"
	case "티셔츠", "니트/스웨터", "가디건", "블라우스", "셔츠", "맨투맨", "후드", "민소매":
		lv1CatName = "상의"
	case "아우터", "상의", "바지", "원피스/세트", "스커트", "라운지/언더웨어", "가방", "신발", "패션잡화":
		lv1CatName = catName
	}

	alloffCat, err := ioc.Repo.AlloffCategories.GetByName(lv1CatName)
	if err != nil {
		return nil, err
	}
	return alloffCat, nil

}

func checkBrandKorName(korName string) string {
	if korName == "막스마라 (인트렌드)" {
		return "막스마라(인트렌드)"
	}
	return korName
}