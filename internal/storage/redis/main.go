package redis

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
)

type RedisDB struct {
	client *redis.Client
}

func NewRedis(conf config.Configuration) *RedisDB {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.REDIS_URL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Println("err in redis connection test", err)
	}

	return &RedisDB{
		client: client,
	}
}

func (conn *RedisDB) RegisterRepos() {
	ioc.Repo.OrderCounts = RedisOrderRepo(conn)
}
