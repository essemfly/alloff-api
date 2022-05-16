package main

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/cmd"
	"github.com/lessbutter/alloff-api/cmd/productDiffNotifier/scripts"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	cmd.SetBaseConfig()
	scripts.InsertDiffNotifi()
	scripts.SendProductDiffNoti()

	// TestProductDiffNoti(pd)
}

func TestProductDiffNoti(product *domain.ProductDAO) {
	mobiles := []string{"01097711882"}
	deviceIDs := []string{}
	for _, mobile := range mobiles {
		user, _ := ioc.Repo.Users.GetByMobile(mobile)
		devices, _ := ioc.Repo.Devices.ListAllowedByUser(user.ID.Hex())
		for _, device := range devices {
			deviceIDs = append(deviceIDs, device.DeviceId)
		}
	}

	messages := "[" + product.ProductInfo.Brand.KorName + "]" + product.AlloffName + "\n"
	messages += utils.PriceFormatter(product.OriginalPrice) + " > " + utils.PriceFormatter(product.DiscountedPrice)
	messages += ", 지금 확인해보세요!"

	productDiffNotification := domain.NotificationDAO{
		ID:               primitive.NewObjectID(),
		Status:           domain.NOTIFICATION_READY,
		NotificationType: domain.NOTIFICATION_PRODUCT_DIFF_NOTIFICATION,
		Title:            "찜한 상품의 가격이 내려갔어요! 🔻",
		Message:          messages,
		DeviceIDs:        deviceIDs,
		NavigateTo:       "/products",
		Notificationid:   "/products" + "/" + product.ID.Hex(),
		ReferenceID:      "/" + product.ID.Hex(),
		Created:          time.Now(),
		Updated:          time.Now(),
	}

	_, err := ioc.Repo.Notifications.Insert(&productDiffNotification)
	if err != nil {
		log.Println(err)
		return
	}

}
