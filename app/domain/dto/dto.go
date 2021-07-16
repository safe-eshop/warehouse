package dto

import "warehouse/app/domain/model"

type AvailableQuantity struct {
	Id                model.ProductId `json:"id,omitempty"`
	AvailableQuantity int             `json:"availableQuantity,omitempty"`
}

func NewAvailableQuantity(id model.ProductId, availableQuantity int) *AvailableQuantity {
	return &AvailableQuantity{Id: id, AvailableQuantity: availableQuantity}
}

func FromWarehouseState(state model.WarehouseState) *AvailableQuantity {
	return &AvailableQuantity{
		Id:                state.GetID(),
		AvailableQuantity: state.GetAvailableQuantity(),
	}
}

func FromWarehouseStates(states []*model.WarehouseState) []*AvailableQuantity {
	res := make([]*AvailableQuantity, len(states))
	for i, state := range states {
		res[i] = FromWarehouseState(*state)
	}
	return res
}
