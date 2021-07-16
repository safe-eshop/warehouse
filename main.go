package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"warehouse/app/application/usecase"
	"warehouse/app/common"
	"warehouse/app/domain/service"
	"warehouse/app/infrastructure/connection"
	"warehouse/app/infrastructure/repository"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func parseInt(idsStr []string) []int {
	res := make([]int, len(idsStr))
	for i, v := range idsStr {
		id, _ := strconv.Atoi(v)
		res[i] = id
	}
	return res
}

func connect(connection string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(connection)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func getAppPrefix() string {
	path := common.GetOsEnvOrDefault("PATH_BASE", "/")
	return path
}

func main() {
	ctx := context.TODO()
	r := gin.Default()
	router := r.Group(getAppPrefix())
	stateRepository := repository.NewWarehouseStateRepository(connection.NewRedisClient(ctx))
	stateService := service.NewWarehouseStateService(stateRepository)
	useCase := usecase.NewWarehouseStateUseCaseUseCase(stateRepository, stateService)
	_ = useCase.SeedDatabase(ctx)
	router.GET("/products/:catalogItemId", func(c *gin.Context) {
		catalogItemId := c.Param("catalogItemId")
		id, err := strconv.Atoi(catalogItemId)

		if err != nil {
			c.Error(err)
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		stock, err := useCase.GetAvailableCatalogItemQuantity(c.Request.Context(), id)
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "unknown error")
		} else {
			c.JSON(200, stock)
		}

	})

	router.GET("/products", func(c *gin.Context) {
		catalogItemIds := c.QueryArray("ids")
		ids := parseInt(catalogItemIds)
		stocks, err := useCase.GetAvailableCatalogItemsQuantity(c.Request.Context(), ids)
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "unknown error")
		} else {
			c.JSON(200, stocks)
		}

	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}
