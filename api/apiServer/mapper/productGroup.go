package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"sort"
)

func MapProductGroup(pgDao *domain.ProductGroupDAO) *model.ProductGroup {
	pds := []*model.Product{}
	soldouts := []*model.Product{}

	sort.Slice(pgDao.Products, func(i, j int) bool {
		return pgDao.Products[i].Priority < pgDao.Products[j].Priority
	})

	for _, productInPg := range pgDao.Products {

		pd := MapProduct(productInPg.Product)
		isSoldOut := true
		for _, inv := range productInPg.Product.Inventory {
			if inv.Quantity > 0 {
				isSoldOut = false
			}
		}
		if isSoldOut || pd.IsSoldout {
			soldouts = append(soldouts, pd)
		} else {
			pds = append(pds, pd)
		}
	}

	pds = append(pds, soldouts...)

	brand := &model.Brand{}
	if pgDao.Brand != nil {
		brand = MapBrandDaoToBrand(pgDao.Brand, false)
	}

	pg := &model.ProductGroup{
		ID:            pgDao.ID.Hex(),
		Brand:         brand,
		Title:         pgDao.Title,
		ShortTitle:    pgDao.ShortTitle,
		ImgURL:        pgDao.ImgUrl,
		Products:      pds,
		TotalProducts: len(pds),
	}
	return pg
}
