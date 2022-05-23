package cmd

import (
	"github.com/lessbutter/alloff-api/internal/storage/redis"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/storage/mongo"
	"github.com/lessbutter/alloff-api/internal/storage/postgres"
)

func SetBaseConfig() {
	config.InitViper()

	conn := mongo.NewMongoDB()
	conn.RegisterRepos()
	pgconn := postgres.NewPostgresDB()
	pgconn.RegisterRepos()
	redisConn := redis.NewRedis()
	redisConn.RegisterRepos()

	config.InitLogger()
	config.InitSlack()
	config.InitIamPort()
	config.InitNotification()
	config.InitSentry()
	config.InitAmplitude()
}
