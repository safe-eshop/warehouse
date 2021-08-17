package service

import (
	"context"
	"errors"
	"testing"
	"warehouse/app/domain/model"
	"warehouse/app/domain/repository"

	"github.com/stretchr/testify/assert"
)

type ErrorRepo struct {
}

func (r ErrorRepo) Count(context context.Context) (int64, error) {
	return 0, errors.New("Test")
}

func (r ErrorRepo) InsertMany(context context.Context, states []*model.WarehouseState) error {
	return errors.New("Test")
}

func (r ErrorRepo) FindById(context context.Context, id model.CatalogItemId) (*model.WarehouseState, error) {
	return nil, repository.ErrNotFound
}

func (r ErrorRepo) FindByIds(context context.Context, ids []model.CatalogItemId) ([]*model.WarehouseState, error) {
	return nil, repository.ErrNotFound
}

func TestFindAvailableQuantityWhenRepositoryReturnError(t *testing.T) {
	ctx := context.TODO()
	service := WarehouseStateService{ErrorRepo{}}
	_, err := service.GetAvailableCatalogItemQuantity(ctx, 1)
	assert.Error(t, err)
	assert.Equal(t, err, repository.ErrNotFound, "")
}

func TestFindAvailableQuantitiesWhenRepositoryReturnError(t *testing.T) {
	ctx := context.TODO()
	service := WarehouseStateService{ErrorRepo{}}
	_, err := service.GetAvailableCatalogItemsQuantity(ctx, []int{1, 2})
	assert.Error(t, err)
	assert.Equal(t, err, repository.ErrNotFound, "")
}

type OkRepo struct {
}

func (r OkRepo) Count(context context.Context) (int64, error) {
	return 0, nil
}

func (r OkRepo) InsertMany(context context.Context, states []*model.WarehouseState) error {
	return nil
}

func (r OkRepo) FindById(context context.Context, id model.CatalogItemId) (*model.WarehouseState, error) {
	return model.NewWarehouseState(id, 10, 1), nil
}

func (r OkRepo) FindByIds(context context.Context, ids []model.CatalogItemId) ([]*model.WarehouseState, error) {
	res := make([]*model.WarehouseState, len(ids))
	for i, id := range ids {
		res[i] = model.NewWarehouseState(id, 10, 1)
	}
	return res, nil
}

func TestFindAvailableQuantity(t *testing.T) {
	service := WarehouseStateService{OkRepo{}}
	ctx := context.TODO()
	subject, err := service.GetAvailableCatalogItemQuantity(ctx, 1)
	assert.Nil(t, err)
	assert.Equal(t, 9, subject.AvailableQuantity)
	assert.Equal(t, 1, subject.Id)
}

func TestFindAvailableQuantities(t *testing.T) {
	service := WarehouseStateService{OkRepo{}}
	ctx := context.TODO()
	subject, err := service.GetAvailableCatalogItemsQuantity(ctx, []int{1, 2})
	assert.Nil(t, err)
	for _, res := range subject {
		assert.Equal(t, 9, res.AvailableQuantity)
	}
}
