package connection

import (
	"context"
	"log"
	"warehouse/app/common"

	"github.com/go-redis/redis/v8"
)

func getRedisConnectionString() string {
	return common.GetOsEnvOrDefault("REDIS_CONNECTION", "localhost:6379")
}

func NewRedisClient(context context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     getRedisConnectionString(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping(context).Result()
	if err != nil {
		log.Fatal(err)
	}
	return client
}
