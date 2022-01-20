package mapper

import (
	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func ProductGroupMapper(pg *domain.ProductGroupDAO) *grpcServer.ProductGroupMessage {
	products := []*grpcServer.ProductInGroupMessage{}
	for _, pd := range pg.Products {
		products = append(products, &grpcServer.ProductInGroupMessage{
			Priority: int32(pd.Priority),
			Product:  ProductMapper(pd.Product),
		})
	}

	return &grpcServer.ProductGroupMessage{
		Title:       pg.Title,
		ShortTitle:  pg.ShortTitle,
		Instruction: pg.Instruction,
		ImageUrl:    pg.ImgUrl,
		Products:    products,
		StartTime:   pg.StartTime.String(),
		FinishTime:  pg.FinishTime.String(),
	}
}