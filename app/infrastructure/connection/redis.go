package connection

import (
	"context"
	"warehouse/app/common"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

func getRedisConnectionString() string {
	return common.GetOsEnvOrDefault("REDIS_CONNECTION", "db:6379")
}

func NewRedisClient(context context.Context) *redis.Client {
	redis_conn := getRedisConnectionString()
	client := redis.NewClient(&redis.Options{
		Addr:     redis_conn,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping(context).Result()
	if err != nil {
		log.WithField("redisConnection", redis_conn).WithContext(context).WithError(err).Fatal("Redis connection error")
	}
	return client
}
