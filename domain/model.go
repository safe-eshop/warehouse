package domain

type WarehouseState struct {
	CatalogItemId string `json:"catalogItemId,omitempty"`
	Quantity      int    `json:"quantity,omitempty"`
	Reservation   int    `json:"reservation,omitempty"`
}

func NewWarehouseState(id string, quantity, reservation int) *WarehouseState {
	return &WarehouseState{
		CatalogItemId: id,
		Quantity:      quantity,
		Reservation:   reservation,
	}
}
