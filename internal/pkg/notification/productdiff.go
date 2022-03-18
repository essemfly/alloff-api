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
				messages += utils.PriceFormatter(like.LastPrice) + " > " + utils.PriceFormatter(newProduct.DiscountedPrice)
				messages += ", 지금 확인해보세요!"
				productDiffNotification := domain.NotificationDAO{
					ID:               primitive.NewObjectID(),
					Status:           domain.NOTIFICATION_READY,
					NotificationType: domain.NOTIFICATION_PRODUCT_DIFF_NOTIFICATION,
					Title:            "찜한 상품의 가격이 내려갔어요! 🔻",
					Message:          messages,
					DeviceIDs:        deviceIDs,
					NavigateTo:       "/products",
					Notificationid:   "/products" + "/" + newProduct.ID.Hex(),
					ReferenceID:      "/" + newProduct.ID.Hex(),
					Created:          time.Now(),
					Updated:          time.Now(),
				}

				_, err := ioc.Repo.Notifications.Insert(&productDiffNotification)
				if err != nil {
					log.Println(err)
					return err
				}

				like.IsPushed = true
				_, err = ioc.Repo.LikeProducts.Update(like)
				if err != nil {
					log.Println("err occured in update like products pushed")
					return err
				}
			}
		}
	}

	return nil
}
