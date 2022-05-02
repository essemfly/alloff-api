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

type productGroupRepo struct {
	col *mongo.Collection
}

func (repo *productGroupRepo) Get(ID string) (*domain.ProductGroupDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	pgId, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": pgId}
	var productGroup *domain.ProductGroupDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&productGroup); err != nil {
		return nil, err
	}
	return productGroup, nil

}

func (repo *productGroupRepo) List(offset, limit int, groupType *domain.ProductGroupType, keyword string) ([]*domain.ProductGroupDAO, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{"grouptype": &groupType}
	options := options.Find()
	options.SetSort(bson.D{{Key: "finishtime", Value: -1}})
	options.SetSkip(int64(offset))
	options.SetLimit(int64(limit))

	if keyword != "" {
		filter["$or"] = []bson.M{
			{"title": primitive.Regex{Pattern: keyword, Options: "i"}},
			{"shorttitle": primitive.Regex{Pattern: keyword, Options: "i"}},
			{"_id": primitive.Regex{Pattern: keyword, Options: "i"}},
		}
	}

	cur, err := repo.col.Find(ctx, filter, options)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	totalCount, _ := repo.col.CountDocuments(ctx, filter)
	var productGroups []*domain.ProductGroupDAO
	err = cur.All(ctx, &productGroups)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	return productGroups, int(totalCount), nil
}

func (repo *productGroupRepo) Upsert(pg *domain.ProductGroupDAO) (*domain.ProductGroupDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	newProductGroupId := ""
	if pg.ID != primitive.NilObjectID {
		filter := bson.M{"_id": pg.ID}
		_, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &pg}, opts)
		if err != nil {
			log.Println(err)
			return pg, nil
		}
		newProductGroupId = pg.ID.Hex()
	} else {
		pg.ID = primitive.NewObjectID()
		insertedId, err := repo.col.InsertOne(ctx, pg)
		if err != nil {
			log.Println(err)
			return pg, nil
		}
		newProductGroupId = insertedId.InsertedID.(primitive.ObjectID).Hex()
	}

	newPg, _ := repo.Get(newProductGroupId)
	return newPg, nil
}

func MongoProductGroupsRepo(conn *MongoDB) repository.ProductGroupsRepository {
	return &productGroupRepo{
		col: conn.productGroupCol,
	}
}
