package productinfo

import (
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.uber.org/zap"
)

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
	//sizeMappingPolicy, err := ioc.Repo.SizeMappingPolicy.GetByDetail(inv.Size, productType, catID)
	sizeMappingPolicies, err := ioc.Repo.SizeMappingPolicy.ListByDetail(inv.Size, productType, catID)
	if err != nil {
		config.Logger.Error("sizemapping policy err on create product", zap.Error(err))
		return &domain.InventoryDAO{
			Quantity: int(inv.Quantity),
			Size:     inv.Size,
		}
	}

	alloffSizes := []*domain.AlloffSizeDAO{}
	for _, policy := range sizeMappingPolicies {
		alloffSizes = append(alloffSizes, policy.AlloffSize)
	}

	return &domain.InventoryDAO{
		AlloffSizes: alloffSizes,
		Quantity:    int(inv.Quantity),
		Size:        inv.Size,
	}

}
