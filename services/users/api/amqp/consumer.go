package amqp

import (
	"github.com/Kmiet/fides/internal/net/amqp"
)

func Run(consumer amqp.Consumer) {
	msgs := consumer.Receive()
	for msg := range msgs {
		handleEvent(msg)
	}
}
