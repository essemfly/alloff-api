package mapper

import (
	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AlloffCategoryMapper(cat *domain.AlloffCategoryDAO) *grpcServer.AlloffCategoryMessage {
	return &grpcServer.AlloffCategoryMessage{
		CategoryId: cat.ID.Hex(),
		Name:       cat.Name,
		Keyname:    cat.KeyName,
		Level:      int32(cat.Level),
		ParentId:   cat.ParentId.Hex(),
		ImgUrl:     cat.ImgURL,
	}
}
