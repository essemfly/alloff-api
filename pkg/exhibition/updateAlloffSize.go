package exhibition

// func UpdateAlloffInventory(exhibitionDao *domain.ExhibitionDAO) {
// 	filter := product.ProductListInput{
// 		Offset:       0,
// 		Limit:        10000,
// 		ExhibitionID: exhibitionDao.ID.Hex(),
// 	}
// 	pds, _, err := product.ListProducts(filter)
// 	if err != nil {
// 		log.Println("err occurred on get product list : ", err)
// 	}

// 	exhibitionInventories := []*domain.AlloffInventoryDAO{}

// 	for _, pd := range pds {
// 		if len(exhibitionInventories) == 0 {
// 			// 상품군의 첫 인벤토리 > 무조건 추가
// 			exhibitionInventories = append(exhibitionInventories, pd.ProductInfo.AlloffInventory...)
// 		} else {
// 			// 상품군의 두번째 인벤토리부터는 기존 인벤토리와 비교해서 수량만 올릴지, 인벤토리 자체를 추가할지 결정한다.
// 			for _, pdInv := range pd.ProductInfo.AlloffInventory {
// 				contains := false

// 				for idx, exhibitionInv := range exhibitionInventories {
// 					if exhibitionInv.AlloffSize.ID == pdInv.AlloffSize.ID {
// 						// 상품의 사이즈가 exhibition에 이미 있는 경우 > 수량만 추가함
// 						exhibitionInventories[idx].Quantity += pdInv.Quantity
// 						contains = true
// 					}
// 				}
// 				// 상품의 사이즈가 exhibition에 포함되지 않은 경우 > 새로 추가함
// 				if !contains {
// 					exhibitionInventories = append(exhibitionInventories, pdInv)
// 				}
// 			}
// 		}
// 	}

// 	exhibitionDao.MetaInfos.AlloffInventories = exhibitionInventories
// 	_, err = ioc.Repo.Exhibitions.Upsert(exhibitionDao)
// 	if err != nil {
// 		log.Println("err occurred on upsert exhibition : ", err)
// 	}
// }
