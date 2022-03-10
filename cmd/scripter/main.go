package main

import (
	"fmt"

	"github.com/lessbutter/alloff-api/cmd"
	"github.com/lessbutter/alloff-api/cmd/scripter/scripts"
)

var (
	GitInfo   = "no info"
	BuildTime = "no datetime"
	Env       = "local"
)

func main() {
	fmt.Println("Git commit information: ", GitInfo)
	fmt.Println("Build date, time: ", BuildTime)

	cmd.SetBaseConfig(Env)

	// scripts.AddAlloffCategory()
	scripts.AddBrandsSeeder()
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
	// product.MakeSnapshot()
	// scripts.RenewAlloffCategories()
	// scripts.AddMockExhibitions()
	// scripts.AddMockHomeTabs()
	// scripts.AddMockTopBanners()
	// offset, limit := 0, 1000
	// _, cnt, err := ioc.Repo.Products.List(0, 10, bson.M{}, nil)
	// if err != nil {
	// 	log.Println("err", err)
	// }
	// log.Println("CNT", cnt)
	// for cnt > offset+limit {
	// 	log.Println("offset", offset)
	// 	pds, _, err := ioc.Repo.Products.List(offset, limit, bson.M{}, nil)
	// 	if err != nil {
	// 		log.Println("err", err)
	// 	}
	// 	for _, pd := range pds {
	// 		pd.OriginalPrice = int(pd.ProductInfo.Price.OriginalPrice)
	// 		_, err = ioc.Repo.Products.Upsert(pd)
	// 		if err != nil {
	// 			log.Println("upsert err", err)
	// 		}
	// 	}
	// 	offset += limit
	// }

	// tag1 := []string{"Premium Outlet", "Daily", "Contemporary"}
	// itemID := "6214de70ec5594445b285202"
	// AddTags(tag1, itemID)

	// tag2 := []string{"주목해야 하는 브랜드", "Outlet Brand Day"}
	// itemID2 := "6214f3f8a1466e95e4bc0695"
	// AddTags(tag2, itemID2)

	// tag3 := []string{"트위드", "원피스", "자켓"}
	// itemID3 := "6214facfa1466e95e4bc0697"
	// AddTags(tag3, itemID3)

	// tag4 := []string{"Premium Outlet", "S/S", "Daily"}
	// itemID4 := "6214fb77a1466e95e4bc0699"
	// AddTags(tag4, itemID4)

	// tag5 := []string{"Premium Outlet", "S/S", "Daily"}
	// itemID5 := "6214fbb5a1466e95e4bc069a"
	// AddTags(tag5, itemID5)

	// pgs, _ := ioc.Repo.ProductGroups.ListTimedeals(0, 100, true)
	// log.Println("개수", len(pgs))
	// for _, pg := range pgs {
	// 	if pg.ID.Hex() != "6214f623a1466e95e4bc0696" {
	// 		pg.FinishTime = time.Now()
	// 		_, err := ioc.Repo.ProductGroups.Upsert(pg)
	// 		if err != nil {
	// 			log.Println("err on timedeals", err)
	// 		}
	// 	}
	// }

	// RemoveManuelAddedProducts()
	// product.MakeSnapshot()
	// ProductSyncer()
	// HomeTabSyncer()
	// AddProductInSize()
	// item, _ := ioc.Repo.HomeTabItems.Get("6214fbb5a1466e95e4bc069a")
	// item.Exhibitions[0].FinishTime = time.Date(
	// 	2022, 2, 28, 14, 59, 59, 0, time.UTC)
	// _, err := ioc.Repo.HomeTabItems.Update(item)
	// if err != nil {
	// 	log.Println("err", err)
	// }

}

// func AddTags(tags []string, itemID string) {
// 	item, err := ioc.Repo.HomeTabItems.Get(itemID)
// 	if err != nil {
// 		log.Println("er", err)
// 	}
// 	item.Tags = tags
// 	_, err = ioc.Repo.HomeTabItems.Update(item)
// 	if err != nil {
// 		log.Println("er2", err)
// 	}

// }

// func RemoveManuelAddedProducts() {
// 	brandKeynames := []string{"THEORY", "LACOSTE", "MAJE", "SANDRO", "MAISONKITSUNE", "MOOSEKNUCKLES", "BURBERRY", "BCBG", "VINCE", "TOMMYHILFIGERW", "RLR", "CLUBMONACO"}

// 	offset, limit := 0, 1000
// 	for _, keyname := range brandKeynames {
// 		filter := bson.M{"productinfo.brand.keyname": keyname, "productinfo.source.crawlmodulename": "manuel"}
// 		pds, cnt, err := ioc.Repo.Products.List(offset, limit, filter, nil)
// 		if err != nil {
// 			log.Println("err in pd", err)
// 		}
// 		log.Println("pd : " + keyname + " : " + strconv.Itoa(cnt))
// 		for _, pd := range pds {
// 			pd.Removed = true
// 			_, err = ioc.Repo.Products.Upsert(pd)
// 			if err != nil {
// 				log.Println("err in upsert pd", err)
// 			}
// 		}
// 	}
// }

// func ProductSyncer() {
// 	offset, limit := 0, 100
// 	pgs, err := ioc.Repo.ProductGroups.ListExhibitionPg(offset, limit)
// 	if err != nil {
// 		log.Println("pg list error", err)
// 	}

// 	for idx, pg := range pgs {
// 		log.Println("IDX", idx)
// 		ProductGroupSyncer(pg)
// 	}
// }

// func ProductGroupSyncer(pgDao *domain.ProductGroupDAO) {
// 	for _, pd := range pgDao.Products {
// 		newPd, err := ioc.Repo.Products.Get(pd.Product.ID.Hex())
// 		if err != nil {
// 			log.Println("err not found product", newPd.ID.Hex())
// 		}
// 		pd.Product = newPd
// 	}

// 	_, err := ioc.Repo.ProductGroups.Upsert(pgDao)
// 	if err != nil {
// 		log.Println("err in upsert pgdao")
// 	}
// }

// func ExhibitionSyncer() {
// 	offset, limit := 0, 100
// 	pgs, err := ioc.Repo.ProductGroups.ListExhibitionPg(offset, limit)
// 	if err != nil {
// 		log.Println("pg list error", err)
// 	}

// 	log.Println("Total Exhibition PGs: ", len(pgs))
// 	for idx, pgDao := range pgs {
// 		log.Println("IDX", idx)
// 		if pgDao.GroupType == domain.PRODUCT_GROUP_EXHIBITION {
// 			ex, err := exhibition.FindExhibitionInProductGroup(pgDao.ID.Hex())
// 			if err != nil {
// 				log.Println("fail in find exhibition in pg", err)
// 				continue
// 			}
// 			if ex == nil {
// 				log.Println("pg have no ex", pgDao.ID)
// 				continue
// 			}
// 			_, err = exhibition.UpdateExhibition(ex)
// 			if err != nil {
// 				log.Println("fail in update exhibition", ex.ID.Hex())
// 			}
// 		}
// 	}
// }

// func HomeTabSyncer() {
// 	items, cnt, err := ioc.Repo.HomeTabItems.List(0, 20, true)
// 	if err != nil {
// 		log.Println("listing hometab item error", err)
// 	}
// 	log.Println("total cnt", cnt)

// 	for idx, item := range items {
// 		log.Println("IDX", idx)
// 		// for idx, exhibition := range item.Exhibitions {
// 		// 	newExhibition, err := ioc.Repo.Exhibitions.Get(exhibition.ID.Hex())
// 		// 	if err != nil {
// 		// 		log.Println("find ex error", err)
// 		// 	}
// 		// 	item.Exhibitions[idx] = newExhibition
// 		// }

// 		for idx, pd := range item.Products {
// 			newPd, err := ioc.Repo.Products.Get(pd.ID.Hex())
// 			if err != nil {
// 				log.Println("find pd error", err)
// 			}
// 			item.Products[idx] = newPd
// 		}

// 		_, err = ioc.Repo.HomeTabItems.Update(item)
// 		if err != nil {
// 			log.Println("HOIT", err)
// 		}
// 	}
// }

// func AddProductInSize() {
// 	filter := bson.M{"productinfo.brand.keyname": "INTREND"}
// 	pds, cnt, err := ioc.Repo.Products.List(0, 7000, filter, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	brand, err := ioc.Repo.Brands.GetByKeyname("INTREND")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	log.Println("total: ", cnt)

// 	for idx, pd := range pds {
// 		if idx%100 == 0 {
// 			log.Println("IDX", idx)
// 		}
// 		pd.ProductInfo.Brand = brand
// 		_, err := ioc.Repo.Products.Upsert(pd)
// 		if err != nil {
// 			log.Println("err found in products upsert", err)
// 		}
// 	}

// }
