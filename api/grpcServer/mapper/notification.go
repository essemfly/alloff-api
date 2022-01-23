package mapper

import (
	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/internal/core/domain"
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
