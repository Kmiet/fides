package rest

import (
	"fmt"

	users "github.com/Kmiet/fides/services/users"
	"github.com/gofiber/fiber"
)

var userService users.UserService

func InitHandlers(service users.UserService) {
	userService = service
}

func fetchUserWithID(ctx *fiber.Ctx) {
	userID := ctx.Params("id")
	if user, err := userService.FindUserWithID(userID); err == nil {
		fmt.Println(user)
		ctx.Send(ctx.JSON(fiber.Map{
			"id":    user.(map[string]interface{})["_id"],
			"email": user.(map[string]interface{})["email"],
		}))
	} else {
		ctx.SendStatus(404)
	}
}

func registerNewUser(ctx *fiber.Ctx) {
	if userID, err := userService.Register("test@test.pl"); err == nil {
		ctx.Send(ctx.JSON(fiber.Map{
			"id": userID,
		}))
	} else {
		ctx.SendStatus(500)
	}
}
