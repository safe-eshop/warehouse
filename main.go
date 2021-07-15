package main

import (
	"context"
	"log"
	"net/http"
	"warehouse/app/application/usecase"
	"warehouse/app/common"
	"warehouse/app/domain/service"
	"warehouse/app/infrastructure/connection"
	"warehouse/app/infrastructure/repository"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func connect(connection string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(connection)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func getAppPrefix() string {
	path := common.GetOsEnvOrDefault("PATH_BASE", "/")
	return path + "states"
}

func main() {
	ctx := context.TODO()
	r := gin.Default()
	router := r.Group(getAppPrefix())
	stateRepository := repository.NewWarehouseStateRepository(connection.NewRedisClient(ctx))
	stateService := service.NewWarehouseStateService(stateRepository)
	useCase := usecase.NewWarehouseStateUseCaseUseCase(stateRepository, stateService)
	_ = useCase.SeedDatabase(ctx)
	router.GET("/:catalogItemId", func(c *gin.Context) {
		catalogItemId := c.Param("catalogItemId")
		stock, err := useCase.GetAvailableCatalogItemQuantity(c.Request.Context(), catalogItemId)
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "unknown error")
		} else {
			c.JSON(200, stock)
		}

	})

	router.GET("/", func(c *gin.Context) {
		catalogItemIds := c.QueryArray("ids")
		stocks, err := useCase.GetAvailableCatalogItemsQuantity(c.Request.Context(), catalogItemIds)
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "unknown error")
		} else {
			c.JSON(200, stocks)
		}

	})
	_ = r.Run("0.0.0.0:9000") // listen and serve on 0.0.0.0:8080
}
