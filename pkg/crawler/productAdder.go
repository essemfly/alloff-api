package crawler

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/classifier"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddProduct(request ProductsAddRequest) {
	pdInfo, err := ioc.Repo.ProductMetaInfos.GetByProductID(request.Brand.KeyName, request.ProductID)

	if err == mongo.ErrNoDocuments {
		// 새로운 상품이 필요한 경우
		pdInfo := &domain.ProductMetaInfoDAO{
			Created: time.Now(),
			Updated: time.Now(),
		}
		pdInfo.SetBrandAndCategory(request.Brand, request.Source)
		pdInfo.SetGeneralInfo(request.ProductName, request.ProductID, request.ProductUrl, request.Images, request.Sizes, request.Colors, request.Description)
		pdInfo.SetPrices(int(request.OriginalPrice), int(request.SalesPrice), domain.CurrencyKRW)

		_, err = ioc.Repo.ProductMetaInfos.Insert(pdInfo)
		if err != nil {
			log.Println("productinfo 1", err)
		}
	} else if err != nil {
		// 에러가 발생한 경우
		log.Println("productinfo 2", err)
	} else {
		// 기존에 상품이 있던 경우
		pdInfo.SetPrices(int(request.OriginalPrice), int(request.SalesPrice), domain.CurrencyKRW)
		_, err := ioc.Repo.ProductMetaInfos.Upsert(pdInfo)
		if err != nil {
			log.Println("productinfo 3", err)
		}
	}

	pdInfo, err = ioc.Repo.ProductMetaInfos.GetByProductID(request.Brand.KeyName, request.ProductID)
	if err != nil {
		log.Println("err", err)
	}

	pd, err := ioc.Repo.Products.GetByMetaID(pdInfo.ID.Hex())
	if err == mongo.ErrNoDocuments {
		pd = &domain.ProductDAO{
			AlloffName:  pdInfo.OriginalName,
			ProductInfo: pdInfo,
			Removed:     false,
			Created:     time.Now(),
			Updated:     time.Now(),
		}
	} else if err != nil {
		log.Println("IsThere?", err)
	} else {
		pd.Updated = time.Now()
	}

	// TODO: Category classifier, Dynamic prices, Dynamic instruction, dynamic scores should be uploaded
	alloffCat := classifier.GetAlloffCategory(pd)
	alloffScore := GetProductScore(pd)
	alloffPrice := GetProductPrice(pd)
	alloffInstruction := GetProductDescription(pd)

	pd.UpdateAlloffCategory(alloffCat)
	pd.UpdateInventory(request.Inventories)
	pd.UpdateScore(alloffScore)
	pd.UpdateInstruction(alloffInstruction)
	if pd.DiscountedPrice > int(alloffPrice) {
		oldPrice := pd.DiscountedPrice
		pd.UpdatePrice(float32(pd.DiscountedPrice), alloffPrice)
		err := ioc.Repo.ProductDiffs.Insert(&domain.ProductDiffDAO{
			OldPrice:   oldPrice,
			NewProduct: pd,
			Type:       "price",
			IsPushed:   false,
		})
		if err != nil {
			log.Println("error occured in product dif update")
		}
	}

	if pd.ID == primitive.NilObjectID {
		// pd.UpdateAlloffCategory(alloffCat)
		_, err = ioc.Repo.Products.Insert(pd)
		if err != nil {
			log.Println("product 1", err)
		}

	} else {
		_, err = ioc.Repo.Products.Upsert(pd)
		if err != nil {
			log.Println("product 2", err)
		}
	}
}
