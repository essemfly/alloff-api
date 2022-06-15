package mongo

import (
	"context"
	"log"
	"time"

	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (repo *notificationRepo) Get(ID string) (*domain.NotificationDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	notiObjID, _ := primitive.ObjectIDFromHex(ID)

	noti := &domain.NotificationDAO{}
	if err := repo.col.FindOne(ctx, bson.M{"_id": notiObjID}).Decode(noti); err != nil {
		return nil, err
	}

	return noti, nil
}

func (repo *notificationRepo) ListByNotiID(notiID string) ([]*domain.NotificationDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	notis := []*domain.NotificationDAO{}
	cursor, err := repo.col.Find(ctx, bson.M{"notificationid": notiID})
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

	notiFilters := []bson.M{}
	for _, notiType := range notiTypes {
		notiFilters = append(notiFilters, bson.M{
			"notificationtype": notiType,
		})
	}
	filter := bson.M{
		"$or": notiFilters,
	}
	if onlyReady {
		filter["status"] = domain.NOTIFICATION_READY
	}

	pipelines := []bson.M{
		{
			"$match": filter,
		},
		{
			"$sort": bson.M{"_id": 1},
		},
		{
			"$group": bson.M{
				"_id":         "$notificationid",
				"notidao":     bson.M{"$last": "$$ROOT"},
				"totalpushed": bson.M{"$sum": "$numuserspushed"},
				"totalfailed": bson.M{"$sum": "$numusersfailed"},
			},
		},
		{
			"$sort": bson.M{"notidao._id": -1},
		},
		{
			"$limit": offset + limit,
		},
		{
			"$skip": offset,
		},
	}
	cursor, err := repo.col.Aggregate(ctx, pipelines, nil)

	if err != nil {
		log.Println("err on notification ", err)
		return nil, err
	}

	var notis []*domain.DistinctNotiResult
	err = cursor.All(ctx, &notis)
	if err != nil {
		return nil, err
	}

	notiDaos := []*domain.NotificationDAO{}
	for _, noti := range notis {
		noti.NotiDAO.NumUsersFailed = noti.TotalFailed
		noti.NotiDAO.NumUsersPushed = noti.TotalPushed
		notiDaos = append(notiDaos, &noti.NotiDAO)
	}
	return notiDaos, nil
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
