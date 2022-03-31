package mapper

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/golang/protos"
)

func NotificationMapper(noti *domain.NotificationDAO) *grpcServer.NotificationMessage {
	return &grpcServer.NotificationMessage{
		NotificationId: noti.ID.Hex(),
		Status:         string(noti.Status),
		NotiType:       string(noti.NotificationType),
		ReferenceId:    noti.ReferenceID,
		Title:          noti.Title,
		Message:        noti.Message,
		SendedAt:       noti.Sended.String(),
	}
}
