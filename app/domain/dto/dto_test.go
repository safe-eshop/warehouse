package dto

import (
	"testing"
	"warehouse/app/domain/model"

	"github.com/stretchr/testify/assert"
)

func TestNewAvailableQuantity(t *testing.T) {
	subject := NewAvailableQuantity(1, 2)
	assert.Equal(t, subject.Id, 1, "they should be equal")
	assert.Equal(t, subject.AvailableQuantity, 2, "they should be equal")
}

func TestFromWarehouseState(t *testing.T) {
	state := model.NewWarehouseState(1, 5, 3)
	subject := FromWarehouseState(*state)
	assert.Equal(t, subject.Id, 1, "they should be equal")
	assert.Equal(t, subject.AvailableQuantity, 2, "they should be equal")
}

func TestFromWarehouseStates(t *testing.T) {
	states := []*model.WarehouseState{model.NewWarehouseState(1, 5, 3), model.NewWarehouseState(2, 7, 3)}
	subject := FromWarehouseStates(states)
	for i, state := range states {

		assert.Equal(t, state.GetID(), subject[i].Id, "they should be equal")
		assert.Equal(t, state.GetAvailableQuantity(), subject[i].AvailableQuantity, "they should be equal")

	}
}
