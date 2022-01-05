package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	brandCol             *mongo.Collection
	productCol           *mongo.Collection
	productMetaInfoCol   *mongo.Collection
	productDiffCol       *mongo.Collection
	productGroupCol      *mongo.Collection
	crawlSourceCol       *mongo.Collection
	crawlRecordCol       *mongo.Collection
	categoryCol          *mongo.Collection
	alloffCategoryCol    *mongo.Collection
	classifyRuleCol      *mongo.Collection
	featuredCol          *mongo.Collection
	homeitemCol          *mongo.Collection
	userCol              *mongo.Collection
	deviceCol            *mongo.Collection
	notificationCol      *mongo.Collection
	alimtalkCol          *mongo.Collection
	likeBrandsCol        *mongo.Collection
	likeProductsCol      *mongo.Collection
	alarmProductGroupCol *mongo.Collection
	exhibitionCol        *mongo.Collection
}

func NewMongoDB(conf config.Configuration) *MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	credential := options.Credential{
		Username: conf.MONGO_USERNAME,
		Password: conf.MONGO_PASSWORD,
	}
	clientOptions := options.Client().ApplyURI(conf.MONGO_URL).SetAuth(credential)

	mongoClient, err := mongo.Connect(ctx, clientOptions)

	checkErr(err, "Connection in mongodb")
	checkErr(mongoClient.Ping(ctx, readpref.Primary()), "Ping error in mongoconnect")
	db := mongoClient.Database(conf.MONGO_DB_NAME)

	return &MongoDB{
		brandCol:             db.Collection("brands"),
		productCol:           db.Collection("products"),
		productMetaInfoCol:   db.Collection("product_infos"),
		productDiffCol:       db.Collection("product_diffs"),
		productGroupCol:      db.Collection("productgroups"),
		crawlSourceCol:       db.Collection("sources"),
		crawlRecordCol:       db.Collection("crawling_records"),
		categoryCol:          db.Collection("categories"),
		alloffCategoryCol:    db.Collection("alloffcategories"),
		classifyRuleCol:      db.Collection("classifier"),
		featuredCol:          db.Collection("featured"),
		homeitemCol:          db.Collection("homeitems"),
		userCol:              db.Collection("users"),
		deviceCol:            db.Collection("devices"),
		notificationCol:      db.Collection("notifications"),
		alimtalkCol:          db.Collection("alimtalks"),
		likeBrandsCol:        db.Collection("likes"),
		likeProductsCol:      db.Collection("likes_products"),
		alarmProductGroupCol: db.Collection("alarm_productgroups"),
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
	ioc.Repo.ProductDiffs = MongoProductDiffsRepo(conn)
	ioc.Repo.Featureds = MongoFeaturedsRepo(conn)
	ioc.Repo.HomeItems = MongoHomeItemsRepo(conn)
	ioc.Repo.LikeBrands = MongoBrandLikesRepo(conn)
	ioc.Repo.LikeProducts = MongoProductLikesRepo(conn)
	ioc.Repo.Users = MongoUsersRepo(conn)
	ioc.Repo.Devices = MongoDevicesRepo(conn)
	ioc.Repo.Alimtalks = MongoAlimtalksRepo(conn)
	ioc.Repo.ProductGroups = MongoProductGroupsRepo(conn)
	ioc.Repo.Exhibitions = MongoExhibitionsRepo(conn)
}

func checkErr(err error, location string) {
	if err != nil {
		fmt.Println("Error occured: " + location)
		fmt.Println("Message: " + err.Error())
	}
}
