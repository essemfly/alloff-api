package malls

import (
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	"github.com/lessbutter/alloff-api/pkg/product"
)

func CrawlSiVillage(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.sivillage.com"),
	)

	num := 0
	totalProducts := 0
	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	c.OnHTML(".ee_goods", func(e *colly.HTMLElement) {
		num += 1
		title := e.ChildText(".ee_goods-data a .ee_tit")
		originalPriceInStr := e.ChildText(".ee_goods-data a .ee_del .ee_number")
		discountedPriceInStr := e.ChildText(".ee_goods-data a .ee_price .ee_number")
		if originalPriceInStr == "" {
			originalPriceInStr = discountedPriceInStr
		}
		originalPrice := utils.ParsePriceString(originalPriceInStr)
		discountedPrice := utils.ParsePriceString(discountedPriceInStr)
		productID := e.Attr("data-goods_no")
		productUrl := "https://www.sivillage.com/goods/initDetailGoods.siv?goods_no=" + productID
		images, colors, sizes, inventories, description := CrawlSiVillageDetail(productID)

		addRequest := &product.ProductCrawlingAddRequest{
			Brand:               brand,
			Source:              source,
			ProductID:           productID,
			ProductName:         title,
			ProductUrl:          productUrl,
			Images:              images,
			Sizes:               sizes,
			Inventories:         inventories,
			Colors:              colors,
			Description:         description,
			OriginalPrice:       float32(originalPrice),
			SalesPrice:          float32(discountedPrice),
			CurrencyType:        domain.CurrencyKRW,
			IsTranslateRequired: false,
		}

		totalProducts += 1
		product.AddProductInCrawling(addRequest)
	})

	c.OnHTML(".ee_paging", func(e *colly.HTMLElement) {
		lastPage := e.ChildAttr("a:last-child", "data-value")
		lastPageInt, _ := strconv.Atoi(lastPage)

		if !strings.Contains(e.Request.URL.Path, "page_idx") {
			currentIdx := 1
			for currentIdx < lastPageInt {
				currentIdx += 1
				pageUrl := strings.Replace(source.CrawlUrl, "initDispCgv", "ctgGoodsTabAjax", 1)
				url := pageUrl + "&page_size=80&page_idx=" + strconv.Itoa(currentIdx)
				c.Visit(url)
			}
		}
	})

	err = c.Visit(source.CrawlUrl + "&page_size=80")
	if err != nil {
		log.Println("err on crawling", err)
	}

	crawler.PrintCrawlResults(source, totalProducts)

	<-worker
	done <- true
}

func CrawlSiVillageDetail(productId string) (images, colors, sizes []string, inventories []domain.InventoryDAO, description map[string]string) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.sivillage.com"),
	)

	productUrl := "https://www.sivillage.com/goods/initDetailGoods.siv?goods_no=" + productId
	description = map[string]string{}
	colors = nil

	c.OnHTML(".ee_product-img #gd_img", func(e *colly.HTMLElement) {
		e.ForEach(".ee_goods", func(_ int, el *colly.HTMLElement) {
			imgUrlRaw := el.ChildAttr("button img", "src")
			imgUrlList := strings.Split(imgUrlRaw, "?")
			images = append(images, imgUrlList[0])
		})
	})

	c.OnHTML(".gds_chk", func(e *colly.HTMLElement) {
		sizeFound := false
		e.ForEach("dl", func(_ int, el *colly.HTMLElement) {
			if el.ChildText("dt") == "옵션" {
				el.ForEach("dd ul li", func(_ int, ele *colly.HTMLElement) {
					sizeInClass := ele.ChildAttr("button", "class")
					size := ele.ChildText("button")
					size = strings.Trim(size, " ")
					if len(size) > 0 {
						sizeFound = true
						sizes = append(sizes, size)
						if !strings.Contains(sizeInClass, "disabled") {
							inventories = append(inventories, domain.InventoryDAO{
								Size:     size,
								Quantity: 10,
							})
						}
					}
				})
				if !sizeFound {
					el.ForEach("dd span ul li", func(_ int, ele *colly.HTMLElement) {
						size := ele.ChildText(".ee_txt")
						size = strings.Trim(size, " ")
						sizes = append(sizes, size)
						if ele.Attr("class") != "ee_disabled" {
							inventories = append(inventories, domain.InventoryDAO{
								Size:     size,
								Quantity: 10,
							})
						}
					})
				}
			}
		})
	})

	infoUrl := "https://www.sivillage.com/goods/getGoodsClssMidAjax.siv"

	formData := url.Values{}
	formData.Set("req_goods_no", productId)
	formData.Set("sel_goods_no", productId)

	errorMessage := "Error occcured in sivilage detail"
	resp, err := utils.RequestRetryer(infoUrl, utils.REQUEST_POST, utils.GetHeader(), formData.Encode(), errorMessage)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	infoDoc, _ := goquery.NewDocumentFromReader(resp.Body)

	infoDoc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		if !strings.Contains(s.Find("th").Text(), "A/S") {
			description[s.Find("th").Text()] = s.Find("td").Text()
		}
	})

	c.Visit(productUrl)
	return
}
