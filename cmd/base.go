package cmd

import (
	"github.com/lessbutter/alloff-api/internal/storage/elasticsearch"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/storage/mongo"
	"github.com/lessbutter/alloff-api/internal/storage/postgres"
	"github.com/lessbutter/alloff-api/internal/storage/redis"
)

func SetBaseConfig() {
	config.InitViper()

	conn := mongo.NewMongoDB()
	conn.RegisterRepos()
	pgconn := postgres.NewPostgresDB()
	pgconn.RegisterRepos()
	redisConn := redis.NewRedis()
	redisConn.RegisterRepos()
	esConn := elasticsearch.NewElasticSearch()
	esConn.RegisterRepos()

	config.InitLogger()
	config.InitSlack()
	config.InitIamPort()
	config.InitNotification()
	config.InitSentry()
	config.InitAmplitude()

}
