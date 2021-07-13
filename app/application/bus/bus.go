package bus

import (
	"github.com/streadway/amqp"
	"log"
)

type ProductCreated struct {
	Id string `json:"id,omitempty"`
}

func HandleProductCreated(conn *amqp.Connection) {
	channel, err := conn.Channel()
	if err != nil {
		log.Fatal("")
	}
	defer channel.Close()
}
