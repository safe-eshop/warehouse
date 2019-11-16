package repository

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"warehouse/app/domain/model"
	model3 "warehouse/app/infrastructure/model"
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

func getRedisKeys(ids []string) []string {
	res := make([]string, len(ids))
	for i, id := range ids {
		res[i] = getRedisKey(id)
	}
	return res
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

func (r *redisWarehouseStateRepository) FindByIds(ids []string) ([]*model.WarehouseState, error) {
	redisKeys := getRedisKeys(ids)
	result := make([]*model.WarehouseState, len(ids))
	res, err := r.redis.MGet(redisKeys...).Result()
	if err == redis.Nil {
		for i, id := range ids {
			result[i] = model.Zero(id)
		}
	} else if err != nil {
		return nil, err
	} else {
		for i, r := range res {
			if r == nil {
				result[i] = model.Zero(ids[i])
			} else {
				var redisModel model3.RedisWarehouseState
				err = json.Unmarshal([]byte(fmt.Sprintf("%v", r)), &redisModel)
				if err != nil {
					return nil, err
				}
				result[i] = redisModel.ToWarehouseState()
			}
		}
	}

	return result, err
}
