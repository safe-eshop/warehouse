package repository

import "rossmann/app/domain/model"

type UserRepository interface {
	FindById(id string) (*model.WarehouseState, error)
}
