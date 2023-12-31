package seeder

import (
	"log"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func NewSizeMappingPolicy() {
	alloffSizes, _, err := ioc.Repo.AlloffSizes.List(0, 1000)
	if err != nil {
		log.Panic("1")
	}

	for _, alloffSize := range alloffSizes {
		sizes := []string{}
		switch alloffSize.AlloffSizeName {
		case "44":
			sizes = []string{"XP", "P/S", "P", "XXS", "XS", "DE32", "IT36", "FR34", "US32", "US0", "US2", "UK4", "EU32", "EU34", "EU0", "EU00", "XXS/XS"}
		case "55":
			sizes = []string{"P/S", "S", "DE34", "DE36", "IT38", "IT40", "FR36", "FR38", "US34", "US36", "US4", "US6", "UK6", "UK8", "EU36", "EU38", "EU1", "EU2", "S/M"}
		case "66":
			sizes = []string{"M/L", "M", "DE36", "DE38", "DE40", "IT40", "IT42", "IT44", "FR38", "FR40", "FR42", "US36", "US38", "US40", "US6", "US8", "US10", "UK8", "UK10", "UK12", "EU38", "EU40", "EU42", "EU3", "EU4", "S/M", "US8/10"}
		case "77":
			sizes = []string{"M/L", "L", "DE40", "DE42", "DE44", "IT44", "IT46", "IT48", "FR42", "FR44", "FR46", "US40", "US42", "US44", "US10", "US12", "US14", "UK12", "UK14", "UK16", "EU42", "EU44", "EU46", "EU4", "L/XL", "US10/12"}
		case "88":
			sizes = []string{"XL", "XXL", "DE44", "DE46", "DE48", "IT48", "IT50", "IT52", "FR46", "FR48", "FR50", "US44", "US46", "US48", "US14", "US16", "US18", "UK16", "UK18", "UK20", "EU46", "EU48", "EU50", "EU5", "L/XL"}
		case "99(이상)":
			sizes = []string{"XXL", "DE48", "IT52", "FR50", "US48", "US18", "UK20", "EU50", "L/XL"}
		case "24-25":
			sizes = []string{"XXS", "XS", "IT36", "IT38", "FR32", "FR34", "US0", "US2", "UK4", "UK6", "EU32", "EU34", "24", "25"}
		case "26-27":
			sizes = []string{"S", "IT40", "FR36", "US4", "UK8", "EU36", "26", "27"}
		case "28-29":
			sizes = []string{"M", "L", "IT42", "IT44", "FR38", "FR40", "US6", "US8", "UK10", "UK12", "EU38", "EU40", "28", "29"}
		case "30-31":
			sizes = []string{"M", "L", "IT42", "IT44", "IT46", "FR38", "FR40", "FR42", "US6", "US8", "US10", "UK10", "UK12", "UK14", "EU38", "EU40", "EU42", "30", "31"}
		case "32-33":
			sizes = []string{"M", "L", "IT44", "IT46", "FR40", "FR42", "US8", "US10", "UK12", "UK14", "EU40", "EU42", "32", "33"}
		case "34(이상)":
			sizes = []string{"XL", "XXL", "IT48", "IT50", "FR44", "FR46", "FR48", "US12", "US14", "US16", "US18", "UK16", "UK18", "EU46", "EU48", "EU50", "34", "35", "36", "37"}
		case "220":
			sizes = []string{"UK2", "EU35", "FR36", "US5"}
		case "225":
			sizes = []string{"UK2.5", "EU35.5", "FR36.5", "US5.5"}
		case "230":
			sizes = []string{"UK3", "EU36", "FR37", "US6"}
		case "235":
			sizes = []string{"UK3.5", "EU36.5", "FR37.5", "US6.5"}
		case "240":
			sizes = []string{"UK4", "EU37", "FR38", "US7"}
		case "245":
			sizes = []string{"UK4.5", "EU37.5", "FR38.5", "US7.5"}
		case "250":
			if alloffSize.ProductType[0] == domain.Female {
				sizes = []string{"UK5", "EU38", "FR39", "US8"}
			} else {
				sizes = []string{"IT40", "US7", "UK6"}
			}
		case "255":
			if alloffSize.ProductType[0] == domain.Female {
				sizes = []string{"UK5.5", "EU38.5", "FR39.5", "US8.5"}
			} else {
				sizes = []string{"IT40.5", "US7.5", "UK6.5"}
			}
		case "260":
			if alloffSize.ProductType[0] == domain.Female {
				sizes = []string{"UK6", "EU39", "FR40", "US9"}
			} else {
				sizes = []string{"IT41", "US8", "UK7"}
			}
		case "265(이상)":
			sizes = []string{"UK6.5", "UK7", "EU39.5", "EU40", "FR40.5", "FR41", "US9.5", "US10"}
		case "265":
			sizes = []string{"IT41.5", "US8.5", "UK7.5"}
		case "270":
			sizes = []string{"IT42", "US9", "UK8"}
		case "275":
			sizes = []string{"IT42.5", "US9.5", "UK8.5"}
		case "280(이상)":
			sizes = []string{"IT43", "US10", "UK9"}
		case "90":
			sizes = []string{"XS", "FR44", "IT44", "EU44", "US34", "UK34", "UK1", "EU0"}
		case "95":
			sizes = []string{"S", "FR46", "IT46", "EU46", "US36", "UK36", "UK2", "EU1"}
		case "100":
			sizes = []string{"M", "FR48", "IT48", "EU48", "US38", "UK38", "UK3", "EU2"}
		case "105":
			sizes = []string{"L", "FR50", "IT50", "EU50", "US40", "UK40", "UK4", "EU3"}
		case "110":
			sizes = []string{"XL", "FR52", "IT52", "EU52", "US42", "UK42", "UK5", "EU4"}
		case "115":
			sizes = []string{"XXL", "FR54", "IT54", "EU54", "US44", "UK44", "UK6", "EU5"}
		case "120(이상)":
			sizes = []string{"XXXL", "US46", "UK46", "EU6"}
		case "28":
			sizes = []string{"IT44", "US32", "US34", "UK34", "FR36", "28"}
		case "30":
			sizes = []string{"IT46", "US36", "UK36", "FR38", "30"}
		case "32":
			sizes = []string{"IT48", "US38", "UK38", "FR40", "32"}
		case "34":
			sizes = []string{"IT50", "US40", "UK40", "FR42", "34"}
		case "36":
			sizes = []string{"IT52", "US42", "UK42", "FR44", "36"}
		case "38(이상)":
			sizes = []string{"IT54", "US44", "US46", "UK44", "UK46", "FR46", "38", "40"}
		case "145":
			sizes = []string{"US8.5", "UK8", "EU25.5"}
		case "150":
			sizes = []string{"US9", "UK8.5", "EU26"}
		case "155":
			sizes = []string{"US9.5", "UK9", "EU26.5"}
		case "160":
			sizes = []string{"US10", "UK9.5", "EU27"}
		case "165":
			sizes = []string{"US10.5", "UK10", "EU27.5"}
		case "170":
			sizes = []string{"US11", "UK10.5", "EU28"}
		case "175":
			sizes = []string{"US11.5", "UK11", "EU28.5"}
		case "180":
			sizes = []string{"US12", "UK11.5", "EU29"}
		case "185":
			sizes = []string{"US12.5", "UK12", "EU30"}
		case "190":
			sizes = []string{"US13", "UK12.5", "EU31"}
		case "195":
			sizes = []string{"US13.5", "UK13", "EU31.5"}
		case "200":
			sizes = []string{"US14", "UK13.5", "EU32"}
		case "205":
			sizes = []string{"US14.5", "UK14", "EU33"}
		case "210(이상)":
			sizes = []string{"US15", "UK14.5", "EU33.5"}
		case "FREE":
			sizes = []string{"-", "Uni", "Unica", "one size"}
		case "3-6개월":
			sizes = []string{"3M", "6M"}
		case "6-12개월":
			sizes = []string{"12M", "6M", "9M"}
		case "12-18개월":
			sizes = []string{"18M", "12M", "15M", "1Y", "1A", "1.5A", "1.5Y"}
		case "18-24개월":
			sizes = []string{"24M", "2Y", "2A", "18M", "21M", "1.5-2Y"}
		case "2-3세":
			sizes = []string{"2Y", "3Y", "2A", "3A"}
		case "3-4세":
			sizes = []string{"3Y", "4Y", "4A", "3A"}
		case "4-5세":
			sizes = []string{"4Y", "4A", "5A", "5Y"}
		case "5-6세":
			sizes = []string{"5A", "5Y", "6A", "6Y"}
		case "6-7세":
			sizes = []string{"6Y", "6A", "7A", "7Y"}
		case "7-8세":
			sizes = []string{"7A", "7Y", "8A", "8Y"}
		case "8-9세":
			sizes = []string{"8Y", "8A", "9A", "9Y"}
		case "9-10세":
			sizes = []string{"9A", "9Y", "10A", "10Y"}
		case "11-12세(이상)":
			sizes = []string{"12A", "12Y", "14A", "14Y", "16A", "16Y"}
		}

		alloffSize.Sizes = sizes
		_, err := ioc.Repo.AlloffSizes.Upsert(alloffSize)
		if err != nil {
			log.Panic(err)
		}
	}
}

func AssignProductsInventoryOfNewSizes() {
	sizes := []string{"XXS/XS", "S/M", "US8/10", "L/XL", "US10/12", "28", "30", "32", "34", "36", "38", "40", "12M", "3M", "6M", "18M", "9M", "12M", "15M", "1Y", "1A", "1.5A", "1.5Y", "2A", "18M", "21M", "1.5-2Y", "2A", "3A", "3A", "12A", "12Y", "14A", "14Y", "16A", "16Y"}
	//sizes := []string{"24M"}

	sizeQueries := []bson.M{}
	for _, size := range sizes {
		sizeQueries = append(sizeQueries, bson.M{"inventory.size": size})
	}
	filter := bson.M{
		"$or": sizeQueries,
	}

	pds, _, err := ioc.Repo.ProductMetaInfos.List(0, 100000, filter, bson.M{})
	if err != nil {
		config.Logger.Error("error occurred on get list of products : ", zap.Error(err))
	}

	for _, pd := range pds {
		productinfo.AssignProductsInventory(pd)
	}
}
