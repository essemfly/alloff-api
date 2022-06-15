package notification

import (
	"errors"
	"math"
	"time"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/push"
	"go.uber.org/zap"
)

const CHUNK_SIZE = 500
const AWAIT_TIME_FOR_CHUNK = 2

func Send(noti *domain.NotificationDAO) error {
	if noti.Status != domain.NOTIFICATION_READY && noti.Status != domain.NOTIFICATION_FAILED {
		return errors.New("notification status not ready")
	}

	noti.Status = domain.NOTIFICATION_SUCCEEDED
	noti.Sended = time.Now()
	noti.Updated = time.Now()

	deviceIDs := loadAppropriateDevices(noti)
	deviceChunks := chunkStringSlice(deviceIDs, CHUNK_SIZE)
	numTotalSend, numTotalFailed := 0, 0
	for _, chunk := range deviceChunks {
		noti.DeviceIDs = chunk
		notiResult, err := push.SendNotification(noti)
		if err != nil || notiResult.Success != "ok" {
			noti.Status = domain.NOTIFICATION_FAILED
			ioc.Repo.Notifications.Update(noti)
			return err
		}
		numTotalFailed += len(notiResult.Logs)
		numTotalSend += len(chunk) - len(notiResult.Logs)
	}

	noti.NumUsersFailed = numTotalFailed
	noti.NumUsersPushed = numTotalSend

	_, err := ioc.Repo.Notifications.Update(noti)
	if err != nil {
		config.Logger.Error("notifucation update fail", zap.Error(err))
		return err
	}

	return nil
}

func loadAppropriateDevices(noti *domain.NotificationDAO) (deviceIDs []string) {
	return
}

func chunkStringSlice(s []string, chunkSize int) [][]string {
	chunkNum := int(math.Ceil(float64(len(s)) / float64(chunkSize)))
	res := make([][]string, 0, chunkNum)
	for i := 0; i < chunkNum-1; i++ {
		res = append(res, s[i*chunkSize:(i+1)*chunkSize])
	}
	res = append(res, s[(chunkNum-1)*chunkSize:])
	return res
}
