package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	model2 "warehouse/app/domain/model"
)

func TestToWarehouseState(t *testing.T) {
	model := RedisWarehouseState{
		CatalogItemId: "Test",
		ShopQuantity:  3,
		Reservation:   1,
	}
	subject := model.ToWarehouseState()
	assert.Equal(t, subject.GetID(), "Test", "they should be equal")
	assert.Equal(t, subject.ShopQuantity, 3, "they should be equal")
	assert.Equal(t, subject.Reservation, 1, "they should be equal")
}

func TestFromWarehouseState(t *testing.T) {
	model := model2.WarehouseState{
		CatalogItemId: "Test",
		ShopQuantity:  3,
		Reservation:   1,
	}
	subject := FromWarehouseState(model)
	assert.Equal(t, subject.CatalogItemId, "Test", "they should be equal")
	assert.Equal(t, subject.ShopQuantity, 3, "they should be equal")
	assert.Equal(t, subject.Reservation, 1, "they should be equal")
}
