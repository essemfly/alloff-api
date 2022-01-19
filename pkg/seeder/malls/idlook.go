package malls

import (
	"encoding/json"
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
)

type IdlookResponseParser struct {
	CateList []IdlookCategory `json:"mOCateList"`
}

type IdlookCategory struct {
	CATEGORY_ID      string `json:"CATEGORY_ID"`
	DISP_YN          string `json:"DISP_YN"`
	LVL              int    `json:"LVL"`
	EMP_DISP_YN      string `json:"EMP_DISP_YN"`
	MOD_ID           string `json:"MOD_ID"`
	UP_CATEGORY_ID   string `json:"UP_CATEGORY_ID"`
	REG_DATE         string `json:"REG_DATE"`
	MENU_ID          string `json:"MENU_ID"`
	DISP_NM          string `json:"DISP_NM"`
	CATEGORY_DIVN_NM string `json:"CATEGORY_DIVN_NM"`
	CATEGORY_DIVN_CD string `json:"CATEGORY_DIVN_CD"`
	SITE_CD          string `json:"SITE_CD"`
	DEL_YN           string `json:"DEL_YN"`
	MALL_DIVN_CD     string `json:"MALL_DIVN_CD"`
	BEST_VISIBLE_SEQ int    `json:"BEST_VISIBLE_SEQ"`
	CATEGORY_NM      string `json:"CATEGORY_NM"`
	VISIBLE_SEQ      int    `json:"VISIBLE_SEQ"`
	REG_ID           string `json:"REG_ID"`
}

func AddIdLook() {
	modulename := "idlook"

	url := "http://www.idlookmall.com/display/display_comm_Info.json"
	headers := map[string]string{
		"accept": "application/json",
	}

	resp, _ := utils.MakeRequest(url, utils.REQUEST_GET, headers, "")

	crawlResponse := &IdlookResponseParser{}
	json.NewDecoder(resp.Body).Decode(crawlResponse)

	brands := map[string]domain.BrandDAO{
		"9": {
			KorName:       "일레븐티",
			EngName:       "eleventy",
			KeyName:       "ELEVENTY",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ELEVENTY.png",
			Onpopular:     false,
			Description:   "컨템포러리 디자이너",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "idlook",
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/eleventy_TOP:BOTTOM.png"},
				{Label: "데님/진", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/eleventy_DENIM:JEAN.png"},
				{Label: "슈즈", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/eleventy_SHOES.png"},
			},
		},
		"T": {
			KorName:     "에센셜",
			EngName:     "ESSENTIEL",
			KeyName:     "ESSENTIEL",
			LogoImgUrl:  "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ESSENTIEL.png",
			Onpopular:   false,
			Description: "컨템포러리",
			Created:     time.Now(),
			IsOpen:      false,
			Modulename:  "idlook",
			SizeGuide:   []domain.SizeGuideDAO{},
		},
		"Q": {
			KorName:       "산드로",
			EngName:       "sandro",
			KeyName:       "SANDRO",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/SANDRO.png",
			Onpopular:     true,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "idlook",
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/sandro_TOP:BOTTOM.png"},
				{Label: "데님/진", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/sandro_DENIM:JEAN.png"},
				{Label: "슈즈", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/sandro_SHOES.png"},
			},
		},
		"4": {
			KorName:       "베르니스",
			EngName:       "Berenice",
			KeyName:       "BERENICE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/BERENICE.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "idlook",
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Berenice_TOP:BOTTOM.png"},
				{Label: "데님/진", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Berenice_DENIM:JEAN.png"},
				{Label: "슈즈", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Berenice_SHOES.png"},
			},
		},
		"Y": {
			KorName:       "마쥬",
			EngName:       "Maje",
			KeyName:       "MAJE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MAJE.png",
			Onpopular:     true,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "idlook",
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Maje_TOP:BOTTOM.png"},
				{Label: "데님/진", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Maje_DENIM:JEAN.png"},
				{Label: "슈즈", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Maje_SHOES.png"},
			},
		},
		"D02": {
			KorName:       "마리메꼬",
			EngName:       "marimekko",
			KeyName:       "MARIMEKKO",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MARIMEKKO.png",
			Onpopular:     false,
			Description:   "라이프 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "idlook",
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/marimekko_TOP:BOTTOM.png"},
			},
		},
		"W": {
			KorName:       "레니본",
			EngName:       "RENEEVON",
			KeyName:       "RENEEVON",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/RENEEVON.png",
			Onpopular:     false,
			Description:   "컨템포러리 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "idlook",
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/RENEEVON_TOP:BOTTOM.png"},
			},
		},
		"E": {
			KorName:       "끌로디피에로",
			EngName:       "CLAUDIEPIERLOT",
			KeyName:       "CLAUDIEPIERLOT",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/CLAUDIEPIERLOT.png",
			Onpopular:     false,
			Description:   "컨템포러리 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "idlook",
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/CLAUDIE+PIERLOT_TOP:BOTTOM.png"},
				{Label: "데님/진", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/CLAUDIE+PIERLOT_DENIM:JEAN.png"},
				{Label: "슈즈", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/CLAUDIE+PIERLOT_SHOES.png"},
			},
		},
		"2": {
			KorName:       "아페쎄",
			EngName:       "A.P.C.",
			KeyName:       "APC",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/APC.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "idlook",
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/A.P.C_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/A.P.C_BOTTOM.png"},
				{Label: "슈즈", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/A.P.C_SHOES.png"},
			},
		},
	}

	for _, brand := range brands {
		_, err := ioc.Repo.Brands.Upsert(&brand)
		if err != nil {
			log.Println(err)
		}
	}

	for _, dat := range crawlResponse.CateList {
		if dat.LVL != 1 {
			continue
		}
		ourCategory := false
		for key := range brands {
			if key == dat.UP_CATEGORY_ID {
				ourCategory = true
				break
			}
		}
		if !ourCategory {
			continue
		}

		if dat.DISP_NM == "악세서리" {
			continue
		}

		brandIdentifier := dat.UP_CATEGORY_ID
		category := domain.CategoryDAO{
			Name:          dat.DISP_NM,
			KeyName:       brands[brandIdentifier].KeyName + "-" + dat.DISP_NM,
			CatIdentifier: dat.CATEGORY_ID,
			BrandKeyname:  brands[brandIdentifier].KeyName,
		}

		updatedCat, err := ioc.Repo.Categories.Upsert(&category)
		if err != nil {
			log.Println(err)
		}

		source := domain.CrawlSourceDAO{
			BrandKeyname:         brands[brandIdentifier].KeyName,
			BrandIdentifier:      brandIdentifier,
			MainCategoryKey:      updatedCat.CatIdentifier,
			Category:             *updatedCat,
			CrawlUrl:             "http://www.idlookmall.com/display/outlet_product_list.do?catLvl=1&cateId=" + updatedCat.CatIdentifier,
			CrawlModuleName:      modulename,
			IsSalesProducts:      true,
			IsForeignDelivery:    false,
			PriceMarginPolicy:    "NORMAL",
			DeliveryPrice:        0,
			EarliestDeliveryDays: 2,
			LatestDeliveryDays:   7,
			DeliveryDesc:         nil,
			RefundAvailable:      true,
			ChangeAvailable:      true,
			RefundFee:            5000,
			ChangeFee:            5000,
			RefundRoughFee:       5500,
			ChangeRoughFee:       5500,
		}

		_, err = ioc.Repo.CrawlSources.Upsert(&source)
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("IDLOOK mall brands, categories & sources are added")
}
