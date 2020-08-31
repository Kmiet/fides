package rest

import (
	users "github.com/Kmiet/fides/services/users"
)

var userService users.UserService

func InitHandlers(service users.UserService) {
	userService = service
}
