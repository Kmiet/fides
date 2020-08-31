package amqp

import (
	"os"
	"strconv"

	"github.com/streadway/amqp"
)

type Channel interface {
	Close()
	DeclareExchange()
	BindQueue(name string, topics []string)
	Publish(body string)
	Receive() <-chan amqp.Delivery
}

type client struct {
	_exchangeName string

	channel *amqp.Channel
	queue   *amqp.Queue
}

const EVENT_EXCHANGE = "EVENT_EXCHANGE"

var conn *amqp.Connection

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func NewChannel(uri string) Channel {
	if conn == nil {
		c, err := amqp.Dial(uri)
		panicOnError(err)
		conn = c
	}

	channel, err := conn.Channel()
	panicOnError(err)

	return &client{
		channel: channel,
	}
}

func Disconnect() {
	conn.Close()
}

func (c *client) Close() {
	c.channel.Close()
}

func (c *client) DeclareExchange() {
	err := c.channel.ExchangeDeclare(
		EVENT_EXCHANGE, // name
		"topic",        // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	panicOnError(err)

	pc, ok := os.LookupEnv("RABBIT_MQ_PREFETCH_COUNT")
	if !ok {
		pc = "1"
	}
	prefetchCount, _ := strconv.Atoi(pc)

	err = c.channel.Qos(
		prefetchCount, // prefetch count
		0,             // prefetch size
		false,         // global
	)
	panicOnError(err)
}

func (c *client) BindQueue(name string, topics []string) {
	queue, err := c.channel.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	panicOnError(err)
	c.queue = &queue

	for _, t := range topics {
		err = c.channel.QueueBind(
			c.queue.Name,    // queue name
			t,               // routing key
			c._exchangeName, // exchange
			false,
			nil,
		)
		panicOnError(err)
	}
}

func (c *client) Publish(body string) {
	c.channel.Publish(
		c._exchangeName, // exchange
		"",              // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		},
	)
}

func (c *client) Receive() <-chan amqp.Delivery {
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
