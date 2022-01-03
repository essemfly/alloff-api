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

type productRepo struct {
	col *mongo.Collection
}

type productDiffRepo struct {
	col *mongo.Collection
}

type productMetaInfoRepo struct {
	col *mongo.Collection
}

func (repo *productRepo) Get(ID string) (*domain.ProductDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	productObjectId, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": productObjectId}
	var product *domain.ProductDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&product); err != nil {
		return nil, err
	}
	return product, nil
}

func (repo *productRepo) GetByMetaID(MetaID string) (*domain.ProductDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	productObjectId, _ := primitive.ObjectIDFromHex(MetaID)
	filter := bson.M{"brand._id": productObjectId}
	var oldProduct domain.ProductDAO
	err := repo.col.FindOne(ctx, filter).Decode(&oldProduct)
	if err != nil {
		return nil, err
	}

	return &oldProduct, nil
}

func (repo *productRepo) List(limit, offset int, filter, sortingOptions interface{}) ([]*domain.ProductDAO, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSort(sortingOptions)
	options.SetLimit(int64(limit))
	options.SetSkip(int64(offset))

	totalCount, _ := repo.col.CountDocuments(ctx, filter)
	cursor, err := repo.col.Find(ctx, filter, options)
	if err != nil {
		return nil, 0, err
	}

	var products []*domain.ProductDAO
	err = cursor.All(ctx, &products)
	if err != nil {
		return nil, 0, err
	}

	return products, int(totalCount), nil
}

func (repo *productRepo) Insert(product *domain.ProductDAO) (*domain.ProductDAO, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var oldProduct domain.ProductDAO
	oid, err := repo.col.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": oid.InsertedID}
	err = repo.col.FindOne(ctx, filter).Decode(&oldProduct)
	if err != nil {
		return nil, err
	}

	return &oldProduct, nil
}

func (repo *productRepo) Upsert(product *domain.ProductDAO) (*domain.ProductDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": product.ID}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &product}, opts); err != nil {
		return nil, err
	}

	var updatedProduct *domain.ProductDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&updatedProduct); err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func MongoProductsRepo(conn *MongoDB) repository.ProductsRepository {
	return &productRepo{
		col: conn.productCol,
	}
}

func (repo *productDiffRepo) Insert(diff *domain.ProductDiffDAO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err := repo.col.InsertOne(ctx, diff)
	if err != nil {
		return err
	}

	return nil
}

func (repo *productDiffRepo) List(filter interface{}) ([]*domain.ProductDiffDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cursor, err := repo.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var diffs []*domain.ProductDiffDAO
	err = cursor.All(ctx, &diffs)
	if err != nil {
		return nil, err
	}
	return diffs, nil
}

func MongoProductDiffsRepo(conn *MongoDB) repository.ProductDiffsRepository {
	return &productDiffRepo{
		col: conn.productDiffCol,
	}
}

func (repo *productMetaInfoRepo) GetByProductID(brandKeyname string, productID string) (*domain.ProductMetaInfoDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var oldProduct domain.ProductMetaInfoDAO
	filter := bson.M{"productid": productID, "brand.keyname": brandKeyname}
	err := repo.col.FindOne(ctx, filter).Decode(&oldProduct)
	if err != nil {
		return nil, err
	}

	return &oldProduct, nil
}

func (repo *productMetaInfoRepo) Insert(pd *domain.ProductMetaInfoDAO) (*domain.ProductMetaInfoDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var oldProduct domain.ProductMetaInfoDAO
	_, err := repo.col.InsertOne(ctx, pd)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"productid": pd.ProductID, "brand.keyname": pd.Brand.KeyName}
	err = repo.col.FindOne(ctx, filter).Decode(&oldProduct)
	if err != nil {
		return nil, err
	}

	return &oldProduct, nil
}

func (repo *productMetaInfoRepo) Upsert(product *domain.ProductMetaInfoDAO) (*domain.ProductMetaInfoDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": product.ID}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &product}, opts); err != nil {
		return nil, err
	}

	var updatedProduct *domain.ProductMetaInfoDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&updatedProduct); err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func MongoProductMetaInfosRepo(conn *MongoDB) repository.ProductMetaInfoRepository {
	return &productMetaInfoRepo{
		col: conn.productMetaInfoCol,
	}
}
