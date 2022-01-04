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

type alloffCategoryRepo struct {
	col *mongo.Collection
}

type classifyRuleRepo struct {
	col *mongo.Collection
}

func (repo *alloffCategoryRepo) Get(ID string) (*domain.AlloffCategoryDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	catID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": catID}

	var cat *domain.AlloffCategoryDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&cat); err != nil {
		return nil, err
	}

	return cat, nil
}
func (repo *alloffCategoryRepo) GetByName(name string) (*domain.AlloffCategoryDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var cat *domain.AlloffCategoryDAO
	if err := repo.col.FindOne(ctx, bson.M{"name": name}).Decode(&cat); err != nil {
		return nil, err
	}

	return cat, nil
}

func (repo *alloffCategoryRepo) GetByKeyname(keyname string) (*domain.AlloffCategoryDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var cat *domain.AlloffCategoryDAO
	if err := repo.col.FindOne(ctx, bson.M{"keyname": keyname}).Decode(&cat); err != nil {
		return nil, err
	}

	return cat, nil
}

func (repo *alloffCategoryRepo) List(parentID *string) ([]*domain.AlloffCategoryDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	sortingOptions := bson.D{{Key: "_id", Value: 1}}
	options := options.Find()
	options.SetSort(sortingOptions)

	filter := bson.M{"level": 1}

	if parentID != nil {
		parentObjId, _ := primitive.ObjectIDFromHex(*parentID)
		filter = bson.M{"parentid": parentObjId}
	}

	cursor, err := repo.col.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	var cats []*domain.AlloffCategoryDAO
	err = cursor.All(ctx, &cats)
	if err != nil {
		return nil, err
	}

	return cats, nil
}

func (repo *alloffCategoryRepo) Upsert(cat *domain.AlloffCategoryDAO) (*domain.AlloffCategoryDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"keyname": cat.KeyName}

	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &cat}, opts); err != nil {
		return nil, err
	}

	var updatedCat *domain.AlloffCategoryDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&updatedCat); err != nil {
		return nil, err
	}

	return updatedCat, nil
}

func MongoAlloffCategoriesRepo(conn *MongoDB) repository.AlloffCategoriesRepository {
	return &alloffCategoryRepo{
		col: conn.alloffCategoryCol,
	}
}

func (repo *classifyRuleRepo) Upsert(rule *domain.ClassifierDAO) (*domain.ClassifierDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	opts := options.Update().SetUpsert(true)

	filter := bson.M{"brandkeyname": rule.BrandKeyname, "categoryname": rule.CategoryName}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &rule}, opts); err != nil {
		return nil, err
	}

	var updatedRule *domain.ClassifierDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&updatedRule); err != nil {
		return nil, err
	}

	return updatedRule, nil
}

func (repo *classifyRuleRepo) GetByKeyname(brandKeyname, categoryKeyname string) (*domain.ClassifierDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	var rule *domain.ClassifierDAO
	if err := repo.col.FindOne(ctx, bson.M{"brandkeyname": brandKeyname, "categoryname": categoryKeyname}).Decode(&rule); err != nil {
		return nil, err
	}
	return rule, nil
}

func MongoClassifyRulesRepo(conn *MongoDB) repository.ClassifyRulesRepository {
	return &classifyRuleRepo{
		col: conn.classifyRuleCol,
	}
}
