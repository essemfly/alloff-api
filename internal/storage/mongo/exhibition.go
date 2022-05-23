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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type exhibitionRepo struct {
	col *mongo.Collection
}

func (repo *exhibitionRepo) Get(ID string) (*domain.ExhibitionDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	exhibitionID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": exhibitionID}
	var exhibition *domain.ExhibitionDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&exhibition); err != nil {
		return nil, err
	}
	return exhibition, nil

}

func (repo *exhibitionRepo) List(offset, limit int, onlyLive bool, exhibitionStatus domain.ExhibitionStatus, exhibitionType domain.ExhibitionType, query string) ([]*domain.ExhibitionDAO, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	now := primitive.NewDateTimeFromTime(time.Now())
	filter := bson.M{
		"exhibitiontype": exhibitionType,
	}
	onGoingOptions := options.Find()
	onGoingOptions.SetSkip(int64(offset))
	onGoingOptions.SetLimit(int64(limit))
	//onGoingOptions.SetSort(sortingOptions)

	switch exhibitionStatus {
	case domain.EXHIBITION_NOTOPEN:
		filter["starttime"] = bson.M{"$gte": now}
		filter["islive"] = true
		onGoingOptions.SetSort(bson.D{{Key: "starttime", Value: -1}})
	case domain.EXHIBITION_LIVE:
		filter["finishtime"] = bson.M{"$gte": now}
		filter["islive"] = true
		onGoingOptions.SetSort(bson.D{{Key: "starttime", Value: -1}})
	case domain.EXHIBITION_CLOSED:
		filter["finishtime"] = bson.M{"$lte": now}
		onGoingOptions.SetSort(bson.D{{Key: "starttime", Value: -1}})
	}

	if onlyLive {
		filter["finishtime"] = bson.M{"$gte": now}
		filter["islive"] = true
		onGoingOptions.SetSort(bson.D{{Key: "starttime", Value: -1}})
	} else {
		onGoingOptions.SetSort(bson.D{{Key: "starttime", Value: -1}})
	}

	if query != "" {
		filter["$or"] = []bson.M{
			{"title": primitive.Regex{Pattern: query, Options: "i"}},
			{"subtitle": primitive.Regex{Pattern: query, Options: "i"}},
		}
	}

	totalCount, _ := repo.col.CountDocuments(ctx, filter)
	cur, err := repo.col.Find(ctx, filter, onGoingOptions)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	var exhibitions []*domain.ExhibitionDAO
	err = cur.All(ctx, &exhibitions)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	return exhibitions, int(totalCount), nil
}

func (repo *exhibitionRepo) Upsert(exhibition *domain.ExhibitionDAO) (*domain.ExhibitionDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	exhibition.UpdatedAt = time.Now()
	opts := options.Update().SetUpsert(true)
	newExhibitionID := ""
	if exhibition.ID != primitive.NilObjectID {
		filter := bson.M{"_id": exhibition.ID}
		_, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &exhibition}, opts)
		if err != nil {
			log.Println(err)
			return exhibition, nil
		}
		newExhibitionID = exhibition.ID.Hex()
	} else {
		exhibition.ID = primitive.NewObjectID()
		exhibition.CreatedAt = time.Now()
		insertedId, err := repo.col.InsertOne(ctx, *exhibition)
		if err != nil {
			log.Println(err)
			return exhibition, nil
		}
		newExhibitionID = insertedId.InsertedID.(primitive.ObjectID).Hex()
	}

	newPg, _ := repo.Get(newExhibitionID)
	return newPg, nil
}

func (repo *exhibitionRepo) ListGroupDeals(offset, limit int, onlyLive bool, exhibitionStatus domain.GroupdealStatus) ([]*domain.ExhibitionDAO, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{
		"exhibitiontype": domain.EXHIBITION_GROUPDEAL,
	}
	onGoingOptions := options.Find()
	onGoingOptions.SetSkip(int64(offset))
	onGoingOptions.SetLimit(int64(limit))

	if onlyLive {
		filter["islive"] = true
		onGoingOptions.SetSort(bson.D{{Key: "starttime", Value: -1}})
	} else {
		onGoingOptions.SetSort(bson.D{{Key: "starttime", Value: -1}})
	}

	now := primitive.NewDateTimeFromTime(time.Now())

	if exhibitionStatus == domain.GROUPDEAL_PENDING {
		filter["starttime"] = bson.M{"$gte": now}
		filter["recruitstarttime"] = bson.M{"$lte": now}
	} else if exhibitionStatus == domain.GROUPDEAL_OPEN {
		filter["starttime"] = bson.M{"$lte": now}
		filter["finishtime"] = bson.M{"$gte": now}
	} else if exhibitionStatus == domain.GROUPDEAL_CLOSED {
		filter["finishtime"] = bson.M{"$lte": now}
	}

	totalCount, _ := repo.col.CountDocuments(ctx, filter)
	cur, err := repo.col.Find(ctx, filter, onGoingOptions)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	var exhibitions []*domain.ExhibitionDAO
	err = cur.All(ctx, &exhibitions)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	return exhibitions, int(totalCount), nil
}

func MongoExhibitionsRepo(conn *MongoDB) repository.ExhibitionsRepository {
	return &exhibitionRepo{
		col: conn.exhibitionCol,
	}
}
