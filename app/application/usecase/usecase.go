package usecase

import (
	"warehouse/app/domain/dto"
	"warehouse/app/domain/repository"
	"warehouse/app/domain/service"
)

type WarehouseStateUseCase interface {
	GetAvailableCatalogItemQuantity(id string) (int, error)
}

type warehouseStateUseCaseUsecase struct {
	repo    repository.WarehouseStateRepository
	service *service.WarehouseStateService
}

func NewWarehouseStateUseCaseUseCase(repo repository.WarehouseStateRepository, service *service.WarehouseStateService) *warehouseStateUseCaseUsecase {
	return &warehouseStateUseCaseUsecase{
		repo:    repo,
		service: service,
	}
}

func (u *warehouseStateUseCaseUsecase) GetAvailableCatalogItemQuantity(id string) (*dto.AvailableQuantity, error) {
	return u.service.GetAvailableCatalogItemQuantity(id)
}

func (u *warehouseStateUseCaseUsecase) GetAvailableCatalogItemsQuantity(ids []string) ([]*dto.AvailableQuantity, error) {
	return u.service.GetAvailableCatalogItemsQuantity(ids)
}
