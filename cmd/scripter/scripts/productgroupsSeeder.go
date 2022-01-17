package scripts

import (
	"log"

	"github.com/lessbutter/alloff-api/pkg/seeder"
)

func AddProductGroupsSeeder() {
	log.Println("Add Dummy Product Groups")
	seeder.AddProductGroups()
}
