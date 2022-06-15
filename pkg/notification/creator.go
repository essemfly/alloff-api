package notification

import (
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateNotiRequest struct {
	NotiType    domain.NotificationType
	Title       string
	Message     string
	ReferenceID string
}

func CreateNotification(request *CreateNotiRequest) (*domain.NotificationDAO, error) {
	notiDao := &domain.NotificationDAO{
		ID:      primitive.NewObjectID(),
		Status:  domain.NOTIFICATION_READY,
		Title:   request.Title,
		Message: request.Message,
	}
	switch request.NotiType {
	case domain.NOTIFICATION_GENERAL_NOTIFICATION:
		notiDao.NotificationType = domain.NOTIFICATION_GENERAL_NOTIFICATION
		notiDao.Notificationid = "/general?" + utils.CreateShortUUID()
		notiDao.ReferenceID = request.ReferenceID + "?notiType=general&title=" + request.Title + "&message=" + request.Message
		notiDao.NavigateTo = "/"
	case domain.NOTIFICATION_PRODUCT_DIFF_NOTIFICATION:
		notiDao.NotificationType = domain.NOTIFICATION_GENERAL_NOTIFICATION
		notiDao.Notificationid = "/productdiff?" + utils.CreateShortUUID()
		notiDao.ReferenceID = request.ReferenceID + "?notiType=product&title=" + request.Title + "&message=" + request.Message
		notiDao.NavigateTo = "/products"
	case domain.NOTIFICATION_TIMEDEAL_OPEN_NOTIFICATION:
		notiDao.NotificationType = domain.NOTIFICATION_TIMEDEAL_OPEN_NOTIFICATION
		notiDao.Notificationid = "/deals?id=" + request.ReferenceID + "?uuid=" + utils.CreateShortUUID()
		notiDao.ReferenceID = "/" + request.ReferenceID + "?notiType=dealopen&title=" + request.Title + "&message=" + request.Message
		notiDao.NavigateTo = "/deals"
	}

	return ioc.Repo.Notifications.Insert(notiDao)
}
