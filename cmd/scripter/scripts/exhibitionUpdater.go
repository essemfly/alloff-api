package scripts

import (
	"log"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.uber.org/zap"
)

func ExhibitionStatusUpdater() {
	offset, limit := 0, 100 // Set 100 for live exhibitions
	onlyLive := true
	exDaos, cnt, err := ioc.Repo.Exhibitions.List(offset, limit, &onlyLive, domain.EXHIBITION_STATUS_ALL, domain.EXHIBITION_TIMEDEAL, "")
	log.Println(len(exDaos))
	if err != nil {
		config.Logger.Error("err on listing exhibitions", zap.Error(err))
	}
	log.Println("total exhibitions", cnt)
	for _, exDao := range exDaos {
		//_, err = exhibition.ExhibitionSyncer(exDao)
		if err != nil {
			config.Logger.Error("exhibitino sync error "+exDao.ID.Hex(), zap.Error(err))
		}
	}
}
