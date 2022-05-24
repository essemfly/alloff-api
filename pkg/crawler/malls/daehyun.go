package malls

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

type DaehyunResponseParser struct {
	Content    string `json:"content"`
	NextPage   int    `json:"next_page"`
	TotalCount string `json:"total_cnt"`
}

func CrawlDaehyun(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	pageNum := 1
	num := 0

	totalProducts := 0

	for {
		baseQuery := "exec.php?exec_file=skin_module/skin_ajax.php&obj_id=prd_basic&_tmp_file_name=shop%2Fbig_section.php&single_module=prd_basic&striplayout=1"
		url := source.CrawlUrl + baseQuery + "&cno1=" + source.MainCategoryKey + "&module_page=" + strconv.Itoa(pageNum)

		errorMessage := "Crawl Failed: Source " + source.Category.KeyName
		resp, err := utils.RequestRetryer(url, utils.REQUEST_GET, utils.GetGeneralHeader(), "", errorMessage)
		if err != nil {
			log.Println("err", err)
			break
		}

		defer resp.Body.Close()

		brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
		if err != nil {
			log.Println(err)
		}

		crawlResponse := &DaehyunResponseParser{}
		json.NewDecoder(resp.Body).Decode(crawlResponse)

		if crawlResponse.Content == "" {
			break
		}

		r := strings.NewReader(crawlResponse.Content)
		doc, _ := goquery.NewDocumentFromReader(r)

		doc.Find("li").Each(func(i int, s *goquery.Selection) {
			productName := s.Find("div.mojo_box div.info p.name a").Text()
			if productName != "" {
				num += 1
				originalPriceStr := s.Find("div.mojo_box div.info div.price p.consumer").Text()
				discountedPriceStr := s.Find("div.mojo_box div.info div.price p.sell").Text()
				if originalPriceStr == "" {
					originalPriceStr = discountedPriceStr
				}
				originalPrice := utils.ParsePriceString(originalPriceStr)
				discountedPrice := utils.ParsePriceString(discountedPriceStr)
				productUrl, _ := s.Find("div.mojo_box div.img div.prdimg a").Attr("href")
				images, colors, sizes, inventories, description := getDaehyunDetailInfo(productUrl)

				s1 := strings.Split(productUrl, "pno=")
				s2 := strings.Split(s1[1], "&")
				productID := s2[0]

				askey := ""
				laundrykey := ""
				for k := range description {
					if strings.Contains(k, "A/S") {
						askey = k
					}
					if strings.Contains(k, "세탁방법") {
						laundrykey = k
					}
				}

				delete(description, askey)
				delete(description, laundrykey)

				addRequest := &product.AddMetaInfoRequest{
					AlloffName:      productName,
					ProductID:       productID,
					ProductUrl:      productUrl,
					ProductType:     []domain.AlloffProductType{domain.Female},
					OriginalPrice:   float32(originalPrice),
					DiscountedPrice: float32(discountedPrice),
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

			}
		})
		pageNum += 1
	}

	crawler.PrintCrawlResults(source, totalProducts)

	<-worker
	done <- true
}

func getDaehyunDetailInfo(producturl string) (imageUrls, colors, sizes []string, inventories []*domain.InventoryDAO, description map[string]string) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.daehyuninside.com"),
	)

	description = map[string]string{}
	sizeKey := map[string]string{}
	colorOptions := []string{}

	c.OnHTML(".imglist #product_add_image2_list", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			imageurl := el.ChildAttr("div img", "src")
			if imageurl != "" {
				imageUrls = append(imageUrls, imageurl)
			}
		})
	})

	c.OnHTML(".info_view .info .info_list .opt_title", func(e *colly.HTMLElement) {
		if e.ChildText("div:first-child") == "컬러" {
			e.ForEach("select option", func(_ int, el *colly.HTMLElement) {
				color := el.Attr("value")
				colors = append(colors, color)
				optionValue := color
				if optionValue != "" && !strings.Contains(optionValue, "품절") {
					optionsValueList := strings.Split(optionValue, "::")
					if len(optionsValueList) > 3 {
						colorOptions = append(colorOptions, optionsValueList[3])
					}
				}
			})
		} else {
			e.ForEach("ul.text_option li", func(_ int, el *colly.HTMLElement) {
				sizeKey[el.ChildAttr("a", "data")] = el.ChildText("a")
			})
		}
	})

	c.OnHTML(".info_view .toggle_detail .section", func(e *colly.HTMLElement) {
		if e.ChildText(".title") == "상품 정보고시" {
			keys := []string{}
			e.ForEach(".common_cnt .prd_info h3", func(_ int, el *colly.HTMLElement) {
				keys = append(keys, el.Text)
			})
			e.ForEach(".common_cnt .prd_info p", func(i int, el *colly.HTMLElement) {
				allInfo := el.Text
				allInfo = strings.ReplaceAll(allInfo, "\n", "")
				allInfo = strings.ReplaceAll(allInfo, "\t", "")
				description[keys[i]] = allInfo
			})
		}
	})

	c.Visit(producturl)

	// TODO 이쪽 Color 및 상품 재고 어떻게 처리되는지 확인 필요
	soldoutSizes := []string{}
	for _, option := range colorOptions {
		stockUrl := "https://www.daehyuninside.com/main/exec.php?exec_file=shop/getAjaxData.php&exec=getOptionStock&item_no=@" + option

		errorMessage := "Crawl error occured in getting daehyun options: " + stockUrl
		resp, err := utils.RequestRetryer(stockUrl, utils.REQUEST_GET, utils.GetGeneralHeader(), "", errorMessage)
		if err != nil {
			continue
		}

		defer resp.Body.Close()

		b, _ := ioutil.ReadAll(resp.Body)
		soldoutSizes = append(soldoutSizes, string(b))
	}

	for key, value := range sizeKey {
		isSoldout := true

		if len(soldoutSizes) == 0 {
			isSoldout = false
		}

		for _, soldoutMergedString := range soldoutSizes {
			if !strings.Contains(soldoutMergedString, key) {
				isSoldout = false
				break
			}
		}
		if !isSoldout {
			sizes = append(sizes, value)
			inventories = append(inventories, &domain.InventoryDAO{
				Quantity: 10,
				Size:     value,
			})
		}
	}

	return
}
