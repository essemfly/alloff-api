package seeder

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

const categoryFileSeeder = "pkg/seeder/category_classifier_seeder_220308.csv"

type CsvData struct {
	OriginalName    string
	InterimKeyName  string
	FirstAlloffKey  string
	SecondAlloffKey string
	FirstName       string
	SecondName      string
}

func MakeClassifyRules() {
	file, err := os.Open(categoryFileSeeder)
	if err != nil {
		log.Panicln("Classify seeder csv file not found")
	}
	reader := csv.NewReader(bufio.NewReader(file))
	reader.LazyQuotes = true

	rows, err := reader.ReadAll()
	if err != nil {
		log.Println("Err found in reader", err)
	}

	catNameMapper := map[string]*domain.AlloffCategoryDAO{}

	cats, _ := ioc.Repo.AlloffCategories.List(nil)
	for _, cat := range cats {
		catNameMapper[cat.Name] = cat
		catID := cat.ID.Hex()
		subcats, _ := ioc.Repo.AlloffCategories.List(&catID)
		for _, subcat := range subcats {
			catNameMapper[subcat.Name] = subcat
		}
	}

	for _, row := range rows {
		var rule map[string]string
		newRow6 := strings.ReplaceAll(row[7], "'", "\"")
		if newRow6 != "" {
			if err := json.Unmarshal([]byte(newRow6), &rule); err != nil {
				panic(err)
			}
		}

		newRule := domain.ClassifierDAO{
			BrandKeyname:    row[2],
			CategoryName:    row[3],
			AlloffCategory1: catNameMapper[row[4]],
			AlloffCategory2: catNameMapper[row[5]],
			HeuristicRules:  rule,
		}

		_, err := ioc.Repo.ClassifyRules.Upsert(&newRule)
		if err != nil {
			log.Println(err)
		}
	}

}
