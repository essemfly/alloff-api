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

func startMigration(pgUrl string) {
	resp, err := utils.MakeRequest(pgUrl, utils.REQUEST_GET, map[string]string{}, "")
	if err != nil {
		log.Panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var result Resp
	if err = json.Unmarshal(body, &result); err != nil {
		log.Panic(err)
	}

	brand, _ := ioc.Repo.Brands.GetByKeyname(result.Brand.Keyname)
	if brand.KorName != result.Products[0].Product.BrandKorName {
		log.Panic("브랜드 이상함 재확인 필요")
	}

	for _, oldPd := range result.Products {
		log.Println("now add : ", oldPd.Product.AlloffName)
		alloffCat, err := ioc.Repo.AlloffCategories.GetByName(oldPd.Product.AlloffCategoryName)
		if err == mongo.ErrNoDocuments {
			alloffCat = &domain.AlloffCategoryDAO{}
		}

		pd := oldPd.Product

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

		log.Println(pd.AlloffProductID)
		_, err = productinfo.AddProductInfo(request)
		if err != nil {
			continue
		} else {
			log.Println(oldPd.Product.AlloffName, ": added")
		}
	}
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
