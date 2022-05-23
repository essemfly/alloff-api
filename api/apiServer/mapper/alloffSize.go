package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapAlloffSizeDaoToAlloffSize(alloffSizeDao *domain.AlloffSizeDAO) *model.AlloffSize {
	return &model.AlloffSize{
		ID:       alloffSizeDao.ID.Hex(),
		SizeName: alloffSizeDao.SizeName,
	}
}
