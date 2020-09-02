package rest

import (
	users "github.com/Kmiet/fides/services/users"
	"github.com/gofiber/fiber"
)

var userService users.UserService

func InitHandlers(service users.UserService) {
	userService = service
}

func fetchUserWithID(ctx *fiber.Ctx) {
	userID := ctx.Params("id")
	if email, err := userService.FindUserWithID(userID); err != nil {
		ctx.Send(ctx.JSON(email))
	} else {
		ctx.SendStatus(404)
	}
}

func registerNewUser(ctx *fiber.Ctx) {
	if userID, err := userService.Register("test@test.pl"); err != nil {
		ctx.Send(ctx.JSON(userID))
	} else {
		ctx.SendStatus(500)
	}
}
