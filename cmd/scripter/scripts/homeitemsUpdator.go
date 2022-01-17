package scripts

import (
	"log"

	"github.com/lessbutter/alloff-api/pkg/homeitem"
)

func UpdateHomeitems() {
	log.Println("Update Home Items")
	homeitem.UpdateHomeItems()
}
