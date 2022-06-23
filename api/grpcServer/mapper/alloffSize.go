package mapper

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
)

func AlloffSizeMapper(alloffSize *domain.AlloffSizeDAO) *grpcServer.AlloffSizeMessage {
	return &grpcServer.AlloffSizeMessage{
		AlloffSizeId:   alloffSize.ID.Hex(),
		AlloffSizeName: alloffSize.AlloffSizeName,
		AlloffCategory: AlloffCategoryMapper(alloffSize.AlloffCategory),
		Sizes:          alloffSize.Sizes,
		ProductTypes:   ProductTypeMapper(alloffSize.ProductType),
	}
}
