package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMqClient struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewRabbitMqClient(connStr string) (*RabbitMqClient, error) {
	conn, err := amqp.Dial(connStr)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &RabbitMqClient{Connection: conn, Channel: ch}, nil
}

func (client *RabbitMqClient) Close() error {
	err := client.Channel.Close()
	if err != nil {
		return fmt.Errorf("channel close error %w", err)
	}
	err = client.Connection.Close()
	if err != nil {
		return fmt.Errorf("client connection close error %w", err)
	}
	return nil
}
