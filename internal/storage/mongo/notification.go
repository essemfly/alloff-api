package mongo

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type notificationRepo struct {
	col *mongo.Collection
}

func (repo *notificationRepo) Insert(*domain.NotificationDAO) (*domain.NotificationDAO, error) {
	return nil, nil
}
func (repo *notificationRepo) List(onlyReady bool) ([]*domain.NotificationDAO, error) {
	return nil, nil
}
func (repo *notificationRepo) Update(*domain.NotificationDAO) (*domain.NotificationDAO, error) {
	return nil, nil
}

func MongoNotificationsRepo(conn *MongoDB) repository.NotificationsRepository {
	return &notificationRepo{
		col: conn.notificationCol,
	}
}
