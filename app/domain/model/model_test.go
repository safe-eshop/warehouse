package model

import (
	"testing"
)

func TestNewWarehouseState(t *testing.T) {
	subject := NewWarehouseState("a", 1, 3)
	if subject.GetID() != "a" {
		t.Error("CatalogItemId should equal a")
	}
	if subject.ShopQuantity != 1 {
		t.Error("quantity should equal 1")
	}
	if subject.Reservation != 3 {
		t.Error("reservation should equal 3")
	}
}

func TestGetAvailableQuantityIfReservationIsGreater(t *testing.T) {
	warehouse := NewWarehouseState("a", 1, 3)
	subject := warehouse.GetAvailableQuantity()
	if subject == 0 {
		t.Log("Correct value")
	} else {
		t.Error("Available quantity should be 0")
	}
}

func TestGetAvailableQuantityIfReservationIsSmaller(t *testing.T) {
	warehouse := NewWarehouseState("a", 3, 1)
	subject := warehouse.GetAvailableQuantity()
	if subject == 2 {
		t.Log("Correct value")
	} else {
		t.Error("Available quantity should be 2")
	}
}
