package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	brandCol             *mongo.Collection
	productCol           *mongo.Collection
	productMetaInfoCol   *mongo.Collection
	productGroupCol      *mongo.Collection
	crawlSourceCol       *mongo.Collection
	crawlRecordCol       *mongo.Collection
	categoryCol          *mongo.Collection
	alloffCategoryCol    *mongo.Collection
	classifyRuleCol      *mongo.Collection
	userCol              *mongo.Collection
	deviceCol            *mongo.Collection
	notificationCol      *mongo.Collection
	alimtalkCol          *mongo.Collection
	likeBrandsCol        *mongo.Collection
	exhibitionCol        *mongo.Collection
	alloffSizeCol        *mongo.Collection
	cartCol              *mongo.Collection
	sizeMappingPolicyCol *mongo.Collection
}

func NewMongoDB() *MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	viper.Get("MONGO_DB_NAME")
	mongoClient, err := makeMongoClient(ctx)
	checkErr(err, "Connection in mongodb")
	checkErr(mongoClient.Ping(ctx, readpref.Primary()), "Ping error in mongoconnect")
	db := mongoClient.Database(viper.GetString("MONGO_DB_NAME"))

	return &MongoDB{
		brandCol:             db.Collection("brands"),
		productCol:           db.Collection("products_v1"),
		productMetaInfoCol:   db.Collection("product_infos_v1"),
		productGroupCol:      db.Collection("productgroups_v1"),
		crawlSourceCol:       db.Collection("sources"),
		crawlRecordCol:       db.Collection("crawling_records"),
		categoryCol:          db.Collection("categories_v1"),
		alloffCategoryCol:    db.Collection("alloffcategories_v1"),
		classifyRuleCol:      db.Collection("classifier_v1"),
		userCol:              db.Collection("users"),
		deviceCol:            db.Collection("devices"),
		notificationCol:      db.Collection("notifications_v1"),
		alimtalkCol:          db.Collection("alimtalks"),
		likeBrandsCol:        db.Collection("likes_brands"),
		exhibitionCol:        db.Collection("exhibitions_v1"),
		alloffSizeCol:        db.Collection("alloff_sizes"),
		cartCol:              db.Collection("carts"),
		sizeMappingPolicyCol: db.Collection("size_mapping_policy"),
	}
}

func (conn *MongoDB) RegisterRepos() {
	ioc.Repo.Brands = MongoBrandsRepo(conn)
	ioc.Repo.Products = MongoProductsRepo(conn)
	ioc.Repo.ProductMetaInfos = MongoProductMetaInfosRepo(conn)
	ioc.Repo.CrawlSources = MongoCrawlSourcesRepo(conn)
	ioc.Repo.CrawlRecords = MongoCrawlRecordRepo(conn)
	ioc.Repo.Categories = MongoCategoriesRepo(conn)
	ioc.Repo.AlloffCategories = MongoAlloffCategoriesRepo(conn)
	ioc.Repo.ClassifyRules = MongoClassifyRulesRepo(conn)
	ioc.Repo.LikeBrands = MongoBrandLikesRepo(conn)
	ioc.Repo.Users = MongoUsersRepo(conn)
	ioc.Repo.Devices = MongoDevicesRepo(conn)
	ioc.Repo.Alimtalks = MongoAlimtalksRepo(conn)
	ioc.Repo.ProductGroups = MongoProductGroupsRepo(conn)
	ioc.Repo.Exhibitions = MongoExhibitionsRepo(conn)
	ioc.Repo.Notifications = MongoNotificationsRepo(conn)
	ioc.Repo.AlloffSizes = MongoAlloffSizeRepo(conn)
	ioc.Repo.Carts = MongoCartsRepo(conn)
}

func makeMongoClient(ctx context.Context) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + viper.GetString("MONGO_URL") + "/" + viper.GetString("MONGO_DB_NAME") + "?&connect=direct&replicaSet=rs0&readPreference=secondaryPreferred&retryWrites=false").SetAuth(options.Credential{
		Username: viper.GetString("MONGO_USERNAME"),
		Password: viper.GetString("MONGO_PASSWORD"),
	})
	mongoClient, err := mongo.Connect(ctx, clientOptions)

	return mongoClient, err
}

func checkErr(err error, location string) {
	if err != nil {
		fmt.Println("Error occured: " + location)
		fmt.Println("Message: " + err.Error())
	}
}
