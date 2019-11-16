package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rossmann/app/domain/model"
)
import "github.com/go-redis/redis/v7"

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func main() {
	r := gin.Default()

	r.GET("/ping/:quantity", func(c *gin.Context) {
		stock := model.NewWarehouseState("hagjhdsag", 2, 4)
		c.JSON(200, stock)
	})
	_ = r.Run("0.0.0.0:9000") // listen and serve on 0.0.0.0:8080
}
