package model

type CatalogItemId = int
type WarehouseState struct {
	CatalogItemId CatalogItemId `json:"catalogItemId,omitempty"`
	ShopQuantity  int           `json:"quantity,omitempty"`
	Reservation   int           `json:"reservation,omitempty"`
}

func NewWarehouseState(id CatalogItemId, quantity, reservation int) *WarehouseState {
	return &WarehouseState{
		CatalogItemId: id,
		ShopQuantity:  quantity,
		Reservation:   reservation,
	}
}

func Zero(id CatalogItemId) *WarehouseState {
	return &WarehouseState{
		CatalogItemId: id,
		ShopQuantity:  0,
		Reservation:   0,
	}
}

func (w *WarehouseState) GetID() CatalogItemId {
	return w.CatalogItemId
}

func (w *WarehouseState) GetAvailableQuantity() int {
	possibleQuantity := w.ShopQuantity - w.Reservation
	if possibleQuantity < 0 {
		return 0
	}
	return possibleQuantity
}
