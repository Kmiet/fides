package amqp

import (
	"github.com/Kmiet/fides/services/users"
	"github.com/streadway/amqp"
)

var userService users.UserService

func InitHandlers(service users.UserService) {
	userService = service
}

func handleEvent(msg amqp.Delivery) {
	return
}
