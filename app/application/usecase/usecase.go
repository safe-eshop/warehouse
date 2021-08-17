package usecase

import (
	"context"
	"warehouse/app/domain/dto"
	"warehouse/app/domain/model"
	"warehouse/app/domain/repository"
	"warehouse/app/domain/service"

	"github.com/brianvoe/gofakeit"
)

type WarehouseStateUseCase interface {
	GetAvailableCatalogItemQuantity(id model.CatalogItemId) (*dto.AvailableQuantity, error)
	GetAvailableCatalogItemsQuantity(ids []model.CatalogItemId) ([]*dto.AvailableQuantity, error)
	SeedDatabase() error
}

type warehouseStateUseCase struct {
	repo    repository.WarehouseStateRepository
	service *service.WarehouseStateService
}

func NewWarehouseStateUseCaseUseCase(repo repository.WarehouseStateRepository, service *service.WarehouseStateService) *warehouseStateUseCase {
	return &warehouseStateUseCase{
		repo:    repo,
		service: service,
	}
}

func (u *warehouseStateUseCase) GetAvailableCatalogItemQuantity(ctx context.Context, id model.CatalogItemId) (*dto.AvailableQuantity, error) {
	return u.service.GetAvailableCatalogItemQuantity(ctx, id)
}

func (u *warehouseStateUseCase) GetAvailableCatalogItemsQuantity(ctx context.Context, ids []model.CatalogItemId) ([]*dto.AvailableQuantity, error) {
	return u.service.GetAvailableCatalogItemsQuantity(ctx, ids)
}

func (u *warehouseStateUseCase) SeedDatabase(ctx context.Context) error {
	q, err := u.repo.Count(ctx)
	if err != nil {
		return err
	}
	if q > 0 {
		return nil
	}

	count := 100
	result := make([]*model.WarehouseState, count)
	gofakeit.Seed(0)
	for i := 0; i < count; i++ {
		result[i] = model.NewWarehouseState(i, gofakeit.Number(0, 1000), gofakeit.Number(0, 1000))
	}
	err = u.repo.InsertMany(ctx, result)
	return err
}
