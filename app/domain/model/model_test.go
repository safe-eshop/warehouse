package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWarehouseState(t *testing.T) {
	subject := NewWarehouseState(1, 1, 3)
	assert.Equal(t, subject.GetID(), "a", "they should be equal")
	assert.Equal(t, subject.ShopQuantity, 1, "they should be equal")
	assert.Equal(t, subject.Reservation, 3, "they should be equal")
}

func TestZeroWarehouseState(t *testing.T) {
	subject := Zero(1)
	assert.Equal(t, subject.GetID(), 1, "they should be equal")
	assert.Equal(t, subject.ShopQuantity, 0, "they should be equal")
	assert.Equal(t, subject.Reservation, 0, "they should be equal")
}

func TestGetAvailableQuantityIfReservationIsGreater(t *testing.T) {
	warehouse := NewWarehouseState(1, 1, 3)
	subject := warehouse.GetAvailableQuantity()
	assert.Equal(t, subject, 0, "they should be equal")
}

func TestGetAvailableQuantityIfReservationIsSmaller(t *testing.T) {
	warehouse := NewWarehouseState(1, 3, 1)
	subject := warehouse.GetAvailableQuantity()
	assert.Equal(t, subject, 2, "they should be equal")
}
