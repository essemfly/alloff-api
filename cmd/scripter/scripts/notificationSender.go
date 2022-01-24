package scripts

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/pkg/notification"
)

func SendNotification() {
	offset, limit := 0, 500
	notis, err := ioc.Repo.Notifications.List(offset, limit, true)
	if err != nil {
		log.Println(err)
	}
	log.Println("# notis: ", len(notis))

	for _, noti := range notis {
		filteredDeviceIDs := []string{}

		for _, deviceID := range noti.DeviceIDs {
			deviceDao, _ := ioc.Repo.Devices.GetByDeviceID(deviceID)
			if deviceDao.AllowNotification {
				filteredDeviceIDs = append(filteredDeviceIDs, deviceID)
			}
		}
		noti.DeviceIDs = filteredDeviceIDs
		err := notification.SendNotification(noti)
		if err != nil {
			log.Println("Error on sending notification", err)
		}
	}

}
