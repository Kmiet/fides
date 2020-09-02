package amqp

import (
	"github.com/Kmiet/fides/internal/contracts/events"
	"github.com/streadway/amqp"
)

type Consumer interface {
	Close()
	Receive() <-chan amqp.Delivery
}

type consumer struct {
	channel *amqp.Channel
	queue   *amqp.Queue
}

func InitConsumer(exchangeName string, queueName string, topics []events.TopicName) Consumer {
	channel := newChannel()
	declareExchange(channel, exchangeName)
	queue := bindQueue(channel, queueName, exchangeName, topics)
	return &consumer{
		channel: channel,
		queue:   queue,
	}
}

func (c *consumer) Close() {
	c.channel.Close()
}

func (c *consumer) Receive() <-chan amqp.Delivery {
	msgs, err := c.channel.Consume(
		c.queue.Name, // queue
		"",           // consumer
		true,         // auto ack
		false,        // exclusive
		false,        // no local
		false,        // no wait
		nil,          // args
	)
	panicOnError(err)

	return msgs
}
