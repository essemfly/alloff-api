package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapSizeMappingPolicy(dao *domain.SizeMappingPolicyDAO) *model.SizeMappingPolicy {
	return &model.SizeMappingPolicy{
		ID:                dao.ID.Hex(),
		AlloffSize:        MapAlloffSizeDaoToAlloffSize(dao.AlloffSize),
		AlloffCategory:    MapAlloffCatDaoToAlloffCat(dao.AlloffCategory),
		Sizes:             dao.Sizes,
		AlloffProductType: MapProductTypes(dao.AlloffProductType),
	}
}
