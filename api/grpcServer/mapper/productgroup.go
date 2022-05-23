package mapper

import (
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
)

func ProductGroupMapper(pg *domain.ProductGroupDAO) *grpcServer.ProductGroupMessage {
	products := []*grpcServer.ProductInGroupMessage{}
	for _, pd := range pg.Products {
		if pd.Product != nil {
			// 이 부분 productgroup에서 어떤 products들을 보내줘야할지에 대해서 고민이 필요하다.
			products = append(products, &grpcServer.ProductInGroupMessage{
				Priority: int32(pd.Priority),
				Product:  ProductInfoMapper(pd.Product.ProductInfo),
			})
		} else {
			pdInfoDao, err := ioc.Repo.ProductMetaInfos.Get(pd.ProductID.Hex())
			if err != nil {
				continue
			}
			products = append(products, &grpcServer.ProductInGroupMessage{
				Priority: int32(pd.Priority),
				Product:  ProductInfoMapper(pdInfoDao),
			})
		}
	}

	brand := &grpcServer.BrandMessage{}
	if pg.Brand != nil {
		brand = BrandMapper(pg.Brand)
	}

	return &grpcServer.ProductGroupMessage{
		Title:          pg.Title,
		ShortTitle:     pg.ShortTitle,
		Instruction:    pg.Instruction,
		ImageUrl:       pg.ImgUrl,
		Products:       products,
		StartTime:      pg.StartTime.String(),
		FinishTime:     pg.FinishTime.String(),
		ProductGroupId: pg.ID.Hex(),
		GroupType:      GroupTypeMapper(pg.GroupType),
		Brand:          brand,
	}
}

func GroupTypeMapper(groupType domain.ProductGroupType) grpcServer.ProductGroupType {
	switch groupType {
	case domain.PRODUCT_GROUP_TIMEDEAL:
		return grpcServer.ProductGroupType_PRODUCT_GROUP_TIMEDEAL
	case domain.PRODUCT_GROUP_EXHIBITION:
		return grpcServer.ProductGroupType_PRODUCT_GROUP_EXHIBITION
	case domain.PRODUCT_GROUP_GROUPDEAL:
		return grpcServer.ProductGroupType_PRODUCT_GROUP_GROUPDEAL
	case domain.PRODUCT_GROUP_BRAND_TIMEDEAL:
		return grpcServer.ProductGroupType_PRODUCT_GROUP_BRAND_TIMEDEAL
	}
	return grpcServer.ProductGroupType_PRODUCT_GROUP_EXHIBITION
}
