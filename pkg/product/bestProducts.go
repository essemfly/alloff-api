package product

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MakeSnapshot() {
	log.Println("Making snapshot Start!!")
	alloffCatIDs := []string{""}
	alloffLev1Cats, _ := ioc.Repo.AlloffCategories.List(&alloffCatIDs[0])

	products := GetBestProductsFromAll()
	snapshot := domain.BestProductDAO{
		ID:               primitive.NewObjectID(),
		AlloffCategoryID: "",
		Products:         products,
	}
	_, err := ioc.Repo.BestProducts.Insert(&snapshot)
	if err != nil {
		log.Println("err occured in make snapshot", err)
	}

	for _, cat := range alloffLev1Cats {
		products := GetAlloffCategoryProducts(cat.ID.Hex())
		snapshot := domain.BestProductDAO{
			ID:               primitive.NewObjectID(),
			AlloffCategoryID: cat.ID.Hex(),
			Products:         products,
		}
		_, err := ioc.Repo.BestProducts.Insert(&snapshot)
		if err != nil {
			log.Println("err occured in make snapshot", err)
		}
	}
}

func GetAlloffCategoryProducts(alloffCatID string) []*domain.ProductDAO {
	pds, _, err := AlloffCategoryProductsListing(0, 100, nil, alloffCatID, "", []string{"70", "50"})
	if err != nil {
		log.Println("err occured in alloff cats product recording")
	}

	return pds
}

func GetBestProductsFromAll() []*domain.ProductDAO {
	productDaos, _, err := ProductsListing(0, 100, "", "", "", []string{"70", "100"})
	if err != nil {
		log.Println("err occured in products listing")
	}

	return productDaos
}
