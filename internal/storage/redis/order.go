package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/lessbutter/alloff-api/internal/core/repository"
)

type orderCountsRepo struct {
	client *redis.Client
}

func (repo *orderCountsRepo) Get(exhibitionID string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	val, err := repo.client.Get(ctx, exhibitionID).Result()
	switch {
	case err == redis.Nil:
		err := repo.client.Set(ctx, exhibitionID, 1, 0).Err()
		if err != nil {
			return 0, nil
		}
	case err != nil:
		err := fmt.Errorf("Get failed" + err.Error())
		return 0, err
	case val == "":
		err := fmt.Errorf("Get failed" + err.Error())
		return 0, err
	}

	counts, _ := strconv.Atoi(val)
	return counts, nil
}

func (repo *orderCountsRepo) Push(exhibitionID string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	val, err := repo.client.Get(ctx, exhibitionID).Result()
	switch {
	case err == redis.Nil:
		err := repo.client.Set(ctx, exhibitionID, 1, 0).Err()
		if err != nil {
			return 0, nil
		}
		return 1, nil
	case err != nil:
		err := fmt.Errorf("Get failed" + err.Error())
		return 0, err
	case val == "":
		err := fmt.Errorf("Get failed" + err.Error())
		return 0, err
	}
	counts, _ := strconv.Atoi(val)
	counts += 1

	err = repo.client.Set(ctx, exhibitionID, counts, 0).Err()
	if err != nil {
		return counts, err
	}
	return counts, nil
}

func (repo *orderCountsRepo) Cancel(exhibitionID string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	val, err := repo.client.Get(ctx, exhibitionID).Result()
	switch {
	case err == redis.Nil:
		err := repo.client.Set(ctx, exhibitionID, 1, 0).Err()
		if err != nil {
			return 0, nil
		}
		return 1, nil
	case err != nil:
		err := fmt.Errorf("Get failed" + err.Error())
		return 0, err
	case val == "":
		err := fmt.Errorf("Get failed" + err.Error())
		return 0, err
	}
	counts, _ := strconv.Atoi(val)
	counts -= 1

	err = repo.client.Set(ctx, exhibitionID, counts, 0).Err()
	if err != nil {
		return counts, err
	}
	return counts, nil
}

func RedisOrderRepo(conn *RedisDB) repository.OrderCountsRepository {
	return &orderCountsRepo{
		client: conn.client,
	}
}
