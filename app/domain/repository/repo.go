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
	FindById(context context.Context, id model.CatalogItemId) (*model.WarehouseState, error)
	FindByIds(context context.Context, ids []model.CatalogItemId) ([]*model.WarehouseState, error)
	InsertMany(context context.Context, states []*model.WarehouseState) error
	Insert(context context.Context, states *model.WarehouseState) error
	Count(context context.Context) (int64, error)
}
