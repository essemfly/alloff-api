package malls

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/crawler"
)

func CrawlIdLook(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.idlookmall.com"),
	)

	totalProducts := 0
	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	c.OnHTML(".item_box", func(e *colly.HTMLElement) {
		originalPriceStr := e.ChildText(".info_area .item_price .before")
		discountedPriceStr := e.ChildText(".info_area .item_price .sellprc")
		originalPrice := utils.ParsePriceString(originalPriceStr)
		discountedPrice := utils.ParsePriceString(discountedPriceStr)
		productID := e.ChildAttr(".info_area a", "href")
		productUrl := "http://idlookmall.com" + productID

		images, colors, sizes, inventories, description := getIdLookDetail(productUrl)
		title := e.ChildAttr(".info_area a", "title")

		askey := ""
		for k, v := range description {
			if strings.Contains(k, "A/S") {
				askey = k
			}
			allInfo := v
			allInfo = strings.ReplaceAll(allInfo, "\n", "")
			allInfo = strings.ReplaceAll(allInfo, "\t", "")
			allInfo = strings.Trim(allInfo, " ")
			description[k] = allInfo
		}

		delete(description, askey)

		addRequest := crawler.ProductsAddRequest{
			Brand:         brand,
			Source:        source,
			ProductID:     productID,
			ProductName:   title,
			ProductUrl:    productUrl,
			Images:        images,
			Sizes:         sizes,
			Inventories:   inventories,
			Colors:        colors,
			Description:   description,
			OriginalPrice: float32(originalPrice),
			SalesPrice:    float32(discountedPrice),
			CurrencyType:  domain.CurrencyKRW,
		}

		totalProducts += 1
		crawler.AddProduct(addRequest)
	})

	c.OnHTML(".paging", func(e *colly.HTMLElement) {
		nextPageStr := e.ChildAttr(".next", "onclick")
		if nextPageStr != "" {
			s1 := strings.Trim(nextPageStr, "javascript:goPage(")
			s2 := strings.Trim(s1, ")")
			s3 := strings.Trim(s2, "/'")
			nextPage, _ := strconv.Atoi(s3)
			url := source.CrawlUrl + "&currentPage=" + strconv.Itoa(nextPage)
			c.Visit(url)
		}
	})
	c.Visit(source.CrawlUrl)

	crawler.WriteCrawlResults(source, totalProducts)

	<-worker
	done <- true
}

func getIdLookDetail(productUrl string) (imageUrls, colors, sizes []string, inventories []domain.InventoryDAO, description map[string]string) {
	c := colly.NewCollector(
		colly.AllowedDomains("idlookmall.com"),
	)

	description = map[string]string{}

	c.OnHTML(".thumb_area > .thumb_wrap:first-child .main_thumb .thumb_imgBox", func(e *colly.HTMLElement) {
		e.ForEach(".slide-container", func(_ int, el *colly.HTMLElement) {
			imageCssCode := el.ChildAttr(".img_box", "style")
			imageCssSplit := strings.Split(imageCssCode, "'")
			if len(imageCssSplit) > 1 {
				imageUrls = append(imageUrls, imageCssSplit[1])
			}
		})
	})

	c.OnHTML(".product_detail_area > .defaultBase", func(e *colly.HTMLElement) {
		e.ForEach(".goodsoption_size label", func(_ int, el *colly.HTMLElement) {
			size := el.ChildText("span")
			attrs := el.Attr("class")
			sizes = append(sizes, size)
			if !strings.Contains(attrs, "disabled") {
				inventories = append(inventories, domain.InventoryDAO{
					Size:     size,
					Quantity: 10,
				})
			}
		})
	})

	c.OnHTML(".product_detail_area > .list_acco ul li", func(e *colly.HTMLElement) {
		if strings.Contains(e.ChildText("a"), "상품필수정보") {
			keys := []string{}

			e.ForEach(".list_info2 dt", func(_ int, el *colly.HTMLElement) {
				keys = append(keys, el.Text)
			})
			e.ForEach(".list_info2 dd", func(i int, el *colly.HTMLElement) {
				description[keys[i]] = el.Text
			})
		}
	})

	c.Visit(productUrl)
	return
}
