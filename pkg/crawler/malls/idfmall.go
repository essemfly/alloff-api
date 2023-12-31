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
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
)

type IDFResponseParser struct {
	Content    string `json:"content"`
	NextPage   int    `json:"next_page"`
	TotalCount string `json:"total_cnt"`
}

func CrawlIDFMall(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	pageNum := 1
	num := 0
	totalProducts := 0

	for {
		baseQuery := "exec.php?exec_file=skin_module/skin_ajax.php&obj_id=prd_basic&_tmp_file_name=shop/big_section.php&single_module=prd_basic&striplayout=1&document_url=/shop/search_result.php&sch_sale=N"
		url := source.CrawlUrl + baseQuery + "&cno1=" + source.MainCategoryKey + "&module_page=" + strconv.Itoa(pageNum)

		errorMessage := "Crawl Failed: Source " + source.Category.KeyName
		resp, err := utils.RequestRetryer(url, utils.REQUEST_GET, utils.GetGeneralHeader(), "", errorMessage)
		if err != nil {
			break
		}

		defer resp.Body.Close()

		brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
		if err != nil {
			log.Println(err)
		}

		crawlResponse := &IDFResponseParser{}
		json.NewDecoder(resp.Body).Decode(crawlResponse)

		if crawlResponse.Content == "" {
			break
		}

		r := strings.NewReader(crawlResponse.Content)
		doc, _ := goquery.NewDocumentFromReader(r)

		doc.Find("li").Each(func(i int, s *goquery.Selection) {
			class, _ := s.Find("div.box").Attr("class")
			productUrl, _ := s.Find("div.box div.img div.prdimg a").Attr("href")

			title := s.Find("div.box div.info p.name a").Text()
			if title != "" {
				if !strings.Contains(class, "out") {
					num += 1
					originalPriceStr := s.Find("div.box div.info div.price span.consumer").Text()
					discountedPriceStr := s.Find("div.box div.info div.price span.sell").Text()
					if originalPriceStr == "" {
						originalPriceStr = discountedPriceStr
					}
					originalPrice := utils.ParsePriceString(originalPriceStr)
					discountedPrice := utils.ParsePriceString(discountedPriceStr)

					s1 := strings.Split(productUrl, "pno=")
					s2 := strings.Split(s1[1], "&")
					productID := s2[0]

					images, colors, sizes, inventories, description := getIdfDetailInfo(productUrl)

					addRequest := &productinfo.AddMetaInfoRequest{
						AlloffName:          title,
						ProductID:           productID,
						ProductUrl:          productUrl,
						ProductType:         []domain.AlloffProductType{domain.Female},
						OriginalPrice:       float32(originalPrice),
						DiscountedPrice:     float32(discountedPrice),
						CurrencyType:        domain.CurrencyKRW,
						Brand:               brand,
						Source:              source,
						AlloffCategory:      &domain.AlloffCategoryDAO{},
						Images:              images,
						Colors:              colors,
						Sizes:               sizes,
						Inventory:           inventories,
						Information:         description,
						IsForeignDelivery:   false,
						IsTranslateRequired: false,
						ModuleName:          source.CrawlModuleName,
						IsRemoved:           false,
						IsSoldout:           false,
					}

					totalProducts += 1
					productinfo.ProcessCrawlingInfoRequests(addRequest)
				}
			}
		})
		pageNum += 1
	}

	crawler.PrintCrawlResults(source, totalProducts)

	<-worker
	done <- true
}

func getIdfDetailInfo(producturl string) (imageUrls, colors, sizes []string, inventories []*domain.InventoryDAO, description map[string]string) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.idfmall.co.kr"),
	)

	description = map[string]string{}

	colorItems := map[string]string{}
	sizeOptions := map[string]string{}
	soldoutItemNos := map[string][]string{}

	stockUrl := "https://www.idfmall.co.kr/main/exec.php?exec_file=shop/getAjaxData.php&exec=getOptionStock&item_no=@"

	c.OnHTML("#mimg_div", func(e *colly.HTMLElement) {
		imageurl := e.ChildAttr("img", "src")
		if imageurl != "" {
			imageUrls = append(imageUrls, imageurl)
		}
	})

	c.OnHTML(".image_view .imglist #product_add_image_list", func(e *colly.HTMLElement) {
		e.ForEach("div", func(_ int, el *colly.HTMLElement) {
			imageurl := el.ChildAttr("img", "src")
			if imageurl != "" {
				imageUrls = append(imageUrls, imageurl)
			}
		})
	})

	c.OnHTML(".info_view .info .info_list ul li", func(e *colly.HTMLElement) {
		if e.ChildText(".fld_title") == "색상" {
			e.ForEach(".fld_cnt ul li", func(i int, el *colly.HTMLElement) {
				if i == 0 {
					colorName := el.ChildText("a .name")
					itemNo := el.ChildAttr("a", "data")
					colorItems[itemNo] = colorName

					errorMessage := "Crawl Failed: Product Stock" + stockUrl + itemNo
					resp, err := utils.RequestRetryer(stockUrl+itemNo, utils.REQUEST_GET, utils.GetGeneralHeader(), "", errorMessage)
					if err != nil {
						return
					}
					defer resp.Body.Close()

					b, _ := ioutil.ReadAll(resp.Body)

					soldoutSizeNos := strings.Split(string(b), "@")
					soldoutItemNos[itemNo] = soldoutSizeNos
				}
			})
		} else if e.ChildText(".fld_title") == "사이즈" {
			e.ForEach(".fld_cnt ul li", func(_ int, el *colly.HTMLElement) {
				sizeName := el.ChildText("a")
				sizeNo := el.ChildAttr("a", "data")
				sizeOptions[sizeNo] = sizeName
			})
		}
	})

	c.OnHTML(".info_table tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			if !strings.Contains(el.ChildText("th"), "A/S") {
				description[el.ChildText("th")] = el.ChildText("td")
			}
		})
	})

	c.Visit(producturl)

	for colorNo, colorItem := range colorItems {
		for sizeNo, sizeOption := range sizeOptions {
			isSoludout := false
			for _, soldoutItem := range soldoutItemNos[colorNo] {
				if soldoutItem == sizeNo {
					isSoludout = true
				}
			}
			if !isSoludout {
				inventories = append(inventories, &domain.InventoryDAO{
					Size:     colorItem + " - " + sizeOption,
					Quantity: 1,
				})
			}
		}
	}

	for _, colorItem := range colorItems {
		colors = append(colors, colorItem)
	}
	for _, sizeOption := range sizeOptions {
		sizes = append(sizes, sizeOption)
	}

	return
}
