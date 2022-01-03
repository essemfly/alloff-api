package mongo

import (
	"context"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type brandsRepo struct {
	col *mongo.Collection
}

type brandLikeRepo struct {
	col *mongo.Collection
}

func (repo *brandsRepo) Get(ID string) (*domain.BrandDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	brandObjectID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": brandObjectID}

	var brand *domain.BrandDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&brand); err != nil {
		return nil, err
	}

	return brand, nil
}

func (repo *brandsRepo) GetByKeyname(keyname string) (*domain.BrandDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{"keyname": keyname}

	var brand *domain.BrandDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&brand); err != nil {
		return nil, err
	}

	return brand, nil
}

func (repo *brandsRepo) List(limit, offset int, filter, sortingOptions interface{}) ([]*domain.BrandDAO, int, error) {
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

	var brands []*domain.BrandDAO
	err = cursor.All(ctx, &brands)
	if err != nil {
		return nil, 0, err
	}

	return brands, int(totalCount), nil
}

func (repo *brandsRepo) Upsert(brand *domain.BrandDAO) (*domain.BrandDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"keyname": &brand.KeyName}
	brand.ID = ""

	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &brand}, opts); err != nil {
		return nil, err
	}

	var updatedBrand *domain.BrandDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&updatedBrand); err != nil {
		return nil, err
	}

	return updatedBrand, nil
}

func MongoBrandsRepo(conn *MongoDB) repository.BrandsRepository {
	return &brandsRepo{
		col: conn.brandCol,
	}
}

func (repo *brandLikeRepo) Like(userID, brandID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	brand, err := ioc.Repo.Brands.Get(brandID)
	if err != nil {
		return false, err
	}

	var likes *domain.LikeBrandDAO
	if err := repo.col.FindOne(ctx, bson.M{"userid": userID}).Decode(&likes); err != nil {
		if _, err := repo.col.InsertOne(ctx,
			bson.M{
				"userid":  userID,
				"brands":  []*domain.BrandDAO{brand},
				"created": time.Now(),
				"updated": time.Now(),
			}); err != nil {
			return false, err
		}
		return true, nil
	}

	brandIndexInLikes := -1
	for i, brand := range likes.Brands {
		if brand.ID == brandID {
			brandIndexInLikes = i
			break
		}
	}

	result := true
	if brandIndexInLikes > -1 {
		result = false
		likes.Brands = removeBrand(likes.Brands, brandIndexInLikes)
	} else {
		likes.Brands = append(likes.Brands, brand)
	}
	likes.Updated = time.Now()
	repo.col.FindOneAndUpdate(ctx, bson.M{"userid": userID}, bson.M{"$set": likes})

	return result, nil
}

func removeBrand(s []*domain.BrandDAO, i int) []*domain.BrandDAO {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func MongoBrandLikesRepo(conn *MongoDB) repository.LikeBrandsRepository {
	return &brandLikeRepo{
		col: conn.likeProductsCol,
	}
}
