package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"warehouse/app/domain/model"
	model3 "warehouse/app/infrastructure/model"

	"github.com/go-redis/redis/v8"
)

type redisWarehouseStateRepository struct {
	redis *redis.Client
}

func NewWarehouseStateRepository(redis *redis.Client) *redisWarehouseStateRepository {
	return &redisWarehouseStateRepository{redis: redis}
}

func getRedisKey(id model.CatalogItemId) string {
	return fmt.Sprintf("{warehouse/%d}", id)
}

func getRedisKeys(ids []model.CatalogItemId) []string {
	res := make([]string, len(ids))
	for i, id := range ids {
		res[i] = getRedisKey(id)
	}
	return res
}

func (r *redisWarehouseStateRepository) Count(context context.Context) (int64, error) {
	result, err := r.redis.DBSize(context).Result()
	if err != nil {
		return 0, nil
	}
	return result, nil
}

func (r *redisWarehouseStateRepository) InsertMany(context context.Context, states []*model.WarehouseState) error {
	pipeline := r.redis.Pipeline()
	for _, state := range states {
		key := getRedisKey(state.CatalogItemId)
		redisVal, err := json.Marshal(model3.FromWarehouseState(*state))
		if err != nil {
			return err
		}

		err = pipeline.Set(context, key, redisVal, 24*time.Hour).Err()
		if err != nil {
			return err
		}
	}
	_, err := pipeline.Exec(context)
	if err != nil {
		return err
	}

	return nil
}

func (r *redisWarehouseStateRepository) Insert(context context.Context, state *model.WarehouseState) error {
	key := getRedisKey(state.CatalogItemId)
	redisVal, err := json.Marshal(model3.FromWarehouseState(*state))
	if err != nil {
		return err
	}
	err = r.redis.Set(context, key, redisVal, 24*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *redisWarehouseStateRepository) FindById(context context.Context, id model.CatalogItemId) (*model.WarehouseState, error) {
	redisKey := getRedisKey(id)
	res, err := r.redis.Get(context, redisKey).Result()
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

func (r *redisWarehouseStateRepository) FindByIds(context context.Context, ids []model.CatalogItemId) ([]*model.WarehouseState, error) {
	redisKeys := getRedisKeys(ids)
	result := make([]*model.WarehouseState, len(ids))
	res, err := r.redis.MGet(context, redisKeys...).Result()
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
