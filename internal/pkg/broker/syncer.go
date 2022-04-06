package broker

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
)

// Pg를 받으면, 안에 있는 PD들 Productgroupid 업데이트하고, 새 것들을 PG에 넣어주는 역할
// PG안에 있는 PD들이 PGID가 잘못박혀서 결제 못되게 하는 것을 막는다.
// 다만 PG가 끝난 PD들에서 PGID들을 없애는 것은 불가능함
func ProductGroupSyncer(pgDao *domain.ProductGroupDAO) (*domain.ProductGroupDAO, error) {
	for _, pd := range pgDao.Products {
		newPd, err := ioc.Repo.Products.Get(pd.Product.ID.Hex())
		if err != nil {
			log.Println("err not found product", pd.Product.ID.Hex())
		}
		if newPd.ProductGroupId != pgDao.ID.Hex() {
			newPd.ProductGroupId = pgDao.ID.Hex()
			updatedPd, err := ioc.Repo.Products.Upsert(newPd)
			if err != nil {
				log.Println("err upsert product", newPd.ID.Hex())
			} else {
				pd.Product = updatedPd
			}

		} else {
			pd.Product = newPd
		}
	}

	newPg, err := ioc.Repo.ProductGroups.Upsert(pgDao)

	return newPg, err
}

// Exhibition안에 들어있는 섹션들에 대해서, 색션들의 PG상태를 업데이트 시켜줌
// Exhibition이 갖고 있는 PG와 PG들이 갖고 있는 EXID들 사이에서 맞춰줘야하는데
// Ex가 갖고 있는 ProductGROUPS가 우선임
func ExhibitionSyncer(exDao *domain.ExhibitionDAO) {
	newPgs := []*domain.ProductGroupDAO{}
	for idx, pg := range exDao.ProductGroups {
		log.Println("Exhibition idx", idx)
		pgDao, err := ioc.Repo.ProductGroups.Get(pg.ID.Hex())
		if err != nil {
			log.Println("Update exhibition not found pgID: "+pg.ID.Hex(), err)
			continue
		}

		newPgDao, err := ProductGroupSyncer(pgDao)
		if err != nil {
			log.Println("product group syncing failed:" + pgDao.ID.Hex())
		}

		newPgDao.StartTime = exDao.StartTime
		newPgDao.FinishTime = exDao.FinishTime
		newPgDao.ExhibitionID = exDao.ID.Hex()
		newPgDao.GroupType = domain.PRODUCT_GROUP_EXHIBITION
		updatedPgDao, err := ioc.Repo.ProductGroups.Upsert(newPgDao)
		if err != nil {
			log.Println("product group update failed", newPgDao.ID.Hex())
		}
		newPgs = append(newPgs, updatedPgDao)
	}

	exDao.ProductGroups = newPgs

	_, err := ioc.Repo.Exhibitions.Upsert(exDao)
	if err != nil {
		log.Println("failed in upsert exhibition", err)
	}
}

// PD가 업데이트되면, 홈탭에 있는 상품들 업데이트 하는 코드는 갖고있기가 어려움. 일단 패스

// HomeTab Syncer
func HomeTabSyncer() {
	items, cnt, err := ioc.Repo.HomeTabItems.List(0, 200, false)
	if err != nil {
		log.Println("listing hometab item error", err)
	}
	log.Println("total cnt", cnt)

	for idx, item := range items {
		log.Println("IDX", idx)
		for idx, exhibition := range item.Exhibitions {
			newExhibition, err := ioc.Repo.Exhibitions.Get(exhibition.ID.Hex())
			if err != nil {
				log.Println("find ex error", err)
			}
			item.Exhibitions[idx] = newExhibition
			item.Products = newExhibition.ListCheifProducts()
		}

		_, err = ioc.Repo.HomeTabItems.Update(item)
		if err != nil {
			log.Println("HOIT", err)
		}
	}
}

func BrandSyncer(brandKeyname string) {
	offset, limit := 0, 20000

	newBrand, _ := ioc.Repo.Brands.GetByKeyname(brandKeyname)

	filter := bson.M{
		"productinfo.brand.keyname": brandKeyname,
	}

	pds, _, err := ioc.Repo.Products.List(offset, limit, filter, nil)
	if err != nil {
		log.Println("err", err)
	}

	for idx, pd := range pds {
		if idx%100 == 0 {
			log.Println("IDX", idx)
		}
		pd.ProductInfo.Brand = newBrand
		_, err := ioc.Repo.Products.Upsert(pd)
		if err != nil {
			log.Println("err occured", err)
		}
	}

}
