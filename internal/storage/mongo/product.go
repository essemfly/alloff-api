package mongo

import (
	"context"
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
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

type productLikeRepo struct {
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

func (repo *productRepo) List(offset, limit int, filter, sortingOptions interface{}) ([]*domain.ProductDAO, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSort(sortingOptions)
	options.SetSkip(int64(offset))
	options.SetLimit(int64(limit))

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

	oid, err := repo.col.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}

	var oldProduct domain.ProductDAO
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

func (repo *productRepo) ListDistinctBrands(alloffCategoryID string) ([]*domain.BrandDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{"removed": false, "alloffcategories.done": true}

	alloffCat, err := ioc.Repo.AlloffCategories.Get(alloffCategoryID)
	if err != nil {
		return nil, err
	}

	if alloffCat.Level == 1 {
		filter["alloffcategories.first._id"] = alloffCat.ID
	} else if alloffCat.Level == 2 {
		filter["alloffcategories.second._id"] = alloffCat.ID
	}

	rows, err := repo.col.Distinct(ctx, "brand", filter)
	if err != nil {
		return nil, err
	}

	brands := []*domain.BrandDAO{}

	for _, row := range rows {
		var brand *domain.BrandDAO
		data, err := bson.Marshal(row)
		if err != nil {
			log.Println("Err in marshaling")
		}

		err = bson.Unmarshal(data, &brand)
		if err != nil {
			log.Println("Err in unmarshaling")
		}
		brands = append(brands, brand)
	}

	return brands, nil
}

func (repo *productRepo) CountNewProducts(brandModules []string) int {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"score.isnewlycrawled": true, "removed": false, "productinfo.source.crawlmodulename": bson.M{"$in": brandModules}}
	newProducts, err := repo.col.CountDocuments(ctx, filter)
	if err != nil {
		log.Println("Find num of new crawled products error", err)
	}

	return int(newProducts)
}

func (repo *productRepo) MakeOutdateProducts(brandModules []string, lastUpdatedDate time.Time) int {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	outProducts, err := repo.col.UpdateMany(
		ctx,
		bson.M{
			"removed":                            false,
			"productinfo.source.crawlmodulename": bson.M{"$in": brandModules},
			"updated": bson.M{
				"$lte": primitive.NewDateTimeFromTime(lastUpdatedDate),
			},
		}, bson.M{"$set": bson.M{"removed": true}})
	if err != nil {
		log.Println("Find num of outdated products error", err)
	}

	return int(outProducts.ModifiedCount)
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

func (repo *productLikeRepo) Like(userID, productID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	product, err := ioc.Repo.Products.Get(productID)
	if err != nil {
		return false, err
	}

	var likes *domain.LikeProductDAO
	if err := repo.col.FindOne(ctx, bson.M{"userid": userID, "productid": productID, "removed": false}).Decode(&likes); err != nil {
		repo.col.InsertOne(
			ctx,
			bson.M{"userid": userID, "created": time.Now(), "updated": time.Now(), "productid": productID, "product": product, "removed": false, "ispushed": false, "lastprice": product.DiscountedPrice},
		)
		return true, nil
	}

	repo.col.FindOneAndUpdate(ctx, bson.M{"userid": userID, "productid": productID, "removed": false}, bson.M{"$set": bson.M{"removed": true, "updated": time.Now()}})

	return false, nil
}

func (repo *productLikeRepo) List(userID string) ([]*domain.LikeProductDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{"userid": userID, "removed": false}
	cursor, err := repo.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var likes []*domain.LikeProductDAO
	err = cursor.All(ctx, &likes)
	if err != nil {
		return nil, err
	}

	return likes, nil
}

func MongoProductLikesRepo(conn *MongoDB) repository.LikeProductsRepository {
	return &productLikeRepo{
		col: conn.likeProductsCol,
	}
}
