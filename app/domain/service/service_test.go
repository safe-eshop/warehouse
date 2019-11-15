package service

import (
	"github.com/stretchr/testify/assert"
	"rossmann/app/domain/model"
	"rossmann/app/domain/repository"
	"testing"
)

type ErrorRepo struct {
}

func (r ErrorRepo) FindById(id string) (*model.WarehouseState, error) {
	return nil, repository.ErrNotFound
}

func TestFindAvailableQuantityWhenRepositoryReturnError(t *testing.T) {
	service := WarehouseStateService{ErrorRepo{}}
	_, err := service.GetAvailableCatalogItemQuantity("")
	assert.Error(t, err)
	assert.Equal(t, err, repository.ErrNotFound, "")
}

type OkRepo struct {
}

func (r OkRepo) FindById(id string) (*model.WarehouseState, error) {
	return model.NewWarehouseState(id, 10, 1), nil
}

func TestFindAvailableQuantity(t *testing.T) {
	service := WarehouseStateService{OkRepo{}}
	subject, err := service.GetAvailableCatalogItemQuantity("test")
	assert.Nil(t, err)
	assert.Equal(t, subject, 9)
}
