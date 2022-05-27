package seeder

import (
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func NewAlloffSizes() {
	// *** 여성 :: 아우터,상의,원피스/세트,스커트 ****
	catOuter, err := ioc.Repo.AlloffCategories.GetByKeyname("1_outer")
	if err != nil {
		log.Panic("can not get cat outer for female")
	}
	catTop, err := ioc.Repo.AlloffCategories.GetByKeyname("1_top")
	if err != nil {
		log.Panic("can not get cat top for female")
	}
	catSet, err := ioc.Repo.AlloffCategories.GetByKeyname("1_onePiece")
	if err != nil {
		log.Panic("can not get cat onePiece for female")
	}
	catSkirt, err := ioc.Repo.AlloffCategories.GetByKeyname("1_skirt")
	if err != nil {
		log.Panic("can not get cat skirt for female")
	}
	cats1 := []*domain.AlloffCategoryDAO{catOuter, catTop, catSet, catSkirt}
	sizes1 := []string{"44", "55", "66", "77", "88", "99(이상)"}

	for _, cat := range cats1 {
		for _, size := range sizes1 {
			alloffSizeDao := &domain.AlloffSizeDAO{
				ID:             primitive.NewObjectID(),
				AlloffCategory: cat,
				AlloffSizeName: size,
				ProductType:    []domain.AlloffProductType{domain.Female},
			}
			inserted, err := ioc.Repo.AlloffSizes.Upsert(alloffSizeDao)
			if err != nil {
				log.Panic(cat, size, alloffSizeDao.ProductType)
			}
			log.Println("inserted : ", inserted)
		}
	}

	// *** 여성 :: 바지 ****
	catPants, err := ioc.Repo.AlloffCategories.GetByKeyname("1_bottom")
	if err != nil {
		log.Panic("can not get cat bottom for female")
	}
	cats2 := []*domain.AlloffCategoryDAO{catPants}
	sizes2 := []string{"24-25", "26-27", "28-29", "30-31", "32-33", "34(이상)"}

	for _, cat := range cats2 {
		for _, size := range sizes2 {
			alloffSizeDao := &domain.AlloffSizeDAO{
				ID:             primitive.NewObjectID(),
				AlloffCategory: cat,
				AlloffSizeName: size,
				ProductType:    []domain.AlloffProductType{domain.Female},
			}
			inserted, err := ioc.Repo.AlloffSizes.Upsert(alloffSizeDao)
			if err != nil {
				log.Panic(cat, size, alloffSizeDao.ProductType)
			}
			log.Println("inserted : ", inserted)
		}
	}

	// **** 여성 :: 신발 ****
	catShoes, err := ioc.Repo.AlloffCategories.GetByKeyname("1_shoes")
	if err != nil {
		log.Panic("1")
	}
	cats3 := []*domain.AlloffCategoryDAO{catShoes}
	sizes3 := []string{"220", "225", "230", "235", "240", "245", "250", "255", "260", "265(이상)"}

	for _, cat := range cats3 {
		for _, size := range sizes3 {
			alloffSizeDao := &domain.AlloffSizeDAO{
				ID:             primitive.NewObjectID(),
				AlloffCategory: cat,
				AlloffSizeName: size,
				ProductType:    []domain.AlloffProductType{domain.Female},
			}
			inserted, err := ioc.Repo.AlloffSizes.Upsert(alloffSizeDao)
			if err != nil {
				log.Panic(cat, size, alloffSizeDao.ProductType)
			}
			log.Println("inserted : ", inserted)
		}
	}

	// **** 남성 :: 아우터,상의  ****
	cats4 := []*domain.AlloffCategoryDAO{catOuter, catTop}
	sizes4 := []string{"90", "95", "100", "105", "110", "115", "120(이상)"}

	for _, cat := range cats4 {
		for _, size := range sizes4 {
			alloffSizeDao := &domain.AlloffSizeDAO{
				ID:             primitive.NewObjectID(),
				AlloffCategory: cat,
				AlloffSizeName: size,
				ProductType:    []domain.AlloffProductType{domain.Male},
			}
			inserted, err := ioc.Repo.AlloffSizes.Upsert(alloffSizeDao)
			if err != nil {
				log.Panic(cat, size, alloffSizeDao.ProductType)
			}
			log.Println("inserted : ", inserted)
		}
	}

	// **** 남성 :: 바지  ****
	cats5 := []*domain.AlloffCategoryDAO{catPants}
	sizes5 := []string{"28", "30", "32", "34", "36", "38(이상)"}

	for _, cat := range cats5 {
		for _, size := range sizes5 {
			alloffSizeDao := &domain.AlloffSizeDAO{
				ID:             primitive.NewObjectID(),
				AlloffCategory: cat,
				AlloffSizeName: size,
				ProductType:    []domain.AlloffProductType{domain.Male},
			}
			inserted, err := ioc.Repo.AlloffSizes.Upsert(alloffSizeDao)
			if err != nil {
				log.Panic(cat, size, alloffSizeDao.ProductType)
			}
			log.Println("inserted : ", inserted)
		}
	}

	// **** 남성 :: 신발  ****
	cats6 := []*domain.AlloffCategoryDAO{catShoes}
	sizes6 := []string{"250", "255", "260", "265", "270", "275", "280(이상)"}

	for _, cat := range cats6 {
		for _, size := range sizes6 {
			alloffSizeDao := &domain.AlloffSizeDAO{
				ID:             primitive.NewObjectID(),
				AlloffCategory: cat,
				AlloffSizeName: size,
				ProductType:    []domain.AlloffProductType{domain.Male},
			}
			inserted, err := ioc.Repo.AlloffSizes.Upsert(alloffSizeDao)
			if err != nil {
				log.Panic(cat, size, alloffSizeDao.ProductType)
			}
			log.Println("inserted : ", inserted)
		}
	}

	// **** 키즈 :: 신발  ****
	cats7 := []*domain.AlloffCategoryDAO{catShoes}
	sizes7 := []string{"145", "150", "155", "160", "165", "170", "180", "185", "190", "195", "200", "205", "210(이상)"}

	for _, cat := range cats7 {
		for _, size := range sizes7 {
			alloffSizeDao := &domain.AlloffSizeDAO{
				ID:             primitive.NewObjectID(),
				AlloffCategory: cat,
				AlloffSizeName: size,
				ProductType:    []domain.AlloffProductType{domain.Male},
			}
			inserted, err := ioc.Repo.AlloffSizes.Upsert(alloffSizeDao)
			if err != nil {
				log.Panic(cat, size, alloffSizeDao.ProductType)
			}
			log.Println("inserted : ", inserted)
		}
	}
}
