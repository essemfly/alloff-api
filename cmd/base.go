package cmd

import (
	"github.com/lessbutter/alloff-api/internal/storage/elasticsearch"
	"github.com/lessbutter/alloff-api/internal/storage/redis"
	"log"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/storage/mongo"
	"github.com/lessbutter/alloff-api/internal/storage/postgres"
)

func SetBaseConfig(Env string) config.Configuration {
	conf := config.GetConfiguration(Env)
	log.Println(conf)

	conn := mongo.NewMongoDB(conf)
	conn.RegisterRepos()
	pgconn := postgres.NewPostgresDB(conf)
	pgconn.RegisterRepos()
	redisConn := redis.NewRedis(conf)
	redisConn.RegisterRepos()
	esConn := elasticsearch.NewElasticSearch(conf)
	esConn.RegisterRepos()

	config.InitSlack(conf)
	config.InitIamPort(conf)
	config.InitNotification(conf)
	config.InitSentry(conf)
	config.InitAmplitude(conf)
	config.InitOmnious(conf)

	return conf
}
