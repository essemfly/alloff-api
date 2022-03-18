package services

import (
	"context"
	"errors"
	"log"

	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/notification"
	"github.com/lessbutter/alloff-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CHUNK_SIZE = 500

type NotiService struct {
	grpcServer.NotificationServer
}

// (TODO) 현재는 timedeal 만 생성 할 수 있다. + 이 코드는 pkg안의 코드로 바껴서 사용되어야 한다.
func (s *NotiService) CreateNoti(ctx context.Context, req *grpcServer.CreateNotiRequest) (*grpcServer.CreateNotiResponse, error) {
	var notiDao *domain.NotificationDAO
	if req.NotiType == string(domain.NOTIFICATION_GENERAL_NOTIFICATION) {
		notiDao = &domain.NotificationDAO{
			Status:           domain.NOTIFICATION_READY,
			NotificationType: domain.NOTIFICATION_GENERAL_NOTIFICATION,
			Title:            req.Title,
			Message:          req.Message,
			Notificationid:   "/" + utils.CreateShortUUID(),
			ReferenceID:      req.ReferenceId,
			NavigateTo:       "/",
		}
	} else if req.NotiType == string(domain.NOTIFICATION_EXHIBITION_OPEN_NOTIFICATION) {
		notiDao = &domain.NotificationDAO{
			Status:           domain.NOTIFICATION_READY,
			NotificationType: domain.NOTIFICATION_EXHIBITION_OPEN_NOTIFICATION,
			Title:            req.Title,
			Message:          req.Message,
			Notificationid:   "/exhibition" + req.ReferenceId,
			ReferenceID:      "/" + req.ReferenceId,
			NavigateTo:       "/exhibition",
		}
	} else {
		return nil, errors.New("invalid notification type")
	}

	devices, err := ioc.Repo.Devices.ListAllowed()
	if err != nil {
		return nil, err
	}

	deviceIDs := []string{}

	for _, device := range devices {
		deviceIDs = append(deviceIDs, device.DeviceId)
		if len(deviceIDs) > 300 {
			notiDao.ID = primitive.NewObjectID()
			notiDao.DeviceIDs = deviceIDs
			_, err := ioc.Repo.Notifications.Insert(notiDao)
			if err != nil {
				return nil, err
			}

			deviceIDs = []string{}
		}
	}

	notiDao.ID = primitive.NewObjectID()
	notiDao.DeviceIDs = deviceIDs
	_, err = ioc.Repo.Notifications.Insert(notiDao)
	if err != nil {
		return nil, err
	}

	return &grpcServer.CreateNotiResponse{
		Succeeded: true,
	}, nil
}

func (s *NotiService) ListNoti(ctx context.Context, req *grpcServer.ListNotiRequest) (*grpcServer.ListNotiResponse, error) {
	onlyReady := false
	notiTypes := []domain.NotificationType{
		domain.NOTIFICATION_PRODUCT_DIFF_NOTIFICATION,
		domain.NOTIFICATION_BRAND_NEW_PRODUCT_NOTIFICATION,
		domain.NOTIFICATION_TIMEDEAL_OPEN_NOTIFICATION,
		domain.NOTIFICATION_BRAND_OPEN_NOTIFICATION,
		domain.NOTIFICATION_EVENT_NOTIFICATION,
		domain.NOTIFICATION_GENERAL_NOTIFICATION,
		domain.NOTIFICATION_EXHIBITION_OPEN_NOTIFICATION,
	}
	notiDaos, err := ioc.Repo.Notifications.List(int(req.Offset), int(req.Limit), notiTypes, onlyReady)

	notis := []*grpcServer.NotificationMessage{}
	for _, noti := range notiDaos {
		notis = append(notis, mapper.NotificationMapper(noti))
	}
	return &grpcServer.ListNotiResponse{
		Notis: notis,
	}, err
}

func (s *NotiService) SendNoti(ctx context.Context, req *grpcServer.SendNotiRequest) (*grpcServer.SendNotiResponse, error) {
	notis, err := ioc.Repo.Notifications.Get(req.NotificationId)
	if err != nil {
		return nil, err
	}

	for _, noti := range notis {
		if noti.Status != domain.NOTIFICATION_READY {
			continue
		}

		err := notification.SendNotification(noti)
		if err != nil {
			log.Println("err occured in send noti", err)
		}
	}
	return &grpcServer.SendNotiResponse{
		IsSent: true,
	}, nil
}
