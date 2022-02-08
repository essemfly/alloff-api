package scripts

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTimedealNotification() error {

	isTest := true

	pushtitle := "Sandro 오픈🎉"
	message := "(광고) 산드로 인기상품 아울렛 가격으로 !"
	timedealID := "6176f430ee2c39d09bebeac8"

	timedeal, err := ioc.Repo.ProductGroups.Get(timedealID)
	if err != nil {
		log.Println("err", err)
		return err
	}

	if isTest {
		mobiles := []string{"01097711882"}
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
			NotificationType: domain.NOTIFICATION_TIMEDEAL_OPEN_NOTIFICATION,
			Title:            pushtitle,
			Message:          message,
			Notificationid:   "/timedeals/" + timedeal.ID.Hex(),
			DeviceIDs:        deviceIds,
			NavigateTo:       "/timedeals",
			ReferenceID:      "/" + timedeal.ID.Hex(),
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
				NotificationType: domain.NOTIFICATION_TIMEDEAL_OPEN_NOTIFICATION,
				Title:            pushtitle,
				Message:          message,
				Notificationid:   "/timedeals/" + timedeal.ID.Hex(),
				DeviceIDs:        deviceIDs,
				NavigateTo:       "/timedeals",
				ReferenceID:      "/" + timedeal.ID.Hex(),
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
		NotificationType: domain.NOTIFICATION_TIMEDEAL_OPEN_NOTIFICATION,
		Title:            pushtitle,
		Message:          message,
		Notificationid:   "/timedeals/" + timedeal.ID.Hex(),
		DeviceIDs:        deviceIDs,
		NavigateTo:       "/timedeals",
		ReferenceID:      "/" + timedeal.ID.Hex(),
	}

	_, err = ioc.Repo.Notifications.Insert(&timedealNoti)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
