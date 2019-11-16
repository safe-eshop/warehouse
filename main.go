package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"warehouse/app/application/usecase"
	"warehouse/app/domain/service"
	"warehouse/app/infrastructure/repository"
)
import "github.com/go-redis/redis/v7"

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	return client

}

func main() {
	r := gin.Default()
	stateRepository := repository.NewWarehouseStateRepository(NewClient())
	stateService := service.NewWarehouseStateService(stateRepository)
	useCase := usecase.NewWarehouseStateUseCaseUseCase(stateRepository, stateService)
	_ = useCase.SeedDatabase()
	r.GET("/state/:catalogItemId", func(c *gin.Context) {
		catalogItemId := c.Param("catalogItemId")
		stock, err := useCase.GetAvailableCatalogItemQuantity(catalogItemId)
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "unknown error")
		} else {
			c.JSON(200, stock)
		}

	})

	r.GET("/state", func(c *gin.Context) {
		catalogItemIds := c.QueryArray("ids")
		stocks, err := useCase.GetAvailableCatalogItemsQuantity(catalogItemIds)
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "unknown error")
		} else {
			c.JSON(200, stocks)
		}

	})
	_ = r.Run("0.0.0.0:9000") // listen and serve on 0.0.0.0:8080
}
