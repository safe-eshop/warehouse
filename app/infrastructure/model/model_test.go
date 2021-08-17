package model

import (
	"testing"
	model2 "warehouse/app/domain/model"

	"github.com/stretchr/testify/assert"
)

func TestToWarehouseState(t *testing.T) {
	model := RedisWarehouseState{
		CatalogItemId: 1,
		ShopQuantity:  3,
		Reservation:   1,
	}
	subject := model.ToWarehouseState()
	assert.Equal(t, subject.GetID(), 1, "they should be equal")
	assert.Equal(t, subject.ShopQuantity, 3, "they should be equal")
	assert.Equal(t, subject.Reservation, 1, "they should be equal")
}

func TestFromWarehouseState(t *testing.T) {
	model := model2.WarehouseState{
		CatalogItemId: 1,
		ShopQuantity:  3,
		Reservation:   1,
	}
	subject := FromWarehouseState(model)
	assert.Equal(t, subject.CatalogItemId, 1, "they should be equal")
	assert.Equal(t, subject.ShopQuantity, 3, "they should be equal")
	assert.Equal(t, subject.Reservation, 1, "they should be equal")
}
