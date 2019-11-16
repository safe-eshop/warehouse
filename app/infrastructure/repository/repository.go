package repository

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"rossmann/app/domain/model"
	model3 "rossmann/app/infrastructure/model"
)

type redisWarehouseStateRepository struct {
	redis *redis.Client
}

func NewWarehouseStateRepository(redis *redis.Client) *redisWarehouseStateRepository {
	return &redisWarehouseStateRepository{redis: redis}
}

func getRedisKey(id string) string {
	return fmt.Sprintf("{warehouse/%s}", id)
}

func (r *redisWarehouseStateRepository) FindById(id string) (*model.WarehouseState, error) {
	redisKey := getRedisKey(id)
	res, err := r.redis.Get(redisKey).Result()
	if err == redis.Nil {
		return model.Zero(id), nil
	} else if err != nil {
		return nil, err
	} else {
		var redisModel model3.RedisWarehouseState
		err = json.Unmarshal([]byte(res), &redisModel)
		if err != nil {
			return nil, err
		}
		return redisModel.ToWarehouseState(), nil
	}
}
