package repository

import (
	"context"
	"errors"
	"warehouse/app/domain/model"
)

type QueryError struct {
	Id  string
	Err error
}

var ErrNotFound = errors.New("WarehouseState not found")

type WarehouseStateRepository interface {
	FindById(context context.Context, id string) (*model.WarehouseState, error)
	FindByIds(context context.Context, ids []string) ([]*model.WarehouseState, error)
	InsertMany(context context.Context, states []*model.WarehouseState) error
	Count(context context.Context) (int64, error)
}
