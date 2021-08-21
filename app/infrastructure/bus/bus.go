package bus

import (
	"context"
	"encoding/json"
	"warehouse/app/application/bus"
	"warehouse/app/infrastructure/rabbitmq"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type rabbitMqMessageSubscriber struct {
	Client    *rabbitmq.RabbitMqClient
	Exchange  string
	QueueName string
	Topic     string
}

func createMessage(msg amqp.Delivery) (bus.ProductCreated, error) {
	var result bus.ProductCreated
	err := json.Unmarshal(msg.Body, &result)
	return result, err
}

func (source rabbitMqMessageSubscriber) StartHandling(ctx context.Context) chan bus.ProductCreated {
	err := source.Client.Channel.ExchangeDeclare(
		source.Exchange, // name
		"topic",         // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)

	if err != nil {
		log.WithError(err).Fatal("declare exchange error")
	}

	q, err := source.Client.Channel.QueueDeclare(
		source.Exchange+"-"+source.QueueName, // name
		true,                                 // durable
		false,                                // delete when usused
		false,                                // exclusive
		false,                                // no-wait
		nil,                                  // arguments
	)

	if err != nil {
		log.WithError(err).Fatal("declare queue error")
	}

	err = source.Client.Channel.QueueBind(
		q.Name,          // queue name
		source.Topic,    // routing key
		source.Exchange, // exchange
		false,
		nil,
	)

	if err != nil {
		log.WithError(err).Fatal("queue bind error")
	}

	msgs, err := source.Client.Channel.Consume(
		q.Name,                               // queue
		source.Exchange+"-"+source.QueueName, // consumer
		true,                                 // auto-ack
		false,                                // exclusive
		false,                                // no-local
		false,                                // no-wait
		nil,                                  // args
	)

	if err != nil {
		log.WithError(err).Fatal("create message channel error")
	}
	stream := make(chan bus.ProductCreated, 10)
	go func(evtChan chan bus.ProductCreated) {
		for msg := range msgs {
			res, err := createMessage(msg)
			if err != nil {
				log.WithError(err).Errorln("Message parsing error")
			}
			evtChan <- res
		}
		close(stream)
	}(stream)

	return stream
}

func NewRabbitMqMessageSubscriber(client *rabbitmq.RabbitMqClient, exchnage, queue, topic string) rabbitMqMessageSubscriber {
	return rabbitMqMessageSubscriber{Client: client, Exchange: exchnage, QueueName: queue, Topic: topic}
}
