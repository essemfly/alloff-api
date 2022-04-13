package mongo

import (
	"context"
	"time"

	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (repo *notificationRepo) Get(notiID string) ([]*domain.NotificationDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	notiObjID, _ := primitive.ObjectIDFromHex(notiID)

	noti := &domain.NotificationDAO{}
	if err := repo.col.FindOne(ctx, bson.M{"_id": notiObjID}).Decode(noti); err != nil {
		return nil, err
	}

	notis := []*domain.NotificationDAO{}
	cursor, err := repo.col.Find(ctx, bson.M{"notificationid": noti.Notificationid})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &notis)
	if err != nil {
		return nil, err
	}

	return notis, nil
}

func (repo *notificationRepo) List(offset, limit int, notiTypes []domain.NotificationType, onlyReady bool) ([]*domain.NotificationDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{}
	notiFilters := []bson.M{}
	for _, notiType := range notiTypes {
		notiFilters = append(notiFilters, bson.M{
			"notificationtype": notiType,
		})
	}
	if onlyReady {
		filter["$or"] = notiFilters
		filter["status"] = domain.NOTIFICATION_READY
	}

	options := options.Find()
	options.SetSkip(int64(offset))
	options.SetLimit(int64(limit))
	options.SetSort(bson.M{"_id": -1})

	cursor, err := repo.col.Find(ctx, filter, options)
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

	_, err := repo.col.UpdateOne(ctx, bson.M{"_id": noti.ID}, bson.M{"$set": &noti})
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
