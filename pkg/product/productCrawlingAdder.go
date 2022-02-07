package product

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/classifier"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddProductInCrawling(request *ProductCrawlingAddRequest) {
	pdInfo, err := ProcessProductInfoRequest(request)
	if err != nil {
		log.Println(err)
		return
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
		pd.ProductInfo = pdInfo
		pd.Updated = time.Now()
		pd.Removed = false
	}

	// PD위의 정보 저것들을 해놓고, 밑에 또 process product request하는 이유가 뭐임?
	ProcessCrawlingProductRequest(pd, request)

}

func ProcessProductInfoRequest(request *ProductCrawlingAddRequest) (*domain.ProductMetaInfoDAO, error) {
	pdInfo, err := ioc.Repo.ProductMetaInfos.GetByProductID(request.Brand.KeyName, request.ProductID)

	var newPdInfo = &domain.ProductMetaInfoDAO{}
	if err == mongo.ErrNoDocuments {
		// 새로운 상품이 필요한 경우
		newPdInfo, err = AddCrawlingProductInfo(request)
		if err != nil {
			log.Println("err on insert product info ", err)
			return nil, err
		}
	} else if err == nil {
		// 기존에 상품이 있던 경우
		newPdInfo, err = UpdateProductInfo(pdInfo, request)
		if err != nil {
			log.Println("err on update product info", err)
			return nil, err
		}
	} else {
		// 에러가 발생한 경우
		log.Println("productinfo 2", err)
		return nil, err
	}

	return newPdInfo, nil
}

func AddCrawlingProductInfo(request *ProductCrawlingAddRequest) (*domain.ProductMetaInfoDAO, error) {
	pdInfo := &domain.ProductMetaInfoDAO{
		Created: time.Now(),
		Updated: time.Now(),
	}
	pdInfo.SetBrandAndCategory(request.Brand, request.Source)
	pdInfo.SetGeneralInfo(request.ProductName, request.ProductID, request.ProductUrl, request.Images, request.Sizes, request.Colors, request.Description)
	pdInfo.SetPrices(int(request.OriginalPrice), int(request.SalesPrice), domain.CurrencyKRW)

	newPdInfo, err := ioc.Repo.ProductMetaInfos.Insert(pdInfo)
	if err != nil {
		log.Println("productinfo 1", err)
		return nil, err
	}
	return newPdInfo, nil
}

func UpdateProductInfo(pdInfo *domain.ProductMetaInfoDAO, request *ProductCrawlingAddRequest) (*domain.ProductMetaInfoDAO, error) {
	pdInfo.SetBrandAndCategory(request.Brand, request.Source)
	pdInfo.SetGeneralInfo(request.ProductName, request.ProductID, request.ProductUrl, request.Images, request.Sizes, request.Colors, request.Description)
	pdInfo.SetPrices(int(request.OriginalPrice), int(request.SalesPrice), request.CurrencyType)
	updatedPdInfo, err := ioc.Repo.ProductMetaInfos.Upsert(pdInfo)
	if err != nil {
		log.Println("productinfo 3", err)
		return nil, err
	}
	return updatedPdInfo, nil
}

func ProcessCrawlingProductRequest(pd *domain.ProductDAO, request *ProductCrawlingAddRequest) {
	if pd.AlloffCategories == nil || !pd.AlloffCategories.Done {
		alloffCat := classifier.GetAlloffCategory(pd)
		pd.UpdateAlloffCategory(alloffCat)
	}

	alloffInstruction := GetProductDescription(pd, request.Source)
	pd.ProductInfo.Source = request.Source // Source가 업데이트될 시 Source 업데이트용이다.
	pd.UpdateInstruction(alloffInstruction)

	alloffScore := GetProductScore(pd)
	pd.UpdateScore(alloffScore)
	pd.UpdateInventory(request.Inventories)

	alloffPrice := GetProductPrice(pd)
	lastPrice := pd.DiscountedPrice
	isPriceUpdated := pd.UpdatePrice(alloffPrice)

	if isPriceUpdated {
		err := InsertProductDiff(pd, lastPrice)
		if err != nil {
			log.Println("error on insert product diff", err)
		}
	}

	if pd.ID == primitive.NilObjectID {
		_, err := ioc.Repo.Products.Insert(pd)
		if err != nil {
			log.Println("product 1", err)
		}

	} else {
		_, err := ioc.Repo.Products.Upsert(pd)
		if err != nil {
			log.Println("product 2", err)
		}
	}
}
