package services

import (
	"context"
	"log"

	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/notification"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CHUNK_SIZE = 500

type NotiService struct {
	grpcServer.NotificationServer
}

// (TODO) 현재는 timedeal 만 생성 할 수 있다. + 이 코드는 pkg안의 코드로 바껴서 사용되어야 한다.
func (s *NotiService) CreateNoti(ctx context.Context, req *grpcServer.CreateNotiRequest) (*grpcServer.CreateNotiResponse, error) {
	notiDao := &domain.NotificationDAO{
		Status:           domain.NOTIFICATION_READY,
		NotificationType: domain.NOTIFICATION_TIMEDEAL_OPEN_NOTIFICATION,
		Title:            req.Title,
		Message:          req.Message,
		Notificationid:   "/timedeals" + req.ReferenceId,
		ReferenceID:      "/" + req.ReferenceId,
		NavigateTo:       "/timedeals",
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
	notiDaos, err := ioc.Repo.Notifications.List(int(req.Offset), int(req.Limit), onlyReady)

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

		if noti.NotificationType != domain.NOTIFICATION_TIMEDEAL_OPEN_NOTIFICATION {
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
