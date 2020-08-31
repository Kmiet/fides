package amqp

import (
	"github.com/Kmiet/fides/internal/net/amqp"
)

var ProducerChannel amqp.Channel

func InitProducer(channel amqp.Channel) {
	channel.DeclareExchange()
	ProducerChannel = channel
}
