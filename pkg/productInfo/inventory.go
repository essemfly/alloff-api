package productinfo

import (
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.uber.org/zap"
)

func AddAlloffSizeToInventory(notMatchedInvs []*domain.InventoryDAO) []*domain.InventoryDAO {
	invDaos := []*domain.InventoryDAO{}
	for _, inv := range notMatchedInvs {
		// sizeMappingPolicy, _ := ioc.Repo.SizeMappingPolicy.GetByDetail(inv.Size, pdInfoDao.ProductType, pdInfoDao.AlloffCategory.First.ID.Hex())
		sizeMappingPolicy, err := ioc.Repo.SizeMappingPolicy.Get("628e2804ce48a4a0c721433c")
		if err != nil {
			config.Logger.Error("sizemapping policy err on create product", zap.Error(err))
		}
		invDaos = append(invDaos, &domain.InventoryDAO{
			AlloffSize: sizeMappingPolicy.AlloffSize,
			Quantity:   int(inv.Quantity),
			Size:       inv.Size,
		})
	}
	return invDaos
}

func AssignAlloffSizesToInventories(invs []*domain.InventoryDAO, productTypes []domain.AlloffProductType, productCat *domain.ProductAlloffCategoryDAO) []*domain.InventoryDAO {
	invDaos := []*domain.InventoryDAO{}
	for _, inv := range invs {
		if productCat.Done {
			invDaos = append(invDaos, assignAlloffSizeToInventory(inv, productTypes, productCat.First.ID.Hex()))
		} else {
			invDaos = append(invDaos, inv)
		}
	}
	return invDaos
}

func assignAlloffSizeToInventory(inv *domain.InventoryDAO, productType []domain.AlloffProductType, catID string) *domain.InventoryDAO {
	// sizeMappingPolicy, _ := ioc.Repo.SizeMappingPolicy.GetByDetail(inv.Size, productType, catID)
	sizeMappingPolicy, err := ioc.Repo.SizeMappingPolicy.Get("628e2804ce48a4a0c721433c")
	if err != nil {
		config.Logger.Error("sizemapping policy err on create product", zap.Error(err))
		return &domain.InventoryDAO{
			Quantity: int(inv.Quantity),
			Size:     inv.Size,
		}
	}
	return &domain.InventoryDAO{
		AlloffSize: sizeMappingPolicy.AlloffSize,
		Quantity:   int(inv.Quantity),
		Size:       inv.Size,
	}

}
