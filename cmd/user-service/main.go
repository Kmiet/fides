package main

import (
	"os"

	"github.com/Kmiet/fides/internal/net/amqp"
	"github.com/Kmiet/fides/internal/storage/mongo"
	"github.com/Kmiet/fides/services/users"
	"github.com/Kmiet/fides/services/users/api/rest"
	"github.com/Kmiet/fides/services/users/repo"
	"github.com/gofiber/fiber"
)

var (
	MONGO_URI     = os.Getenv("MONGO_URI")
	PORT          = os.Getenv("PORT")
	RABBIT_MQ_URI = os.Getenv("RABBIT_MQ_URI")
)

func main() {
	dbClient := mongo.InitClient(MONGO_URI)

	db := repo.InitDatabase(dbClient)

	consumerChannel := amqp.NewChannel(RABBIT_MQ_URI)
	consumerChannel.DeclareExchange()
	defer consumerChannel.Close()
	producerChannel := amqp.NewChannel(RABBIT_MQ_URI)
	producerChannel.DeclareExchange()
	defer producerChannel.Close()

	defer amqp.Disconnect()

	userService := users.InitService(nil, db, producerChannel)

	app := fiber.New()

	rest.InitHandlers(userService)
	rest.InitRouter(app)

	app.Listen(PORT)
}
