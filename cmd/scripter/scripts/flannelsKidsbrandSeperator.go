package scripts

import (
	"fmt"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

// SeparateFlannelsProductsBrand :
// 1. listing all products which needs to be separated by kids brands
// 2. if pd has kids brands, update brand of that product
// 3. if pd has not kids brands, do nothing
func SeparateFlannelsProductsBrand() {
	errors := []map[string]string{}

	err := kidsBrandsAdder()
	if err != nil {
		log.Panic("PANIC: can not upsert kids brands")
	}

	filters := bson.M{"source.crawlmodulename": "flannels", "producttype": []string{"KIDS"}}
	pds, _, err := ioc.Repo.ProductMetaInfos.List(0, 10000, filters, bson.M{})

	for _, pd := range pds {
		switch pd.Brand.KeyName {
		case "BALMAIN", "BURBERRY", "BARBOUR_INTERNATIONAL", "BARBOUR", "CANADAGOOSE", "CHLOE", "EMPORIOARMANI", "FREDPERRY",
			"FENDI", "GIVENCHY", "KENZO", "LACOSTE", "MOSCHINO", "OFFWHITE", "POLORALPHLAUREN", "PALMANGELS", "STONEISLAND",
			"STELLAMCCARTNEY", "UGG":
			kidsBrandKeyname := pd.Brand.KeyName + "KIDS"
			newBrand, err := ioc.Repo.Brands.GetByKeyname(kidsBrandKeyname)
			if err != nil {
				errorDetail := fmt.Sprintf("err : %v <- from %s (brands : %s)", err.Error(), pd.ID.Hex(), pd.Brand.KeyName)
				errors = append(errors, map[string]string{"type": "get kids brands", "detail": errorDetail})
			}

			pd.Brand = newBrand
			_, err = productinfo.Update(pd)
			if err != nil {
				errDetail := fmt.Sprintf("err : %v <- from %s (brands : %s)", err.Error(), pd.ID.Hex(), pd.Brand.KeyName)
				errors = append(errors, map[string]string{"type": "update kids brands", "detail": errDetail})
			}
		default:
			continue
		}
	}

	for _, error := range errors {
		log.Println(error)
	}
}

// kidsBrandsAdder: Upsert new kids brands in https://docs.google.com/spreadsheets/d/1vrzlMJJCSpO2V8HygtdPJVcYp0ectSuVFRfpump_LGE/edit#gid=1589106977
func kidsBrandsAdder() error {
	brands := map[string]*domain.BrandDAO{
		"balmain-kids": {
			KorName:       "발망 키즈",
			EngName:       "BALMAIN KIDS",
			KeyName:       "BALMAINKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"burberry-kids": {
			KorName:       "버버리 키즈",
			EngName:       "BURBERRY KIDS",
			KeyName:       "BURBERRYKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"barbour-international-kids": {
			KorName:       "바버인터네셔널 키즈",
			EngName:       "Barbour International Kids",
			KeyName:       "BARBOUR_INTERNATIONALKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"barbour-kids": {
			KorName:       "바버 키즈",
			EngName:       "Barbour Kids",
			KeyName:       "BARBOURKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"canadagoose-kids": {
			KorName:       "캐나다구스 키즈",
			EngName:       "CANADAGOOSE KIDS",
			KeyName:       "CANADAGOOSEKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"chloe-kids": {
			KorName:       "끌로에 키즈",
			EngName:       "Chloe Kids",
			KeyName:       "CHLOEKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"emporioarmani-kids": {
			KorName:       "엠포리오아르마니 키즈",
			EngName:       "EMPORIOARMANI KIDS",
			KeyName:       "EMPORIOARMANIKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"fred-perry-kids": {
			KorName:       "프레드페리 키즈",
			EngName:       "FRED PERRY KIDS",
			KeyName:       "FREDPERRYKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"fendi-kids": {
			KorName:       "펜디 키즈",
			EngName:       "Fendi Kids",
			KeyName:       "FENDIKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"givenchy-kids": {
			KorName:       "지방시 키즈",
			EngName:       "Givenchy Kids",
			KeyName:       "GIVENCHYKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"kenzo-kids": {
			KorName:       "겐조 키즈",
			EngName:       "Kenzo Kids",
			KeyName:       "KENZOKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"lacoste-kids": {
			KorName:       "라코스테 키즈",
			EngName:       "Lacoste Kids",
			KeyName:       "LACOSTEKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"moschino-kids": {
			KorName:       "모스키노 키즈",
			EngName:       "MOSCHINO KIDS",
			KeyName:       "MOSCHINOKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"off-white-kids": {
			KorName:       "오프화이트 키즈",
			EngName:       "Off-White Kids",
			KeyName:       "OFFWHITEKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"polo-ralph-lauren-kids": {
			KorName:       "폴로 랄프로렌 키즈",
			EngName:       "Polo Ralph Lauren Kids",
			KeyName:       "POLORALPHLAURENKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"palmangels-kids": {
			KorName:       "팜엔젤스 키즈",
			EngName:       "PALMANGELS KIDS",
			KeyName:       "PALMANGELSKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"stone-island-junior": {
			KorName:       "스톤아일랜드 주니어",
			EngName:       "Stone Island Junior",
			KeyName:       "STONEISLANDJUNIOR",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"stella-mccartney-kids": {
			KorName:       "스텔라맥카트니 키즈",
			EngName:       "Stella Mccartney Kids",
			KeyName:       "STEELLAMCCARTNEYKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
		"ugg-kids": {
			KorName:       "어그 키즈",
			EngName:       "UGG KIDS",
			KeyName:       "UGGKIDS",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
		},
	}
	for _, brand := range brands {
		_, err := ioc.Repo.Brands.Upsert(brand)
		if err != nil {
			return fmt.Errorf("error on %s", brand.KeyName)
		}
	}
	return nil
}
