package main

import (
	"fmt"
	"rossmann/products"
	"rossmann/repository"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
)

func main() {
	r := gin.Default()
	r.GET("/ping/:quantity", func(c *gin.Context) {
		quantity := c.Param("quantity")
		i, err := strconv.ParseInt(quantity, 10, 64)
		if err != nil {
			c.Error(err)
		}
		fmt.Println(quantity)
		p, err := products.GetRossmannPricesFromApi(int(i))
		if err != nil {
			c.Error(err)
		}
		c.JSON(200, p)
	})
	var db *pg.DB = pg.Connect(&pg.Options{
		User:     "user",
		Password: "pass",
		Addr:     "db:5432",

		Database: "data",
	})
	repository.CreateSchema(db)
	r.Run("0.0.0.0:9000") // listen and serve on 0.0.0.0:8080
}
