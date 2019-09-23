package repository

import (
	"fmt"
	"rossmann/model"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

func CreateSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*model.Product)(nil)} {
		fmt.Println("Elo")
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			println(err.Error())
			return err
		}
	}
	return nil
}
