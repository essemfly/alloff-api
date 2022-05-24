package malls

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	"github.com/lessbutter/alloff-api/pkg/product"
)

type OptionParser struct {
	Option []OptionDetail
}

type OptionDetail struct {
	Stock    string
	Opt1     string
	Opt2     string
	Totstock string
}

func CrawlTheamall(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	pageNum := 1
	totalProducts := 0

	crawlUrl := source.CrawlUrl + "&mode=ajaxItems"
	stockUrl := "http://www.theamall.com/product/ajax"

	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	for {
		numProducts := 0

		newlUrl := crawlUrl + "&page=" + strconv.Itoa(pageNum)
		errorMessage := "Crawl Failed: Source " + source.Category.KeyName
		resp, err := utils.RequestRetryer(newlUrl, utils.REQUEST_POST, utils.GetGeneralHeader(), "", errorMessage)
		if err != nil {
			break
		}

		defer resp.Body.Close()

		doc, _ := goquery.NewDocumentFromReader(resp.Body)
		doc.Find("li").Each(func(i int, s *goquery.Selection) {
			numProducts += 1
			title := s.Find("a .caption .goodsnm").Text()
			origPriceString := s.Find("a .caption .price strike i").Text()
			origPrice := utils.ParsePriceString(origPriceString)
			curPriceString := s.Find("a .caption .price > i").Text()
			curPrice := utils.ParsePriceString(curPriceString)
			if curPrice == 0 {
				curPrice = origPrice
			} else if origPrice == 0 {
				origPrice = curPrice
			}

			productUrlNaive, _ := s.Find("a").Attr("href")
			productID := strings.Split(productUrlNaive, "/")[2]
			productUrl := "http://www.theamall.com/product/" + productID
			mainImgTag, _ := s.Find("a div").Attr("style")

			sizes := []string{}
			colors := []string{}
			images := []string{}
			inventories := []*domain.InventoryDAO{}

			if len(strings.Split(mainImgTag, "'")) > 1 {
				val := strings.Split(mainImgTag, "'")[1]
				if !strings.HasPrefix(val, "http") {
					images = append(images, "https:"+val)
				} else {
					images = append(images, val)
				}
			}

			description := map[string]string{}

			errorMessage = "Crawl Failed: Product Detail" + source.Category.KeyName + " - " + productUrl
			resp, err := utils.RequestRetryer(productUrl, utils.REQUEST_GET, utils.GetGeneralHeader(), "", errorMessage)
			if err != nil {
				return
			}

			defer resp.Body.Close()

			detailDoc, _ := goquery.NewDocumentFromReader(resp.Body)

			detailDoc.Find("div.thumbnails a img").Each(func(i int, s *goquery.Selection) {
				if i != 0 {
					val, _ := s.Attr("src")
					if !strings.HasPrefix(val, "http") {
						images = append(images, "http:"+val)
					} else {
						images = append(images, val)
					}

				}
			})

			// 일반 상세이미지
			detailDoc.Find("div.content > div > img").Each(func(i int, s *goquery.Selection) {
				val, _ := s.Attr("src")
				if !strings.HasPrefix(val, "http") {
					images = append(images, "http:"+val)
				} else {
					images = append(images, val)
				}
			})

			// 엄청 긴 상세이미지
			detailDoc.Find("div.content > p > img").Each(func(i int, s *goquery.Selection) {
				val, _ := s.Attr("src")
				if !strings.HasPrefix(val, "http") {
					images = append(images, "http:"+val)
				} else {
					images = append(images, val)
				}
			})

			detailDoc.Find("div.product-info div.option dl dd").Each(func(i int, s *goquery.Selection) {
				if i == 0 {
					s.Find("select option").Each(func(j int, s2 *goquery.Selection) {
						if j != 0 {
							_, exist := s2.Attr("disabled")
							if !exist {
								optionValue, _ := s2.Attr("value")
								parser := &OptionParser{}

								errorMessage = "Crawl Failed: Product Stock" + source.Category.KeyName + " - " + stockUrl
								stockResp, err := utils.RequestRetryer(stockUrl+"?mode=option2&goodsno="+productID+"&optno="+optionValue, utils.REQUEST_GET, utils.GetGeneralHeader(), "", errorMessage)
								if err != nil {
									return
								}

								defer stockResp.Body.Close()
								json.NewDecoder(stockResp.Body).Decode(parser)
								for _, val := range parser.Option {
									colors = append(colors, val.Opt1)
									sizes = append(sizes, val.Opt1+" - "+val.Opt2)
									stock, _ := strconv.Atoi(val.Totstock)
									inventories = append(inventories, &domain.InventoryDAO{
										Size:     val.Opt1 + " - " + val.Opt2,
										Quantity: stock,
									})
								}
							}
						}
					})
				}
			})

			detailDoc.Find("table.table-info").Each(func(i int, s *goquery.Selection) {
				s.Find("tbody tr").Each(func(j int, s2 *goquery.Selection) {
					if !strings.Contains(s2.Find("th").Text(), "AS") {
						description[s2.Find("th").Text()] = s2.Find("td").Text()
					}
				})
			})

			addRequest := &product.AddMetaInfoRequest{
				AlloffName:      title,
				ProductID:       productID,
				ProductUrl:      productUrl,
				ProductType:     []domain.AlloffProductType{domain.Female},
				OriginalPrice:   float32(origPrice),
				DiscountedPrice: float32(curPrice),
				CurrencyType:    domain.CurrencyKRW,
				Brand:           brand,
				Source:          source,
				// AlloffCategory:  nil,
				Images:              images,
				Colors:              colors,
				Sizes:               sizes,
				Inventory:           inventories,
				Information:         description,
				IsForeignDelivery:   false,
				IsTranslateRequired: false,
				ModuleName:          source.CrawlModuleName,
			}

			totalProducts += 1
			product.ProcessAddProductInfoRequests(addRequest)

		})

		if numProducts > 0 {
			pageNum += 1
			numProducts = 0
		} else {
			break
		}
	}

	crawler.PrintCrawlResults(source, totalProducts)
	<-worker
	done <- true
}
