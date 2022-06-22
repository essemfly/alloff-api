package scripts

import (
	"context"
	"encoding/json"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"io/ioutil"
	"log"
	"time"
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

func MigrateFromOldDb(filter bson.M) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, pdCol := oldDbConnector(ctx, "products")
	defer client.Disconnect(ctx)

	cursor, err := pdCol.Find(ctx, filter)
	if err != nil {
		log.Println(err)
	}
	var pdRes []bson.M

	if err = cursor.All(ctx, &pdRes); err != nil {
		log.Println(err)
	}

	brand, err := ioc.Repo.Brands.GetByKeyname("THOMBROWNE")
	if err != nil {
		log.Panic()
	}

	requests := []*productinfo.AddMetaInfoRequest{}
	for _, pd := range pdRes {
		pdInfo, ok := pd["productinfo"].(primitive.M)
		if !ok {
			log.Panic("can not decode productinfo to map : ", pd["_id"])
		}

		alloffcat, ok := pd["alloffcategories"].(primitive.M)
		if !ok {
			log.Panic("can not decode alloffcategories to map : ", pd["_id"])
		}

		alloffCat := &domain.AlloffCategoryDAO{}
		firstCat, ok := alloffcat["first"].(primitive.M)
		if !ok {
			log.Println("can not decode firstCat to map : ", pd["_id"])
		} else {
			alloffCat, err = mapAlloffCatByName(firstCat["name"].(string))
			if err == mongo.ErrNoDocuments {
				alloffCat = &domain.AlloffCategoryDAO{}
				log.Panic("can not find alloff-category by name : ", firstCat["name"])
			}
		}

		sizes := []string{}
		pdSizes, ok := pdInfo["sizes"].(primitive.A)
		invDaos := []*domain.InventoryDAO{}
		for _, size := range pdSizes {
			sizes = append(sizes, size.(string))
			invDaos = append(invDaos, &domain.InventoryDAO{
				Quantity: 1,
				Size:     size.(string),
			})
		}

		salesInstruction, ok := pd["salesinstruction"].(primitive.M)
		if !ok {
			log.Panic("can not decode salesinstruction to map : ", pd["_id"])
		}

		description, ok := salesInstruction["description"].(primitive.M)
		if !ok {
			log.Panic("can not decode description to map : ", pd["_id"])
		}

		images := []string{}
		pdImages, ok := pd["images"].(primitive.A)
		if ok {
			for _, image := range pdImages {
				images = append(images, image.(string))
			}
		}

		thumnailImage := ""
		pdThumbnail, ok := pd["thumbnailimage"].(string)
		if ok {
			thumnailImage = pdThumbnail
		}

		colors := []string{}
		pdColors, ok := pd["colors"].(primitive.A)
		if ok {
			for _, color := range pdColors {
				colors = append(colors, color.(string))
			}
		}

		texts := []string{}
		pdTexts, ok := description["texts"].(primitive.A)
		if ok {
			for _, text := range pdTexts {
				texts = append(texts, text.(string))
			}
		}

		descImages := []string{}
		pdDescImages, ok := description["images"].(primitive.A)
		if ok {
			for _, descImg := range pdDescImages {
				descImages = append(descImages, descImg.(string))
			}
		}

		infos := map[string]string{}
		pdInfos, ok := description["infos"].(primitive.M)
		if ok {
			for k, v := range pdInfos {
				infos[k] = v.(string)
			}
		}

		log.Println("========================================================================================")
		log.Println("alloffname 			: ", pd["alloffname"].(string))
		log.Println("productid 			: ", "colognese-boutique"+pdInfo["originalname"].(string))
		log.Println("producturl 			: ", pdInfo["producturl"].(string))
		log.Println("producttype 		: ", []domain.AlloffProductType{domain.Female})
		log.Println("originalprice 		: ", float32(pd["originalprice"].(int32)))
		log.Println("discountedprice		: ", float32(pd["discountedprice"].(int32)))
		log.Println("currencytype		: ", domain.CurrencyKRW)
		log.Println("brand			: ", brand.KeyName)
		log.Println("source			: ", "manual")
		log.Println("alloffCat			: ", alloffCat.KeyName)
		log.Println("images			: ", images)
		log.Println("thumbnail			: ", thumnailImage)
		log.Println("colors				: ", colors)
		log.Println("sizes				: ", sizes)
		log.Println("inventory			: ", invDaos)
		log.Println("description			: ", texts)
		log.Println("descriptionimages	: ", descImages)
		log.Println("descriptioninfos 	: ", infos)
		log.Println("descriptionrawinfos	: ", infos)
		log.Println("information 	: ", infos)
		log.Println("rawinformation	: ", infos)
		log.Println("isforeigndelivery	: ", true)
		log.Println("earlistdeliveryday	: ", 14)
		log.Println("latestdeliveryday	: ", 21)
		log.Println("isrefundpossible	: ", true)
		log.Println("refundfee	 	: ", 100000)
		log.Println("modulename	 	: ", "manual")
		log.Println("========================================================================================")
		request := &productinfo.AddMetaInfoRequest{
			AlloffName:           pd["alloffname"].(string),
			ProductID:            "colognese-boutique" + pdInfo["originalname"].(string),
			ProductUrl:           pdInfo["producturl"].(string),
			ProductType:          []domain.AlloffProductType{domain.Female},
			OriginalPrice:        float32(pd["originalprice"].(int32)),
			DiscountedPrice:      float32(pd["discountedprice"].(int32)),
			CurrencyType:         domain.CurrencyKRW,
			Brand:                brand,
			Source:               &domain.CrawlSourceDAO{CrawlModuleName: "manual"},
			AlloffCategory:       alloffCat,
			Images:               images,
			ThumbnailImage:       thumnailImage,
			Colors:               colors,
			Sizes:                sizes,
			Inventory:            invDaos,
			Description:          texts,
			DescriptionImages:    descImages,
			DescriptionInfos:     infos,
			DescriptionRawInfos:  infos,
			Information:          infos,
			RawInformation:       infos,
			IsForeignDelivery:    true,
			EarliestDeliveryDays: 14,
			LatestDeliveryDays:   21,
			IsRefundPossible:     true,
			RefundFee:            100000,
			ModuleName:           "manual",
			IsTranslateRequired:  false,
			IsInventoryMapped:    false,
			IsRemoved:            false,
			IsSoldout:            false,
		}
		requests = append(requests, request)
	}

	totalProducts := 0
	for _, request := range requests {
		_, err = productinfo.AddProductInfo(request)
		if err != nil {
			log.Println(request.AlloffName, "<== 못넣음")
			continue
		} else {
			totalProducts += 1
			log.Println(request.AlloffName, ": added")
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

func mapAlloffCatByName(catName string) (*domain.AlloffCategoryDAO, error) {
	lv1CatName := ""
	switch catName {
	case "코트", "점퍼", "자켓", "베스트", "패딩", "야상":
		lv1CatName = "아우터"
	case "슬랙스", "데님", "팬츠", "팬츠/데님":
		lv1CatName = "바지"
	case "티셔츠", "니트/스웨터", "가디건", "블라우스", "셔츠", "맨투맨", "후드", "민소매", "니트웨어", "셔츠/블라우스":
		lv1CatName = "상의"
	case "원피스":
		lv1CatName = "원피스/세트"
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

func oldDbConnector(ctx context.Context, colName string) (*mongo.Client, *mongo.Collection) {
	options := options.Client().ApplyURI("mongodb://" + viper.GetString("MONGO_URL") + "/" + viper.GetString("MONGO_DB_NAME") + "?&connect=direct&replicaSet=rs0&readPreference=secondaryPreferred&retryWrites=false").SetAuth(options.Credential{
		Username: viper.GetString("MONGO_USERNAME"),
		Password: viper.GetString("MONGO_PASSWORD"),
	})
	client, err := mongo.NewClient(options)
	if err != nil {
		log.Println(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
	}

	log.Println(client)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(viper.GetString("MONGO_DB_NAME"))
	return client, db.Collection(colName)
}
