package notification

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertBrandOpenNotification(brands []*domain.BrandDAO) error {

	devices, err := ioc.Repo.Devices.ListAllowed()
	if err != nil {
		return err
	}

	deviceIDs := []string{}
	for _, device := range devices {
		deviceIDs = append(deviceIDs, device.DeviceId)
	}

	// Admins
	// adminDeviceIDs = []string{
	// 	"c0IfW-xzdUx3nu6XmpzXNz:APA91bGYmuaoWsdoNzhp_SqUiaKm_8bPK-MfUgmJpmvprUui6run4qKTEeF8QpPSnrR7f5oK2Suy5BJduO04C2DuKWmYHbYZJCRy_FtI6Rm6kAxEopwiRSGjymiqGzXVpz8i8nffFYV1", // 주영
	// 	"dJ4rSV2L60aLpHxQjxoFz-:APA91bGJR2YXQkmBbv0rELM6caPUeZ3C1MnBkz1wlI68wCzDRhc9Bsma3stSCRXTGip-6mxdtj2GfuMT4c0XV85AWDuLr0lkH33VDNRsuc8nqo24JGOHmDDpYl_wetLh9vYL-3I0A6ID", // 석민
	// 	"cBRNxXqZwEeapVVni8KJcG:APA91bG7UnRCfxWvNR7ngSYNfhTazApx9yAlQXtXqCJDWpn_X-cwVMnnUDLmjLUCso9s7_oiP_xrBkOqoa-1ie3LaRsckENluZTaxWcNAKpdUvVZtV9Pq_TRgRdmwtpA0kE_-Mx-_Nzl", // 명규
	// 	"eEx9IUeGQtGgo__aVsJih4:APA91bG0NaGdVP_2NyfBWAXXXdZ0-iR9J-pRTf6JxjueMowaIg9HQDYiZ5PkGeqTypblQNDg-Fecdask8W6o15lNHOhzCHfHzIzofE63APEZPNQQbRVjaaBCTk6m3kAeIqoBuJl9rmHh", // 석민 안드로이드
	// }

	brandsList := []string{}

	for _, brand := range brands {
		brandsList = append(brandsList, brand.KorName)
	}

	brandsString := strings.Join(brandsList, ", ")

	notiID := "/" + time.Now().String()
	timedealNoti := domain.NotificationDAO{
		ID:               primitive.NewObjectID(),
		Status:           domain.NOTIFICATION_READY,
		NotificationType: domain.NOTIFICATION_BRAND_OPEN_NOTIFICATION,
		Title:            "신규 브랜드 " + strconv.Itoa(len(brandsList)) + "개가 오픈했어요! 🎉",
		Message:          brandsString,
		DeviceIDs:        deviceIDs,
		NavigateTo:       "/home",
		ReferenceID:      notiID,
		Notificationid:   "/home" + notiID,
		Created:          time.Now(),
		Updated:          time.Now(),
	}

	_, err = ioc.Repo.Notifications.Insert(&timedealNoti)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
