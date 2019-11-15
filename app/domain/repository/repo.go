package repository

import "rossmann/app/domain/model"

type WarehouseStateRepository interface {
	FindById(id string) (*model.WarehouseState, error)
}
