package main

import (
	"github.com/gin-gonic/gin"
	"rossmann/domain"
)

func main() {
	r := gin.Default()

	r.GET("/ping/:quantity", func(c *gin.Context) {
		stock := domain.NewWarehouseState("hagjhdsag", 2, 4)
		c.JSON(200, stock)
	})
	_ = r.Run("0.0.0.0:9000") // listen and serve on 0.0.0.0:8080
}
