package amqp

import (
	"os"
	"strconv"

	"github.com/Kmiet/fides/internal/contracts/events"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection

func Connect(uri string) {
	var err error
	conn, err = amqp.Dial(uri)
	panicOnError(err)
}

func Disconnect() {
	conn.Close()
}

func newChannel() *amqp.Channel {
	channel, err := conn.Channel()
	panicOnError(err)
	return channel
}

func declareExchange(channel *amqp.Channel, name string) {
	err := channel.ExchangeDeclare(
		name,    // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
	panicOnError(err)

	pc, ok := os.LookupEnv("RABBIT_MQ_PREFETCH_COUNT")
	if !ok {
		pc = "1"
	}
	prefetchCount, _ := strconv.Atoi(pc)

	err = channel.Qos(
		prefetchCount, // prefetch count
		0,             // prefetch size
		false,         // global
	)
	panicOnError(err)
}

func bindQueue(
	channel *amqp.Channel,
	queueName string,
	exchangeName string,
	topics []events.TopicName,
) *amqp.Queue {
	queue, err := channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	panicOnError(err)

	for _, t := range topics {
		err = channel.QueueBind(
			queue.Name,   // queue name
			t,            // routing key
			exchangeName, // exchange
			false,
			nil,
		)
		panicOnError(err)
	}

	return &queue
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
