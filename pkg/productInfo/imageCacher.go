package productinfo

import (
	"log"

	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
)

func CacheProductsImage(pd *domain.ProductMetaInfoDAO) error {
	url := "https://s70q2owcf2.execute-api.ap-northeast-2.amazonaws.com/?productId=" + pd.ID.Hex()
	_, err := utils.RequestRetryer(url, utils.REQUEST_GET, utils.GetGeneralHeader(), "", "")
	if err != nil {
		log.Println("ERR", err)
		return err
	}
	return nil
}
