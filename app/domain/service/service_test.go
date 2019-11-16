package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/app/domain/model"
	"warehouse/app/domain/repository"
)

type ErrorRepo struct {
}

func (r ErrorRepo) FindById(id string) (*model.WarehouseState, error) {
	return nil, repository.ErrNotFound
}

func (r ErrorRepo) FindByIds(ids []string) ([]*model.WarehouseState, error) {
	return nil, repository.ErrNotFound
}

func TestFindAvailableQuantityWhenRepositoryReturnError(t *testing.T) {
	service := WarehouseStateService{ErrorRepo{}}
	_, err := service.GetAvailableCatalogItemQuantity("")
	assert.Error(t, err)
	assert.Equal(t, err, repository.ErrNotFound, "")
}

func TestFindAvailableQuantitiesWhenRepositoryReturnError(t *testing.T) {
	service := WarehouseStateService{ErrorRepo{}}
	_, err := service.GetAvailableCatalogItemsQuantity([]string{"a", "b"})
	assert.Error(t, err)
	assert.Equal(t, err, repository.ErrNotFound, "")
}

type OkRepo struct {
}

func (r OkRepo) FindById(id string) (*model.WarehouseState, error) {
	return model.NewWarehouseState(id, 10, 1), nil
}

func (r OkRepo) FindByIds(ids []string) ([]*model.WarehouseState, error) {
	res := make([]*model.WarehouseState, len(ids))
	for i, id := range ids {
		res[i] = model.NewWarehouseState(id, 10, 1)
	}
	return res, nil
}

func TestFindAvailableQuantity(t *testing.T) {
	service := WarehouseStateService{OkRepo{}}
	subject, err := service.GetAvailableCatalogItemQuantity("test")
	assert.Nil(t, err)
	assert.Equal(t, 9, subject.AvailableQuantity)
	assert.Equal(t, "test", subject.Id)
}

func TestFindAvailableQuantities(t *testing.T) {
	service := WarehouseStateService{OkRepo{}}
	subject, err := service.GetAvailableCatalogItemsQuantity([]string{"a", "b"})
	assert.Nil(t, err)
	for _, res := range subject {
		assert.Equal(t, 9, res.AvailableQuantity)
	}
}
