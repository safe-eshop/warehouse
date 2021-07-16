package model

type ProductId = int
type WarehouseState struct {
	CatalogItemId ProductId `json:"catalogItemId,omitempty"`
	ShopQuantity  int       `json:"quantity,omitempty"`
	Reservation   int       `json:"reservation,omitempty"`
}

func NewWarehouseState(id ProductId, quantity, reservation int) *WarehouseState {
	return &WarehouseState{
		CatalogItemId: id,
		ShopQuantity:  quantity,
		Reservation:   reservation,
	}
}

func Zero(id ProductId) *WarehouseState {
	return &WarehouseState{
		CatalogItemId: id,
		ShopQuantity:  0,
		Reservation:   0,
	}
}

func (w *WarehouseState) GetID() ProductId {
	return w.CatalogItemId
}

func (w *WarehouseState) GetAvailableQuantity() int {
	possibleQuantity := w.ShopQuantity - w.Reservation
	if possibleQuantity < 0 {
		return 0
	}
	return possibleQuantity
}
