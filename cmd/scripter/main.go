package main

import (
	"github.com/lessbutter/alloff-api/cmd"
	"github.com/lessbutter/alloff-api/cmd/scripter/scripts"
)

func main() {
	cmd.SetBaseConfig()

	// scripts.AddAlloffCategory()
	// scripts.AddBrandsSeeder()
	// scripts.AddClassifyRules()
	// scripts.UpdateBrands()
	// scripts.UpdateHomeitems()
	// scripts.AddProductDiffNotificaton()
	// scripts.ConfirmOrders()
	// scripts.SendNotification()
	// scripts.AddProductGroups()
	// scripts.AddMockOrders()
	// scripts.AddMockHomeItems()
	// scripts.InsertDiffNotification()
	// scripts.TestAlimtalk()
	title := "MaxMara(Intrend) 시즌오프:짠:"
	message := "(광고) 품절임박! 시즌오프 세일:폭탄:"
	timedealID := "61a9d41dee2c39d09bebeedb"
	scripts.AddTimedealNotification(timedealID, title, message)
}
