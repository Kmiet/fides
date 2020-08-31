package amqp

import (
	"github.com/Kmiet/fides/internal/net/amqp"
)

const X = "SESESE"

func InitConsumer(channel amqp.Channel) {
	channel.DeclareExchange()
	channel.BindQueue(X, []string{
		"sagas.client",
	})
}
