package crawler

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BabatheStokPostData struct {
	Wrapper []BabatheStockInfo `json:"itemAmtStockQtyVOList"`
}

type BabatheInfoGetData struct {
	Wrapper []BabatheInfo `json:"goodsNotifyList"`
}

type BabatheStockInfo struct {
	Quantity int    `json:"optQty"`
	Size     string `json:"attrVal1"`
}

type BabatheInfo struct {
	ItemVal string `json:"itemVal"`
	Name    string `json:"itemNm"`
}

func CrawlBabathe(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	pageNum := 1
	crawlUrl := source.CrawlUrl

	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	for {
		numProducts := 0

		jsonStr := buildBabatheJsonBody(pageNum, source)
		errorMessage := "Crawl Failed: Source " + source.Category.KeyName
		resp, err := utils.RequestRetryer(crawlUrl, utils.REQUEST_POST, utils.GetHeader(), jsonStr, errorMessage)
		if err != nil {
			break
		}
		defer resp.Body.Close()

		doc, _ := goquery.NewDocumentFromReader(resp.Body)
		doc.Find("li").Each(func(i int, s *goquery.Selection) {
			numProducts += 1
			productName, productID, productUrl, origPrice, curPrice := ParseHtml(s)
			images, sizes, colors, inventories, description := CrawlBabatheDetail(productUrl, productID, source)
			_, err := ioc.Repo.ProductMetaInfos.GetByProductID(brand.KeyName, productID)

			if err == mongo.ErrNoDocuments {
				pdInfo := &domain.ProductMetaInfoDAO{
					Created: time.Now(),
					Updated: time.Now(),
				}
				pdInfo.SetBrandAndCategory(brand, source)
				pdInfo.SetGeneralInfo(productName, productID, productUrl, images, sizes, colors, description)
				pdInfo.SetPrices(origPrice, curPrice, domain.CurrencyKRW)

				_, err = ioc.Repo.ProductMetaInfos.Insert(pdInfo)
				if err != nil {
					log.Println(err)
				}
			} else if err != nil {
				log.Println(err)
			}

			pdInfo, err := ioc.Repo.ProductMetaInfos.GetByProductID(brand.KeyName, productID)
			if err != nil {
				log.Println("err", err)
			}

			pd, err := ioc.Repo.Products.GetByMetaID(pdInfo.ID.Hex())
			if err == mongo.ErrNoDocuments {
				pd = &domain.ProductDAO{
					ProductInfo: pdInfo,
					Removed:     false,
					Created:     time.Now(),
					Updated:     time.Now(),
				}
			} else if err != nil {
				log.Println(err)
			}

			pd.UpdateInventory(inventories)

			// TODO: Category classifier, Dynamic prices, Dynamic instruction, dynamic scores should be uploaded
			alloffCat := GetAlloffCategory(pd)
			alloffScore := GetProductScore(pd)
			alloffPrice := GetProductPrice(pd)
			alloffInstruction := GetProductDescription(pd)

			pd.UpdateAlloffCategory(alloffCat)
			pd.UpdateScore(alloffScore)
			pd.UpdatePrice(pdInfo.Price.OriginalPrice, alloffPrice)
			pd.UpdateInstruction(alloffInstruction)

			if pd.ID == primitive.NilObjectID {
				_, err = ioc.Repo.Products.Insert(pd)
			} else {
				_, err = ioc.Repo.Products.Upsert(pd)
			}

			if err != nil {
				log.Println(err)
			}
		})

		if numProducts > 0 {
			log.Println("numProducts", numProducts)
			pageNum += 1
			numProducts = 0
		} else {
			break
		}
	}

	<-worker
	done <- true
}

func ParseHtml(s *goquery.Selection) (productName, productID, productUrl string, origPrice, curPrice int) {
	baseUrl := "https://pc.babathe.com/goods/indexGoodsDetail?goodsId="
	productName = s.Find("a div.prd-info div.name div.tx-ovf").Text()
	origPriceString := s.Find("a div.prd-info div.prd-price div.price-org").Text()
	origPrice = utils.ParsePriceString(origPriceString)
	curPriceString := s.Find("a div.prd-info div.prd-price div.price strong").Text()
	curPrice = utils.ParsePriceString(curPriceString)
	if curPrice == 0 {
		curPrice = origPrice
	} else if origPrice == 0 {
		origPrice = curPrice
	}

	productID, _ = s.Find("button.btn-wish").Attr("goodsid")
	productUrl = baseUrl + productID

	return
}

func CrawlBabatheDetail(productUrl, productId string, source *domain.CrawlSourceDAO) (images, sizes, colors []string, inventories []domain.InventoryDAO, description map[string]string) {
	description = map[string]string{}

	stockUrl := "https://pc.babathe.com/goods/includeGoodsDtlItemStockQtyList"
	infoUrl := "https://pc.babathe.com/goods/includeGoodsDtlNotifyList?goodsId="

	errorMessage := "Crawl Failed: Product Detail" + source.Category.KeyName + " - " + productUrl
	resp, err := utils.RequestRetryer(productUrl, utils.REQUEST_GET, utils.GetJsonHeader(), "", errorMessage)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	detailDoc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("Crawl Failed: Goquery", err, resp.Body)
		return
	}

	detailDoc.Find("div.gallery-top div.swiper-wrapper div.swiper-slide").Each(func(i int, s *goquery.Selection) {
		image, _ := s.Find("div img").Attr("src")
		images = append(images, image)
	})

	formData := url.Values{}
	formData.Set("goodsId", productId)

	errorMessage = "Crawl Failed: Product Stock" + source.Category.KeyName + " - " + stockUrl
	resp, err = utils.RequestRetryer(stockUrl, utils.REQUEST_POST, utils.GetHeader(), formData.Encode(), errorMessage)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	stockCrawlResponse := &BabatheStokPostData{}
	json.NewDecoder(resp.Body).Decode(stockCrawlResponse)
	for _, val := range stockCrawlResponse.Wrapper {
		sizes = append(sizes, val.Size)
		if val.Quantity > 0 {
			inventories = append(inventories, domain.InventoryDAO{
				Size:     val.Size,
				Quantity: val.Quantity,
			})
		}
	}

	resp, err = utils.RequestRetryer(infoUrl+productId, utils.REQUEST_GET, utils.GetHeader(), "", errorMessage)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	infoCrawlResponse := &BabatheInfoGetData{}
	json.NewDecoder(resp.Body).Decode(infoCrawlResponse)
	for _, val := range infoCrawlResponse.Wrapper {
		cleanItem := strings.Replace(val.ItemVal, "<br>", "", -1)
		if !strings.Contains(val.Name, "A/S") {
			description[val.Name] = cleanItem
		}
	}

	colors = nil

	return
}

func buildBabatheJsonBody(pageNum int, source *domain.CrawlSourceDAO) string {
	pureString := `searchDisplay=100&pageNumber=` + strconv.Itoa(pageNum) + `&searchCategory=` + source.MainCategoryKey + `&brandNo=` + source.BrandIdentifier + `&outletYn=Y&gubun=display`

	newString := strings.ReplaceAll(pureString, " ", "")
	return newString
}
