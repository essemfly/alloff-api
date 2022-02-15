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
	// scripts.AddProductGroups()
	// scripts.AddMockOrders()
	// scripts.AddMockHomeItems()
	// scripts.InsertDiffNotification()
	// scripts.TestAlimtalk()
	// scripts.AddTimedealNotification()
	// scripts.SendNotification()
	scripts.AddMockExhibitions()
	scripts.AddMockHomeTabs()
	scripts.AddMockTopBanners()
}
