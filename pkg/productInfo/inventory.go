package productinfo

import (
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.uber.org/zap"
)

func AssignProductsInventory(pdDao *domain.ProductMetaInfoDAO) {
	newInv := AssignAlloffSizesToInventories(pdDao.Inventory, pdDao.ProductType, pdDao.AlloffCategory)
	pdDao.Inventory = newInv

	_, err := Update(pdDao)
	if err != nil {
		config.Logger.Error("error occurred on upsert product-meta-info : ", zap.Error(err))
	}
}

func AssignAlloffSizesToInventories(invs []*domain.InventoryDAO, productTypes []domain.AlloffProductType, productCat *domain.ProductAlloffCategoryDAO) []*domain.InventoryDAO {
	invDaos := []*domain.InventoryDAO{}
	for _, inv := range invs {
		if productCat != nil && productCat.Done {
			invDaos = append(invDaos, assignAlloffSizeToInventory(inv, productTypes, productCat.First.ID.Hex()))
		} else {
			invDaos = append(invDaos, inv)
		}
	}
	return invDaos
}

func assignAlloffSizeToInventory(inv *domain.InventoryDAO, productType []domain.AlloffProductType, catID string) *domain.InventoryDAO {
	alloffSizes, err := ioc.Repo.AlloffSizes.ListByDetail(inv.Size, productType, catID)
	if err != nil {
		config.Logger.Error("alloffsize listing err on create product", zap.Error(err))
		return &domain.InventoryDAO{
			Quantity: int(inv.Quantity),
			Size:     inv.Size,
		}
	}

	return &domain.InventoryDAO{
		AlloffSizes: alloffSizes,
		Quantity:    int(inv.Quantity),
		Size:        inv.Size,
	}
}
