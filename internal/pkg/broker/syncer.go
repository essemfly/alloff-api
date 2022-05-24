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
// Product의 ExhibitionID를 여기서 syncing 해주는게 맞을까?
func ProductGroupSyncer(pgDao *domain.ProductGroupDAO) (*domain.ProductGroupDAO, error) {
	for _, pd := range pgDao.Products {
		newPd, err := ioc.Repo.Products.Get(pd.Product.ID.Hex())
		if err != nil {
			log.Println("err not found product", pd.Product.ID.Hex())
		}

		if newPd.ExhibitionID != pgDao.ExhibitionID {
			newPd.ExhibitionID = pgDao.ExhibitionID
			newPd, _ = ioc.Repo.Products.Upsert(newPd)
		}

		if newPd.ProductGroupID != pgDao.ID.Hex() {
			newPd.ProductGroupID = pgDao.ID.Hex()
			newPd.ExhibitionID = pgDao.ExhibitionID
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
