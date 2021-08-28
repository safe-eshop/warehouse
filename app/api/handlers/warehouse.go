package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"warehouse/app/application/usecase"

	"github.com/gin-gonic/gin"
)

type warehouseApi struct {
	router  *gin.RouterGroup
	usecase usecase.WarehouseStateUseCase
}

func NewWarehouseApi(router *gin.RouterGroup, usecase usecase.WarehouseStateUseCase) *warehouseApi {
	return &warehouseApi{router: router, usecase: usecase}
}

func parseInt(idsStr []string) []int {
	res := make([]int, len(idsStr))
	for i, v := range idsStr {
		id, _ := strconv.Atoi(v)
		res[i] = id
	}
	return res
}

func (api *warehouseApi) Start(ctx context.Context) {
	api.router.GET("/products/:catalogItemId", func(c *gin.Context) {
		catalogItemId := c.Param("catalogItemId")
		id, err := strconv.Atoi(catalogItemId)

		if err != nil {
			c.Error(err)
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		stock, err := api.usecase.GetAvailableCatalogItemQuantity(c.Request.Context(), id)
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "unknown error")
		} else {
			c.JSON(200, stock)
		}

	})

	api.router.GET("/products", func(c *gin.Context) {
		catalogItemIds := c.QueryArray("ids")
		ids := parseInt(catalogItemIds)
		stocks, err := api.usecase.GetAvailableCatalogItemsQuantity(c.Request.Context(), ids)
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "unknown error")
		} else {
			c.JSON(200, stocks)
		}

	})
}
