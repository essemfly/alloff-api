package scripts

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/pkg/classifier"
	"github.com/lessbutter/alloff-api/pkg/seeder"
)

func RenewAlloffCategories() {

	// 그전에 Remove alloff cats, & classify rules

	log.Println("Add alloff categories")
	seeder.AddNewAlloffCats()
	log.Println("Add new classify rules")
	AddClassifyRules()

	offset, limit := 0, 1000
	_, cnt, err := ioc.Repo.Products.List(offset, limit, nil, nil)
	log.Println("cnt", cnt)
	if err != nil {
		log.Println("err occured inrenew alloff cats", err)
	}

	for cnt > limit+offset {
		log.Println("current offset", offset, cnt)
		pds, _, _ := ioc.Repo.Products.List(offset, limit, nil, nil)
		for _, pd := range pds {
			alloffCat := classifier.GetAlloffCategory(pd)
			pd.UpdateAlloffCategory(alloffCat)
			_, err := ioc.Repo.Products.Upsert(pd)
			if err != nil {
				log.Println("err occured in upsert products", err)
			}
		}
		offset = offset + limit
	}
}
