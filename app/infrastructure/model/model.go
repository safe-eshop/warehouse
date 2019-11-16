package model

import "warehouse/app/domain/model"

type RedisWarehouseState struct {
	CatalogItemId string `json:"catalogItemId,omitempty"`
	ShopQuantity  int    `json:"quantity,omitempty"`
	Reservation   int    `json:"reservation,omitempty"`
}

func NewWarehouseState(id string, quantity, reservation int) *RedisWarehouseState {
	return &RedisWarehouseState{
		CatalogItemId: id,
		ShopQuantity:  quantity,
		Reservation:   reservation,
	}
}

func (w *RedisWarehouseState) ToWarehouseState() *model.WarehouseState {
	return &model.WarehouseState{
		CatalogItemId: w.CatalogItemId,
		ShopQuantity:  w.ShopQuantity,
		Reservation:   w.Reservation,
	}
}

func FromWarehouseState(model model.WarehouseState) *RedisWarehouseState {
	return NewWarehouseState(model.GetID(), model.ShopQuantity, model.Reservation)
}
