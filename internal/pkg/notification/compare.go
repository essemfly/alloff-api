package notification

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertDiffNotification(newProduct *domain.ProductDAO, oldPrice int) error {
	likes, err := ioc.Repo.LikeProducts.ListProductsLike(newProduct.ID.Hex())
	if err != nil {
		return err
	}

	for _, like := range likes {
		deviceIDs := []string{}
		// Admins
		// adminDeviceIDs := []string{
		// 	"c0IfW-xzdUx3nu6XmpzXNz:APA91bGYmuaoWsdoNzhp_SqUiaKm_8bPK-MfUgmJpmvprUui6run4qKTEeF8QpPSnrR7f5oK2Suy5BJduO04C2DuKWmYHbYZJCRy_FtI6Rm6kAxEopwiRSGjymiqGzXVpz8i8nffFYV1", // 주영
		// 	"dJ4rSV2L60aLpHxQjxoFz-:APA91bGJR2YXQkmBbv0rELM6caPUeZ3C1MnBkz1wlI68wCzDRhc9Bsma3stSCRXTGip-6mxdtj2GfuMT4c0XV85AWDuLr0lkH33VDNRsuc8nqo24JGOHmDDpYl_wetLh9vYL-3I0A6ID", // 석민
		// 	"cBRNxXqZwEeapVVni8KJcG:APA91bG7UnRCfxWvNR7ngSYNfhTazApx9yAlQXtXqCJDWpn_X-cwVMnnUDLmjLUCso9s7_oiP_xrBkOqoa-1ie3LaRsckENluZTaxWcNAKpdUvVZtV9Pq_TRgRdmwtpA0kE_-Mx-_Nzl", // 명규
		// 	"eEx9IUeGQtGgo__aVsJih4:APA91bG0NaGdVP_2NyfBWAXXXdZ0-iR9J-pRTf6JxjueMowaIg9HQDYiZ5PkGeqTypblQNDg-Fecdask8W6o15lNHOhzCHfHzIzofE63APEZPNQQbRVjaaBCTk6m3kAeIqoBuJl9rmHh", // 석민 안드로이드
		// }

		// Admins_DEV
		// adminDeviceIDs = []string{
		// 	"eK15WFKbUERjlijXdXzyQz:APA91bGSClbWJAxwfWuDEQFQ4povXbbZhmOQNynZqIxvvKeVjrmNJkEoV_qlUrK8832dH3gML0Ltk4Ll9zI8syyLdsdYXy1KgBmFL6gDibFdmGwIea-q6Z3HOa9NAOjqXaBZJf9iPdjw",
		// }

		if like.OldProduct.DiscountedPrice > newProduct.DiscountedPrice {
			devices, err := ioc.Repo.Devices.ListAllowedByUser(like.Userid)
			if err != nil {
				log.Println("User " + like.Userid + " failed")
				continue
			}
			for _, device := range devices {
				deviceIDs = append(deviceIDs, device.DeviceId)
			}

			if len(deviceIDs) > 0 {
				messages := "[" + newProduct.ProductInfo.Brand.KorName + "]" + newProduct.AlloffName + "\n"
				messages += utils.PriceFormatter(like.OldProduct.DiscountedPrice) + " > " + utils.PriceFormatter(newProduct.DiscountedPrice)
				messages += ", 지금 확인해보세요!"
				productDiffNotification := domain.NotificationDAO{
					ID:               primitive.NewObjectID(),
					Status:           domain.NOTIFICATION_READY,
					NotificationType: domain.NOTIFICATION_PRODUCT_DIFF_NOTIFICATION,
					Title:            "찜한 상품의 가격이 내려갔어요! 🔻",
					Message:          messages,
					DeviceIDs:        deviceIDs,
					NavigateTo:       "/products",
					ReferenceID:      "/" + newProduct.ID.Hex(),
					ScheduledDate:    time.Now(),
					Created:          time.Now(),
					Updated:          time.Now(),
				}

				_, err := ioc.Repo.Notifications.Insert(&productDiffNotification)
				if err != nil {
					log.Println(err)
					return err
				}

				for _, like := range likes {
					_, err = ioc.Repo.LikeProducts.Update(like)
					if err != nil {
						log.Println("err occured in update like products pushed")
						return err
					}
				}
			}
		}
	}

	return nil
}
