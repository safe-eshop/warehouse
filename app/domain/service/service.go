package service

import "warehouse/app/domain/repository"

type WarehouseStateService struct {
	repo repository.WarehouseStateRepository
}

func NewWarehouseStateService(repo repository.WarehouseStateRepository) *WarehouseStateService {
	return &WarehouseStateService{repo: repo}
}

func (s *WarehouseStateService) GetAvailableCatalogItemQuantity(id string) (int, error) {
	state, err := s.repo.FindById(id)
	if err != nil {
		return 0, err
	}
	q := state.GetAvailableQuantity()
	return q, nil
}
