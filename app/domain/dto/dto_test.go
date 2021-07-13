package dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/app/domain/model"
)

func TestNewAvailableQuantity(t *testing.T) {
	subject := NewAvailableQuantity("a", 2)
	assert.Equal(t, subject.Id, "a", "they should be equal")
	assert.Equal(t, subject.AvailableQuantity, 2, "they should be equal")
}

func TestFromWarehouseState(t *testing.T) {
	state := model.NewWarehouseState("a", 5, 3)
	subject := FromWarehouseState(*state)
	assert.Equal(t, subject.Id, "a", "they should be equal")
	assert.Equal(t, subject.AvailableQuantity, 2, "they should be equal")
}

func TestFromWarehouseStates(t *testing.T) {
	states := []*model.WarehouseState{model.NewWarehouseState("a", 5, 3), model.NewWarehouseState("b", 7, 3)}
	subject := FromWarehouseStates(states)
	for i, state := range states {

		assert.Equal(t, state.GetID(), subject[i].Id, "they should be equal")
		assert.Equal(t, state.GetAvailableQuantity(), subject[i].AvailableQuantity, "they should be equal")

	}
}
