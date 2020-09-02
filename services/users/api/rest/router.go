package rest

import "github.com/gofiber/fiber"

func InitRouter(app *fiber.App) {
	api := app.Group("/users")

	api.Post("/", registerNewUser)
	api.Get("/:id", fetchUserWithID)
}
