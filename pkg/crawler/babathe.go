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
	stockUrl := "https://pc.babathe.com/goods/includeGoodsDtlItemStockQtyList"
	productUrl := "https://pc.babathe.com/goods/indexGoodsDetail?goodsId="
	infoUrl := "https://pc.babathe.com/goods/includeGoodsDtlNotifyList?goodsId="

	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	newProducts := []*domain.ProductDAO{}
	headers := map[string]string{
		"accept":          "*/*",
		"content-type":    "application/x-www-form-urlencoded;charset=UTF-8",
		"connection":      "keep-alive",
		"user-agent":      "Crawler",
		"accept-language": "ko-KR",
	}

	jsonHeaders := map[string]string{
		"accept": "application/json",
	}

	for {
		products := []*domain.ProductDAO{}

		jsonStr := buildBabatheJsonBody(pageNum, source)
		errorMessage := "Crawl Failed: Source " + source.Category.KeyName
		resp, err := utils.RequestRetryer(crawlUrl, utils.REQUEST_POST, headers, jsonStr, errorMessage)
		if err != nil {
			break
		}

		defer resp.Body.Close()

		doc, _ := goquery.NewDocumentFromReader(resp.Body)
		doc.Find("li").Each(func(i int, s *goquery.Selection) {
			title := s.Find("a div.prd-info div.name div.tx-ovf").Text()
			origPriceString := s.Find("a div.prd-info div.prd-price div.price-org").Text()
			origPrice := utils.ParsePriceString(origPriceString)
			curPriceString := s.Find("a div.prd-info div.prd-price div.price strong").Text()
			curPrice := utils.ParsePriceString(curPriceString)
			if curPrice == 0 {
				curPrice = origPrice
			} else if origPrice == 0 {
				origPrice = curPrice
			}

			discountRate := utils.CalculateDiscountRate(float32(origPrice), float32(curPrice))
			productId, _ := s.Find("button.btn-wish").Attr("goodsid")
			productUrl := productUrl + productId

			var images []string
			var sizes []string
			description := map[string]string{}

			errorMessage = "Crawl Failed: Product Detail" + source.Category.KeyName + " - " + productUrl
			resp, err := utils.RequestRetryer(productUrl, utils.REQUEST_GET, jsonHeaders, "", errorMessage)
			if err != nil {
				return
			}

			defer resp.Body.Close()
			detailDoc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				log.Println("FAIL?", err, resp.Body)
				return
			}

			detailDoc.Find("div.gallery-top div.swiper-wrapper div.swiper-slide").Each(func(i int, s *goquery.Selection) {
				image, _ := s.Find("div img").Attr("src")
				images = append(images, image)
			})

			formData := url.Values{}
			formData.Set("goodsId", productId)

			errorMessage = "Crawl Failed: Product Stock" + source.Category.KeyName + " - " + stockUrl
			resp, err = utils.RequestRetryer(stockUrl, utils.REQUEST_POST, headers, formData.Encode(), errorMessage)
			if err != nil {
				return
			}

			defer resp.Body.Close()
			stockCrawlResponse := &BabatheStokPostData{}
			json.NewDecoder(resp.Body).Decode(stockCrawlResponse)
			for _, val := range stockCrawlResponse.Wrapper {
				if val.Quantity > 0 {
					sizes = append(sizes, val.Size)
				}
			}

			resp, err = utils.RequestRetryer(infoUrl+productId, utils.REQUEST_GET, headers, "", errorMessage)
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

			soldout := false
			if len(sizes) == 0 {
				soldout = true
			}
			// (TODO) 0인 것들도 다 가져오는 Sizes가 있어야하고, inventory를 0으로 바꿔주는 코드가 있어야함

			product := domain.ProductDAO{
				ProductInfo: &domain.ProductMetaInfoDAO{
					ProductID:    productId,
					Category:     &source.Category,
					Brand:        brand,
					OriginalName: title,
					Price: &domain.PriceDAO{
						OriginalPrice: float32(origPrice),
						CurrencyType:  domain.CurrencyKRW,
						SellersPrice:  float32(curPrice),
					},
					Images:         images,
					ProductUrl:     productUrl,
					Description:    description,
					SizeAvailable:  sizes,
					ColorAvailable: nil,
				},
				// (TODO) 이 부분 수정 더 되어야함
				DiscountedPrice:  curPrice,
				DiscountRate:     discountRate,
				Soldout:          soldout,
				Removed:          false,
				Inventory:        nil,
				Score:            &domain.ProductScoreInfoDAO{},
				SalesInstruction: &domain.AlloffInstructionDAO{},
				Created:          time.Now(),
				Updated:          time.Now(),
			}

			products = append(products, &product)
		})

		newProducts = append(newProducts, products...)

		if len(products) > 0 {
			pageNum += 1
			products = []*domain.ProductDAO{}
		} else {
			break
		}
	}

	// ProcessCralwedProducts(source, newProducts)
	<-worker
	done <- true
}

func buildBabatheJsonBody(pageNum int, source *domain.CrawlSourceDAO) string {
	pureString := `searchDisplay=100&pageNumber=` + strconv.Itoa(pageNum) + `&searchCategory=` + source.MainCategoryKey + `&brandNo=` + source.BrandIdentifier + `&outletYn=Y&gubun=display`

	newString := strings.ReplaceAll(pureString, " ", "")
	return newString
}
