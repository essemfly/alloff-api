package scripts

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTimedealNotification(timedealID, pushtitle, message string) error {
	// Admins
	adminDeviceIDs := []string{
		"c0IfW-xzdUx3nu6XmpzXNz:APA91bGYmuaoWsdoNzhp_SqUiaKm_8bPK-MfUgmJpmvprUui6run4qKTEeF8QpPSnrR7f5oK2Suy5BJduO04C2DuKWmYHbYZJCRy_FtI6Rm6kAxEopwiRSGjymiqGzXVpz8i8nffFYV1", // 주영
		"dJ4rSV2L60aLpHxQjxoFz-:APA91bGJR2YXQkmBbv0rELM6caPUeZ3C1MnBkz1wlI68wCzDRhc9Bsma3stSCRXTGip-6mxdtj2GfuMT4c0XV85AWDuLr0lkH33VDNRsuc8nqo24JGOHmDDpYl_wetLh9vYL-3I0A6ID", // 석민
		"cBRNxXqZwEeapVVni8KJcG:APA91bG7UnRCfxWvNR7ngSYNfhTazApx9yAlQXtXqCJDWpn_X-cwVMnnUDLmjLUCso9s7_oiP_xrBkOqoa-1ie3LaRsckENluZTaxWcNAKpdUvVZtV9Pq_TRgRdmwtpA0kE_-Mx-_Nzl", // 명규
		"eEx9IUeGQtGgo__aVsJih4:APA91bG0NaGdVP_2NyfBWAXXXdZ0-iR9J-pRTf6JxjueMowaIg9HQDYiZ5PkGeqTypblQNDg-Fecdask8W6o15lNHOhzCHfHzIzofE63APEZPNQQbRVjaaBCTk6m3kAeIqoBuJl9rmHh", // 석민 안드로이드
		"eO3-GwnKS0U5iAUTm_13Fk:APA91bH8O6aDYWdXSzlvBKTXpAF9e4sTD4mxzL_wDe0xhkch-ia2gpjPI0IGNTynCle4A5cjl5QZBI3SkBk3PI1N8OVunv_gPAErTk7R47-J2qteM68VVvw9kw6Udiv77b1t2oA3CuYq",
		"eEx9IUeGQtGgo__aVsJih4:APA91bG0NaGdVP_2NyfBWAXXXdZ0-iR9J-pRTf6JxjueMowaIg9HQDYiZ5PkGeqTypblQNDg-Fecdask8W6o15lNHOhzCHfHzIzofE63APEZPNQQbRVjaaBCTk6m3kAeIqoBuJl9rmHh",
	}

	isTest := true

	timedeal, err := ioc.Repo.ProductGroups.Get(timedealID)
	if err != nil {
		log.Println("err", err)
		return err
	}

	if isTest {
		timedealNoti := domain.NotificationDAO{
			ID:               primitive.NewObjectID(),
			Status:           domain.NOTIFICATION_READY,
			NotificationType: domain.NOTIFICATION_TIMEDEAL_OPEN_NOTIFICATION,
			Title:            pushtitle,
			Message:          message,
			Notificationid:   "/timedeals/" + timedeal.ID.Hex(),
			DeviceIDs:        adminDeviceIDs,
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
