package cmd

import (
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
	config.InitSlack(conf)
	config.InitIamPort(conf)
	config.InitNotification(conf)
	config.InitSentry(conf)

	return conf
}
