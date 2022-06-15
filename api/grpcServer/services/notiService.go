package services

import (
	"context"
	"errors"

	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/notification"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
	"go.uber.org/zap"
)

const CHUNK_SIZE = 500

type NotiService struct {
	grpcServer.NotificationServer
}

// (TODO) 현재는 timedeal 만 생성 할 수 있다. + 이 코드는 pkg안의 코드로 바껴서 사용되어야 한다.
func (s *NotiService) CreateNoti(ctx context.Context, req *grpcServer.CreateNotiRequest) (*grpcServer.CreateNotiResponse, error) {
	notiRequest := notification.CreateNotiRequest{
		Title:       req.Title,
		Message:     req.Message,
		ReferenceID: req.ReferenceId,
	}

	if req.NotiType == string(domain.NOTIFICATION_GENERAL_NOTIFICATION) {
		notiRequest.NotiType = domain.NOTIFICATION_GENERAL_NOTIFICATION
	} else if req.NotiType == string(domain.NOTIFICATION_PRODUCT_DIFF_NOTIFICATION) {
		notiRequest.NotiType = domain.NOTIFICATION_PRODUCT_DIFF_NOTIFICATION
	} else if req.NotiType == string(domain.NOTIFICATION_EXHIBITION_OPEN_NOTIFICATION) {
		notiRequest.NotiType = domain.NOTIFICATION_TIMEDEAL_OPEN_NOTIFICATION
	} else {
		return nil, errors.New("invalid notification type")
	}

	_, err := notification.CreateNotification(&notiRequest)
	if err != nil {
		config.Logger.Error("Err occured when create notification", zap.Error(err))
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
	notis, err := ioc.Repo.Notifications.ListByNotiID(req.NotificationId)
	if err != nil {
		return &grpcServer.SendNotiResponse{
			IsSent: false,
		}, err
	}

	// Notis를 장기적으로 slice가 아닌 단건으로 바꿀 예정
	for _, noti := range notis {
		go notification.Send(noti)
	}

	return &grpcServer.SendNotiResponse{
		IsSent: true,
	}, nil
}
