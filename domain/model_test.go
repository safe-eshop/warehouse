package domain

import "testing"

func TestNewWarehouseState(t *testing.T) {
	subject := NewWarehouseState("a", 1, 3)
	if subject.CatalogItemId != "a" {
		t.Error("CatalogItemId should equal a")
	}
	if subject.Quantity != 1 {
		t.Error("quantity should equal 1")
	}
	if subject.Reservation != 3 {
		t.Error("reservation should equal 3")
	}
}
