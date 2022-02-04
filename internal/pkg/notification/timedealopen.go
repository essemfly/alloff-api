package notification

import (
	"log"
	"strconv"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTimedealNotification(timedeal *domain.ProductGroupDAO, pushtitle string) error {
	// Admins
	adminDeviceIDs := []string{
		"c0IfW-xzdUx3nu6XmpzXNz:APA91bGYmuaoWsdoNzhp_SqUiaKm_8bPK-MfUgmJpmvprUui6run4qKTEeF8QpPSnrR7f5oK2Suy5BJduO04C2DuKWmYHbYZJCRy_FtI6Rm6kAxEopwiRSGjymiqGzXVpz8i8nffFYV1", // 주영
		"dJ4rSV2L60aLpHxQjxoFz-:APA91bGJR2YXQkmBbv0rELM6caPUeZ3C1MnBkz1wlI68wCzDRhc9Bsma3stSCRXTGip-6mxdtj2GfuMT4c0XV85AWDuLr0lkH33VDNRsuc8nqo24JGOHmDDpYl_wetLh9vYL-3I0A6ID", // 석민
		"cBRNxXqZwEeapVVni8KJcG:APA91bG7UnRCfxWvNR7ngSYNfhTazApx9yAlQXtXqCJDWpn_X-cwVMnnUDLmjLUCso9s7_oiP_xrBkOqoa-1ie3LaRsckENluZTaxWcNAKpdUvVZtV9Pq_TRgRdmwtpA0kE_-Mx-_Nzl", // 명규
		"eEx9IUeGQtGgo__aVsJih4:APA91bG0NaGdVP_2NyfBWAXXXdZ0-iR9J-pRTf6JxjueMowaIg9HQDYiZ5PkGeqTypblQNDg-Fecdask8W6o15lNHOhzCHfHzIzofE63APEZPNQQbRVjaaBCTk6m3kAeIqoBuJl9rmHh", // 석민 안드로이드
	}

	isTest := true

	pg, err := ioc.Repo.ProductGroups.Get(timedeal.ID.Hex())
	if err != nil {
		return err
	}

	alloffpds := pg.Products

	maxDiscountRate := 0
	for _, pd := range alloffpds {
		pdDao, _ := ioc.Repo.Products.Get(pd.ProductID.Hex())
		if maxDiscountRate < pdDao.DiscountRate {
			maxDiscountRate = pdDao.DiscountRate
		}
	}

	if maxDiscountRate == 0 {
		log.Println("Max discount rate " + strconv.Itoa(maxDiscountRate) + "% for timedeal " + timedeal.ID.Hex())
	}

	// message := timedeal.ShortTitle + " 최대 " + strconv.Itoa(maxDiscountRate) + "% 할인"
	message := "(광고) 1년차 아울렛 최대 83% 할인"

	if isTest {
		timedealNoti := &domain.NotificationDAO{
			ID:               primitive.NewObjectID(),
			Status:           domain.NOTIFICATION_READY,
			NotificationType: domain.NOTIFICATION_TIMEDEAL_OPEN_NOTIFICATION,
			Title:            pushtitle,
			Message:          message,
			DeviceIDs:        adminDeviceIDs,
			NavigateTo:       "/timedeals",
			ReferenceID:      "/" + timedeal.ID.Hex(),
			Notificationid:   "/timedeals/" + timedeal.ID.Hex(),
			Created:          time.Now(),
			Updated:          time.Now(),
		}

		_, err := ioc.Repo.Notifications.Insert(timedealNoti)
		if err != nil {
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
				DeviceIDs:        deviceIDs,
				NavigateTo:       "/timedeals",
				ReferenceID:      "/" + timedeal.ID.Hex(),
				Notificationid:   "/timedeals/" + timedeal.ID.Hex(),
				Created:          time.Now(),
				Updated:          time.Now(),
			}

			_, err := ioc.Repo.Notifications.Insert(&timedealNoti)
			if err != nil {
				log.Println(err)
				return err
			}
			deviceIDs = []string{}
		}
	}

	timedealNoti := &domain.NotificationDAO{
		ID:               primitive.NewObjectID(),
		Status:           domain.NOTIFICATION_READY,
		NotificationType: domain.NOTIFICATION_TIMEDEAL_OPEN_NOTIFICATION,
		Title:            pushtitle,
		Message:          message,
		DeviceIDs:        deviceIDs,
		NavigateTo:       "/timedeals",
		ReferenceID:      "/" + timedeal.ID.Hex(),
		Notificationid:   "/timedeals/" + timedeal.ID.Hex(),
		Created:          time.Now(),
		Updated:          time.Now(),
	}

	_, err = ioc.Repo.Notifications.Insert(timedealNoti)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
