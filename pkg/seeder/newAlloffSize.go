package seeder

import (
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"log"
)

// TODO alloffCat, alloffSizeName, ProductType 으로 Document 특정해서 Upsert 하려고했는데, 잘안되서 일단 새로운 alloffSize에 대해서 inser만 하도록 해놓음 -> 구조체 플래그에 omitempty 앞에 공백
func NewAlloffSizes() {
	// *** 여성 :: 아우터,상의,원피스/세트,스커트 ****
	catOuter, err := ioc.Repo.AlloffCategories.GetByKeyname("1_outer")
	if err != nil {
		log.Panic("can not get cat outer")
	}
	catTop, err := ioc.Repo.AlloffCategories.GetByKeyname("1_top")
	if err != nil {
		log.Panic("can not get cat top")
	}
	catSet, err := ioc.Repo.AlloffCategories.GetByKeyname("1_onePiece")
	if err != nil {
		log.Panic("can not get cat onePiece")
	}
	catSkirt, err := ioc.Repo.AlloffCategories.GetByKeyname("1_skirt")
	if err != nil {
		log.Panic("can not get cat skirt")
	}
	catAccessory, err := ioc.Repo.AlloffCategories.GetByKeyname("1_accessory")
	if err != nil {
		log.Panic("can not get cat accessory")
	}
	cats1 := []*domain.AlloffCategoryDAO{catOuter, catTop, catSet, catSkirt}
	sizes1 := []string{"44", "55", "66", "77", "88", "99(이상)"}

	for _, cat := range cats1 {
		for _, size := range sizes1 {
			alloffSizeDao := &domain.AlloffSizeDAO{
				AlloffCategory: cat,
				AlloffSizeName: size,
				ProductType:    []domain.AlloffProductType{domain.Female},
			}
			upserted, err := ioc.Repo.AlloffSizes.Upsert(alloffSizeDao)
			if err != nil {
				log.Panic(cat, size, alloffSizeDao.ProductType)
			}
			log.Println("inserted : ", upserted)
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
				AlloffCategory: cat,
				AlloffSizeName: size,
				ProductType:    []domain.AlloffProductType{domain.Female},
			}
			upserted, err := ioc.Repo.AlloffSizes.Upsert(alloffSizeDao)
			if err != nil {
				log.Panic(cat, size, alloffSizeDao.ProductType)
			}
			log.Println("inserted : ", upserted)
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
				AlloffCategory: cat,
				AlloffSizeName: size,
				ProductType:    []domain.AlloffProductType{domain.Female},
			}
			upserted, err := ioc.Repo.AlloffSizes.Upsert(alloffSizeDao)
			if err != nil {
				log.Panic(cat, size, alloffSizeDao.ProductType)
			}
			log.Println("inserted : ", upserted)
		}
	}

	// **** 남성 :: 아우터,상의  ****
	cats4 := []*domain.AlloffCategoryDAO{catOuter, catTop}
	sizes4 := []string{"90", "95", "100", "105", "110", "115", "120(이상)"}

	for _, cat := range cats4 {
		for _, size := range sizes4 {
			alloffSizeDao := &domain.AlloffSizeDAO{
				AlloffCategory: cat,
				AlloffSizeName: size,
				ProductType:    []domain.AlloffProductType{domain.Male},
			}
			upserted, err := ioc.Repo.AlloffSizes.Upsert(alloffSizeDao)
			if err != nil {
				log.Panic(cat, size, alloffSizeDao.ProductType)
			}
			log.Println("inserted : ", upserted)
		}
	}

	// **** 남성 :: 바지  ****
	cats5 := []*domain.AlloffCategoryDAO{catPants}
	sizes5 := []string{"28", "30", "32", "34", "36", "38(이상)"}

	for _, cat := range cats5 {
		for _, size := range sizes5 {
			alloffSizeDao := &domain.AlloffSizeDAO{
				AlloffCategory: cat,
				AlloffSizeName: size,
				ProductType:    []domain.AlloffProductType{domain.Male},
			}
			upserted, err := ioc.Repo.AlloffSizes.Upsert(alloffSizeDao)
			if err != nil {
				log.Panic(cat, size, alloffSizeDao.ProductType)
			}
			log.Println("inserted : ", upserted)
		}
	}

	// **** 남성 :: 신발  ****
	cats6 := []*domain.AlloffCategoryDAO{catShoes}
	sizes6 := []string{"250", "255", "260", "265", "270", "275", "280(이상)"}

	for _, cat := range cats6 {
		for _, size := range sizes6 {
			alloffSizeDao := &domain.AlloffSizeDAO{
				AlloffCategory: cat,
				AlloffSizeName: size,
				ProductType:    []domain.AlloffProductType{domain.Male},
			}
			upserted, err := ioc.Repo.AlloffSizes.Upsert(alloffSizeDao)
			if err != nil {
				log.Panic(cat, size, alloffSizeDao.ProductType)
			}
			log.Println("inserted : ", upserted)
		}
	}

	// **** 키즈 :: 신발  ****
	cats7 := []*domain.AlloffCategoryDAO{catShoes}
	sizes7 := []string{"145", "150", "155", "160", "165", "170", "175", "180", "185", "190", "195", "200", "205", "210(이상)"}

	for _, cat := range cats7 {
		for _, size := range sizes7 {
			alloffSizeDao := &domain.AlloffSizeDAO{
				AlloffCategory: cat,
				AlloffSizeName: size,
				ProductType:    []domain.AlloffProductType{domain.Male},
			}
			upserted, err := ioc.Repo.AlloffSizes.Upsert(alloffSizeDao)
			if err != nil {
				log.Panic(cat, size, alloffSizeDao.ProductType)
			}
			log.Println("inserted : ", upserted)
		}
	}

	// **** 프리사이즈 ****
	cats8 := []*domain.AlloffCategoryDAO{catOuter, catTop, catSet, catSkirt, catAccessory}
	sizes8 := []string{"FREE"}

	for _, cat := range cats8 {
		for _, size := range sizes8 {
			alloffSizeDao := &domain.AlloffSizeDAO{
				AlloffCategory: cat,
				AlloffSizeName: size,
				ProductType:    []domain.AlloffProductType{domain.Male, domain.Female, domain.Kids},
			}
			upserted, err := ioc.Repo.AlloffSizes.Upsert(alloffSizeDao)
			if err != nil {
				log.Panic(cat, size, alloffSizeDao.ProductType)
			}
			log.Println("upserted : ", upserted)
		}
	}
}
