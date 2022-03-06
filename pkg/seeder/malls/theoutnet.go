package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddTheoutnet() {
	modulename := "theoutnet"
	categories := map[string]string{
		"상의":   "/clothing/tops",
		"코트":   "/clothing/coats",
		"자켓":   "/clothing/jackets",
		"니트":   "/clothing/knitwear",
		"드레스":  "/clothing/dresses",
		"스커트":  "/clothing/skirts",
		"캐시미어": "/clothing/cashmere",
		"라운지":  "/clothing/evening",
		"가방":   "/bags",
		"구두":   "/shoes",
		"악세서리": "/accessories",
	}

	brands := map[string]domain.BrandDAO{
		"sandro": {
			KorName:       "산드로",
			EngName:       "sandro",
			KeyName:       "SANDRO",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/SANDRO.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"maje": {
			KorName:       "마쥬",
			EngName:       "Maje",
			KeyName:       "MAJE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MAJE.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"tory-burch": {
			KorName:       "토리버치",
			EngName:       "TORYBURCH",
			KeyName:       "TORYBURCH",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/TORYBURCH.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"burberry": {
			KorName:       "버버리",
			EngName:       "BURBERRY",
			KeyName:       "BURBERRY",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/BURBERRY.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"ganni": {
			KorName:       "가니",
			EngName:       "GANNI",
			KeyName:       "GANNI",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/GANNI.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"isabel-marant": {
			KorName:       "이자벨마랑",
			EngName:       "Isabel Marant",
			KeyName:       "ISABELMARANT",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ISABELMARANT.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"isabel-marant-etoile": {
			KorName:       "이자벨마랑 에뚜왈",
			EngName:       "Isabel Marant Etoile",
			KeyName:       "ISABELMARANTETOILE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ISABELMARANTETOILE.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"vince": {
			KorName:       "빈스",
			EngName:       "Vince",
			KeyName:       "VINCE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/VINCE.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"tod-s": {
			KorName:       "토즈",
			EngName:       "Tod's",
			KeyName:       "TODS",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/TODS.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"stella-mccartney": {
			KorName:       "스텔라 맥카트니",
			EngName:       "Stella McCartney",
			KeyName:       "STELLAMCCARTNEY",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/STELLAMCCARTNEY.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"marni": {
			KorName:       "마르니",
			EngName:       "MARNI",
			KeyName:       "MARNI",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MARNI.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"chloe": {
			KorName:       "끌로에",
			EngName:       "Chloé",
			KeyName:       "CHLOE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/CHLOE.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"apc": {
			KorName:       "아페쎄",
			EngName:       "A.P.C",
			KeyName:       "APC",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/APC.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"acne-studios": {
			KorName:       "아크네 스튜디오",
			EngName:       "Acne Studios",
			KeyName:       "ACNESTUDIOS",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ACNESTUDIOS.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"alexander-mcqueen": {
			KorName:       "알렉산더 맥퀸",
			EngName:       "Alexander McQueen",
			KeyName:       "ALEXANDERMCQUEEN",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ALEXANDERMCQUEEN.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"bash": {
			KorName:       "바앤쉬",
			EngName:       "ba&sh",
			KeyName:       "BASH",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/BASH.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"balenciaga": {
			KorName:       "발렌시아가",
			EngName:       "Balenciaga",
			KeyName:       "BALENCIAGA",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/BALENCIAGA.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"brunello-cucinelli": {
			KorName:       "브루넬로 쿠치넬리",
			EngName:       "Brunello Cucinelli",
			KeyName:       "BRUNELLOCUCINELLI",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/BRUNELLOCUCINELLI.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"balmain": {
			KorName:       "발망",
			EngName:       "Balmain",
			KeyName:       "BALMAIN",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/BALMAIN.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"calvin-klein": {
			KorName:       "캘빈클라인",
			EngName:       "Calvin Klein",
			KeyName:       "CALVINKLEIN",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/CALVINKLEIN.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"coach": {
			KorName:       "코치",
			EngName:       "Coach",
			KeyName:       "COACH",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/COACH.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"claudie-pierlot": {
			KorName:       "끌로디 피에로",
			EngName:       "Claudie Pierlot",
			KeyName:       "CLAUDIEPIERLOT",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/CLAUDIEPIERLOT.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"dkny": {
			KorName:       "디케이앤와이",
			EngName:       "DKNY",
			KeyName:       "DKNY",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/DKNY.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"dries-van-noten": {
			KorName:       "드리스 반 노튼",
			EngName:       "Dries Van Noten",
			KeyName:       "DRIESVANNOTEN",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/DRIESVANNOTEN.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"dolce-gabbana": {
			KorName:       "돌체&가바나",
			EngName:       "Dolce & Gabbana",
			KeyName:       "DOLCEGABBANA",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/DOLCEGABBANA.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"etro": {
			KorName:       "에트로",
			EngName:       "Etro",
			KeyName:       "ETRO",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ETRO.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"iro": {
			KorName:       "이로",
			EngName:       "IRO",
			KeyName:       "IRO",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/IRO.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"ienki-ienki": {
			KorName:       "옌키옌키",
			EngName:       "IENKI IENKI",
			KeyName:       "IENKIIENKI",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/IENKIIENKI.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"jimmy-choo": {
			KorName:       "지미추",
			EngName:       "JImmy Choo",
			KeyName:       "JIMMYCHOO",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/JIMMYCHOO.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"jil-sander": {
			KorName:       "질센더",
			EngName:       "Jil Sander",
			KeyName:       "JILSANDER",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/JILSANDER.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"joseph": {
			KorName:       "조셉",
			EngName:       "JOSEPH",
			KeyName:       "JOSEPH",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/JOSEPH.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"kenzo": {
			KorName:       "겐조",
			EngName:       "KENZO",
			KeyName:       "KENZO",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/KENZO.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"maison-margiela": {
			KorName:       "메종 마르지엘라",
			EngName:       "Maison Margiela",
			KeyName:       "MAISONMARGIELA",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MAISONMARGIELA.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"msgm": {
			KorName:       "MSGM",
			EngName:       "MSGM",
			KeyName:       "MSGM",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MSGM.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"mm6-maison-margiela": {
			KorName:       "MM6 메종 마르지엘라",
			EngName:       "MM6 MAISON MARGIELA",
			KeyName:       "MM6MAISONMARGIELA",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MM6MAISONMARGIELA.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"paul-smith": {
			KorName:       "폴스미스",
			EngName:       "Paul Smith",
			KeyName:       "PAULSMITH",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/PAULSMITH.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"rick-owens": {
			KorName:       "릭 오웬스",
			EngName:       "Rick Owens",
			KeyName:       "RICKOWENS",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/RICKOWENS.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"rag-bone": {
			KorName:       "랙앤본",
			EngName:       "rag & bone",
			KeyName:       "RAGANDBONE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/RAGANDBONE.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"st-john": {
			KorName:       "세인트 존",
			EngName:       "St.John",
			KeyName:       "STJOHN",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/STJOHN.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"toteme": {
			KorName:       "토템",
			EngName:       "Toteme",
			KeyName:       "TOTEME",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/TOTEME.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"valentino": {
			KorName:       "발렌티노",
			EngName:       "Valentino",
			KeyName:       "VALENTINO",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/VALENTINO.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
		"vanessa-bruno": {
			KorName:       "바네사브루노",
			EngName:       "Vanessabruno",
			KeyName:       "VANESSABRUNO",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/VANESSABRUNO.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
		},
	}

	for shopID, brand := range brands {
		upsertedBrand, err := ioc.Repo.Brands.Upsert(&brand)
		if err != nil {
			log.Println(err)
		}

		for key, val := range categories {
			category := domain.CategoryDAO{
				Name:          key,
				KeyName:       brand.KeyName + "-" + key,
				CatIdentifier: val,
				BrandKeyname:  upsertedBrand.KeyName,
			}

			updatedCat, err := ioc.Repo.Categories.Upsert(&category)
			if err != nil {
				log.Println(err)
			}

			source := domain.CrawlSourceDAO{
				BrandKeyname:      upsertedBrand.KeyName,
				BrandIdentifier:   shopID,
				MainCategoryKey:   updatedCat.KeyName,
				Category:          *updatedCat,
				CrawlUrl:          "https://www.theoutnet.com/api/yoox/ton/search/resources/store/theoutnet_DE/productview/byCategory",
				CrawlModuleName:   modulename,
				IsSalesProducts:   true,
				IsForeignDelivery: true,
				PriceMarginPolicy: "THEOUTNET",
				DeliveryPrice:     0,
			}

			_, err = ioc.Repo.CrawlSources.Upsert(&source)
			if err != nil {
				log.Println(err)
			}
		}
	}
	log.Println("Theoutnet mall brands, categories & sources are added")
}