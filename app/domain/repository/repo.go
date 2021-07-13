package repository

import (
	"errors"
	"warehouse/app/domain/model"
)

type QueryError struct {
	Id  string
	Err error
}

var ErrNotFound = errors.New("WarehouseState not found")

type WarehouseStateRepository interface {
	FindById(id string) (*model.WarehouseState, error)
	FindByIds(ids []string) ([]*model.WarehouseState, error)
	InsertMany(states []*model.WarehouseState) error
	Count() (int64, error)
}
