package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapCategoryDaoToCategory(catDao *domain.CategoryDAO) *model.Category {
	return &model.Category{
		ID:      catDao.ID.Hex(),
		Name:    catDao.Name,
		KeyName: catDao.KeyName,
	}
}

func MapAlloffCatDaoToAlloffCat(catDao *domain.AlloffCategoryDAO) *model.AlloffCategory {
	if catDao == nil {
		return nil
	}
	if catDao.CategoryType == "NORMAL" {
		newItem := model.AlloffCategory{
			ID:       catDao.ID.Hex(),
			Name:     catDao.Name,
			KeyName:  catDao.KeyName,
			Level:    catDao.Level,
			ParentID: catDao.ParentId.Hex(),
			ImgURL:   catDao.ImgURL,
		}
		return &newItem
	}
	return nil
}
