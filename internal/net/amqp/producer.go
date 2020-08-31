package amqp

import (
	"github.com/Kmiet/fides/internal/contracts/events"
	"github.com/streadway/amqp"
)

type Producer interface {
	Close()
	Publish(body string)
}

type producer struct {
	exchangeName string
	channel      *amqp.Channel
	topic        events.TopicName
}

func InitProducer(exchangeName string, topic events.TopicName) Producer {
	channel := newChannel()
	declareExchange(channel, exchangeName)
	return &producer{
		exchangeName: exchangeName,
		channel:      channel,
		topic:        topic,
	}
}

func (p *producer) Close() {
	p.channel.Close()
}

func (p *producer) Publish(body string) {
	p.channel.Publish(
		p.exchangeName, // exchange
		p.topic,        // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		},
	)
}
