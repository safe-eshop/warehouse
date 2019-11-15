package repository

import (
	"errors"
	"rossmann/app/domain/model"
)

type QueryError struct {
	Id  string
	Err error
}

var ErrNotFound = errors.New("WarehouseState not found")

type WarehouseStateRepository interface {
	FindById(id string) (*model.WarehouseState, error)
}
