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

type MongoRepo struct {
	brandsCol   *mongo.Collection
	productsCol *mongo.Collection
	sourceCol   *mongo.Collection
	categoryCol *mongo.Collection
}

func NewMongoDB(conf config.Configuration) *MongoRepo {
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

	return &MongoRepo{
		brandsCol:   db.Collection("brands"),
		productsCol: db.Collection("products"),
		sourceCol:   db.Collection("sources"),
		categoryCol: db.Collection("categories"),
	}
}

func (conn *MongoRepo) RegisterRepos() {
	ioc.Repo.Brands = MongoBrandsRepo(conn)
}

func checkErr(err error, location string) {
	if err != nil {
		fmt.Println("Error occured: " + location)
		fmt.Println("Message: " + err.Error())
	}
}
