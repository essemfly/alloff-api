package scripts

import (
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/pkg/seeder"
)

func AddAlloffCategory() {
	catDaos := seeder.LoadAlloffCats()

	for _, catDao := range catDaos {
		ioc.Repo.AlloffCategories.Upsert(catDao)
	}

	seeder.MakeClassifyRules()
}
