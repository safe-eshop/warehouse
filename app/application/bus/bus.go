package bus

import (
	"context"
	"warehouse/app/domain/model"
	"warehouse/app/domain/repository"

	log "github.com/sirupsen/logrus"
)

type MessageSubscriber interface {
	StartHandling(ctx context.Context) chan ProductCreated
}

type ProductCreated struct {
	Id int `json:"id,omitempty"`
}

func HandleProductCreated(ctx context.Context, subscriber MessageSubscriber, repository repository.WarehouseStateRepository) {
	for msg := range subscriber.StartHandling(ctx) {
		stock := model.NewWarehouseState(msg.Id, 10, 0)
		err := repository.Insert(ctx, stock)
		if err != nil {
			log.WithField("Id", stock.CatalogItemId).WithError(err).WithContext(ctx).Errorln("Insert new warehouse failed")
		}
	}
}
