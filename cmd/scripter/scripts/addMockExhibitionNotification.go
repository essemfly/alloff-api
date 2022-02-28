package scripts

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddExhibitionNotification() error {

	isTest := true

	pushtitle := "Exhibitino adder test"
	message := "(광고) 산드로 인기상품 아울렛 가격으로 !"
	exID := "6214e01ae8c186b66cd41538"

	ex, err := ioc.Repo.Exhibitions.Get(exID)
	if err != nil {
		log.Println("err", err)
		return err
	}

	if isTest {
		mobiles := []string{"01097711882", "01034671612"}
		deviceIds := []string{}
		for _, mobile := range mobiles {
			user, _ := ioc.Repo.Users.GetByMobile(mobile)
			devices, _ := ioc.Repo.Devices.ListAllowedByUser(user.ID.Hex())
			for _, device := range devices {
				deviceIds = append(deviceIds, device.DeviceId)
			}
		}

		timedealNoti := domain.NotificationDAO{
			ID:               primitive.NewObjectID(),
			Status:           domain.NOTIFICATION_READY,
			NotificationType: domain.NOTIFICATION_EXHIBITION_OPEN_NOTIFICATION,
			Title:            pushtitle,
			Message:          message,
			Notificationid:   "/exhibition/" + ex.ID.Hex(),
			DeviceIDs:        deviceIds,
			NavigateTo:       "/exhibition",
			ReferenceID:      "/" + ex.ID.Hex(),
		}

		_, err := ioc.Repo.Notifications.Insert(&timedealNoti)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}

	devices, err := ioc.Repo.Devices.ListAllowed()
	if err != nil {
		return err
	}

	deviceIDs := []string{}

	for _, device := range devices {
		deviceIDs = append(deviceIDs, device.DeviceId)
		if len(deviceIDs) >= 300 {
			timedealNoti := domain.NotificationDAO{
				ID:               primitive.NewObjectID(),
				Status:           domain.NOTIFICATION_READY,
				NotificationType: domain.NOTIFICATION_EXHIBITION_OPEN_NOTIFICATION,
				Title:            pushtitle,
				Message:          message,
				Notificationid:   "/exhibition/" + ex.ID.Hex(),
				DeviceIDs:        deviceIDs,
				NavigateTo:       "/exhibition",
				ReferenceID:      "/" + ex.ID.Hex(),
			}

			_, err := ioc.Repo.Notifications.Insert(&timedealNoti)
			if err != nil {
				log.Println(err)
				return err
			}
			deviceIDs = []string{}
		}
	}

	timedealNoti := domain.NotificationDAO{
		ID:               primitive.NewObjectID(),
		Status:           domain.NOTIFICATION_READY,
		NotificationType: domain.NOTIFICATION_EXHIBITION_OPEN_NOTIFICATION,
		Title:            pushtitle,
		Message:          message,
		Notificationid:   "/exhibition/" + ex.ID.Hex(),
		DeviceIDs:        deviceIDs,
		NavigateTo:       "/exhibition",
		ReferenceID:      "/" + ex.ID.Hex(),
	}

	_, err = ioc.Repo.Notifications.Insert(&timedealNoti)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
