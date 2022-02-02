package mongo

import (
	"context"
	"time"

	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type crawlSourceRepo struct {
	col *mongo.Collection
}

type crawlRecordRepo struct {
	col *mongo.Collection
}

func (repo *crawlSourceRepo) List(filter interface{}) ([]*domain.CrawlSourceDAO, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	options := options.Find()

	totalCount, _ := repo.col.CountDocuments(ctx, filter)
	cursor, err := repo.col.Find(ctx, filter, options)
	if err != nil {
		return nil, 0, err
	}

	var sources []*domain.CrawlSourceDAO
	err = cursor.All(ctx, &sources)
	if err != nil {
		return nil, 0, err
	}

	return sources, int(totalCount), nil
}

func (repo *crawlSourceRepo) Upsert(source *domain.CrawlSourceDAO) (*domain.CrawlSourceDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)

	filter := bson.M{"maincategorykey": source.MainCategoryKey, "brandidentifier": source.BrandIdentifier, "category.catidentifier": source.Category.CatIdentifier, "brandkeyname": source.BrandKeyname}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": source}, opts); err != nil {
		return nil, err
	}

	var updatedSource *domain.CrawlSourceDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&updatedSource); err != nil {
		return nil, err
	}

	return updatedSource, nil
}

func MongoCrawlSourcesRepo(conn *MongoDB) repository.CrawlSourcesRepository {
	return &crawlSourceRepo{
		col: conn.crawlSourceCol,
	}
}

func (repo *crawlRecordRepo) GetLast() (*domain.CrawlRecordDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	options := options.FindOne()
	options.SetSort(bson.M{"_id": -1})

	var lastRecord domain.CrawlRecordDAO
	err := repo.col.FindOne(ctx, bson.M{}, options).Decode(&lastRecord)
	if err != nil {
		return nil, err
	}

	return &lastRecord, nil
}

func (repo *crawlRecordRepo) Insert(crawlRecord *domain.CrawlRecordDAO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.col.InsertOne(ctx, bson.M{"date": time.Now(), "crawledbrands": crawlRecord.CrawledBrands, "newproducts": crawlRecord.NewProducts, "outproducts": crawlRecord.OldProducts})

	return err
}

func MongoCrawlRecordRepo(conn *MongoDB) repository.CrawlRecordsRepository {
	return &crawlRecordRepo{
		col: conn.crawlRecordCol,
	}
}
