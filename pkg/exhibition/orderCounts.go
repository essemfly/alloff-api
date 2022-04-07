package exhibition

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func GetCurrentSales(exDao *domain.ExhibitionDAO) int {
	num, err := ioc.Repo.OrderCounts.Get(exDao.ID.Hex())
	if err != nil {
		log.Println("err occured on getting order counts", err)
		num = 9999 // Temp code for debug
	}
	return num
}
