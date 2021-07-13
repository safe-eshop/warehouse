package service

import (
	"warehouse/app/domain/dto"
	"warehouse/app/domain/repository"
)

type WarehouseStateService struct {
	repo repository.WarehouseStateRepository
}

func NewWarehouseStateService(repo repository.WarehouseStateRepository) *WarehouseStateService {
	return &WarehouseStateService{repo: repo}
}

func (s *WarehouseStateService) GetAvailableCatalogItemQuantity(id string) (*dto.AvailableQuantity, error) {
	state, err := s.repo.FindById(id)
	if err != nil {
		return dto.NewAvailableQuantity(id, 0), err
	}
	return dto.FromWarehouseState(*state), nil
}

func (s *WarehouseStateService) GetAvailableCatalogItemsQuantity(ids []string) ([]*dto.AvailableQuantity, error) {
	state, err := s.repo.FindByIds(ids)
	if err != nil {
		return nil, err
	}
	result := dto.FromWarehouseStates(state)
	return result, nil
}
