package mongo

import (
	"context"
	"time"

	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type notificationRepo struct {
	col *mongo.Collection
}

func (repo *notificationRepo) Insert(noti *domain.NotificationDAO) (*domain.NotificationDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	noti.Created = time.Now()
	noti.Updated = time.Now()
	_, err := repo.col.InsertOne(ctx, noti)

	if err != nil {
		return nil, err
	}

	return noti, nil
}
func (repo *notificationRepo) List(onlyReady bool) ([]*domain.NotificationDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	now := time.Now()

	filter := bson.M{"notificationtype": "PRODUCT_DIFF_NOTIFICATION", "scheduleddate": bson.M{"$lte": now}}
	if onlyReady {
		filter["status"] = domain.NOTIFICATION_READY
	}

	cursor, err := repo.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var notis []*domain.NotificationDAO
	err = cursor.All(ctx, &notis)
	if err != nil {
		return nil, err
	}

	return notis, nil
}
func (repo *notificationRepo) Update(noti *domain.NotificationDAO) (*domain.NotificationDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err := repo.col.UpdateOne(ctx, bson.M{"_id": noti.ID}, bson.M{"$set": bson.M{"status": domain.NOTIFICATION_SUCCEEDED, "updated": time.Now(), "sended": time.Now()}})
	if err != nil {
		return nil, err
	}

	return noti, nil
}

func MongoNotificationsRepo(conn *MongoDB) repository.NotificationsRepository {
	return &notificationRepo{
		col: conn.notificationCol,
	}
}
