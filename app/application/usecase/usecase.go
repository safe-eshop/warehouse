package usecase

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"warehouse/app/domain/dto"
	"warehouse/app/domain/model"
	"warehouse/app/domain/repository"
	"warehouse/app/domain/service"
)

type WarehouseStateUseCase interface {
	GetAvailableCatalogItemQuantity(id string) (*dto.AvailableQuantity, error)
	GetAvailableCatalogItemsQuantity(ids []string) ([]*dto.AvailableQuantity, error)
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

func (u *warehouseStateUseCase) GetAvailableCatalogItemQuantity(id string) (*dto.AvailableQuantity, error) {
	return u.service.GetAvailableCatalogItemQuantity(id)
}

func (u *warehouseStateUseCase) GetAvailableCatalogItemsQuantity(ids []string) ([]*dto.AvailableQuantity, error) {
	return u.service.GetAvailableCatalogItemsQuantity(ids)
}

func (u *warehouseStateUseCase) SeedDatabase() error {
	q, err := u.repo.Count()
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
		result[i] = model.NewWarehouseState(fmt.Sprint(gofakeit.UUID()), gofakeit.Number(0, 1000), gofakeit.Number(0, 1000))
	}
	err = u.repo.InsertMany(result)
	return err
}
