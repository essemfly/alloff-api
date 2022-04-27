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

func (repo *productGroupRepo) List(numPassedItem int) ([]*domain.ProductGroupDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	now := primitive.NewDateTimeFromTime(time.Now())
	filter := bson.M{"finishtime": bson.M{"$gte": now}, "grouptype": domain.PRODUCT_GROUP_TIMEDEAL}
	onGoingOptions := options.Find()
	onGoingOptions.SetSort(bson.D{{Key: "starttime", Value: 1}})
	cur, err := repo.col.Find(ctx, filter, onGoingOptions)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var productGroups []*domain.ProductGroupDAO
	err = cur.All(ctx, &productGroups)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	outDateFilter := bson.M{"finishtime": bson.M{"$lt": now}}
	outDateOptions := options.Find()
	outDateOptions.SetSort(bson.D{{Key: "finishtime", Value: -1}})
	outDateOptions.SetLimit(int64(numPassedItem)) // Out date timedeals 10개 제한

	cur, err = repo.col.Find(ctx, outDateFilter, outDateOptions)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var outdatedProductGroups []*domain.ProductGroupDAO
	err = cur.All(ctx, &outdatedProductGroups)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	productGroups = append(productGroups, outdatedProductGroups...)

	return productGroups, nil
}

func (repo *productGroupRepo) ListTimedeals(offset, limit int, isLive bool) ([]*domain.ProductGroupDAO, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if isLive {
		now := primitive.NewDateTimeFromTime(time.Now())
		filter := bson.M{"finishtime": bson.M{"$gte": now}, "grouptype": domain.PRODUCT_GROUP_BRAND_TIMEDEAL}
		onGoingOptions := options.Find()
		onGoingOptions.SetSort(bson.D{{Key: "starttime", Value: 1}})
		cur, err := repo.col.Find(ctx, filter, onGoingOptions)
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

	outDateFilter := bson.M{"grouptype": domain.PRODUCT_GROUP_BRAND_TIMEDEAL}
	outDateOptions := options.Find()
	outDateOptions.SetSort(bson.D{{Key: "finishtime", Value: -1}})
	outDateOptions.SetSkip(int64(offset))
	outDateOptions.SetLimit(int64(limit))

	cur, err := repo.col.Find(ctx, outDateFilter, outDateOptions)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	totalCount, _ := repo.col.CountDocuments(ctx, outDateFilter)
	var productGroups []*domain.ProductGroupDAO
	err = cur.All(ctx, &productGroups)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	return productGroups, int(totalCount), nil
}

func (repo *productGroupRepo) ListExhibitionPg(offset, limit int, keyword string) ([]*domain.ProductGroupDAO, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{"grouptype": domain.PRODUCT_GROUP_EXHIBITION}
	outDateOptions := options.Find()
	outDateOptions.SetSort(bson.D{{Key: "finishtime", Value: -1}})
	outDateOptions.SetSkip(int64(offset))
	outDateOptions.SetLimit(int64(limit))

	if keyword != "" {
		filter["$or"] = []bson.M{
			{"title": primitive.Regex{Pattern: keyword, Options: "i"}},
			{"shorttitle": primitive.Regex{Pattern: keyword, Options: "i"}},
			{"_id": primitive.Regex{Pattern: keyword, Options: "i"}},
		}
	}

	totalCount, _ := repo.col.CountDocuments(ctx, filter)
	cur, err := repo.col.Find(ctx, filter, outDateOptions)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	var productGroups []*domain.ProductGroupDAO
	err = cur.All(ctx, &productGroups)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	return productGroups, int(totalCount), nil
}

func (repo *productGroupRepo) ListPgInExhibition(exhibitionID string) ([]*domain.ProductGroupDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{"exhibitionid": exhibitionID}
	cur, err := repo.col.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var productGroups []*domain.ProductGroupDAO
	err = cur.All(ctx, &productGroups)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return productGroups, nil
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
