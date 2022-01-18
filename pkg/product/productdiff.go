package product

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func InsertProductDiff(pd *domain.ProductDAO, oldPrice int) error {
	err := ioc.Repo.ProductDiffs.Insert(&domain.ProductDiffDAO{
		OldPrice:   oldPrice,
		NewProduct: pd,
		Type:       "price",
		IsPushed:   false,
	})
	if err != nil {
		log.Println("error occured in product dif update")
		return err
	}
	return nil
}
