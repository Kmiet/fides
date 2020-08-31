package main

import (
	"os"

	"github.com/Kmiet/fides/internal/net/amqp"
	"github.com/Kmiet/fides/internal/storage/mongo"
	"github.com/Kmiet/fides/internal/storage/redis"
	"github.com/Kmiet/fides/services/users"
	amqpAPI "github.com/Kmiet/fides/services/users/api/amqp"
	restAPI "github.com/Kmiet/fides/services/users/api/rest"
	"github.com/Kmiet/fides/services/users/repo"

	"github.com/gofiber/fiber"
)

var (
	MONGO_URI     = os.Getenv("MONGO_URI")
	PORT          = os.Getenv("PORT")
	RABBIT_MQ_URI = os.Getenv("RABBIT_MQ_URI")
	REDIS_URI     = os.Getenv("REDIS_URI")
)

func main() {
	dbClient := mongo.InitClient(MONGO_URI)
	mongo.TestConnection(dbClient)
	defer mongo.Disconnect(dbClient)

	db := repo.InitDatabase(dbClient)

	redisClient := redis.InitClient(REDIS_URI)
	redis.TestConnection(redisClient)
	defer redis.Disconnect(redisClient)

	cache := repo.InitCache(redisClient)

	amqp.Connect(RABBIT_MQ_URI)
	defer amqp.Disconnect()
	amqpConsumer := amqp.InitConsumer("", "", []string{""})
	defer amqpConsumer.Close()
	amqpProducer := amqp.InitProducer("", "")
	defer amqpProducer.Close()

	userService := users.InitService(cache, db, amqpProducer)

	amqpAPI.InitHandlers(userService)
	go amqpAPI.Run(amqpConsumer)

	app := fiber.New()

	restAPI.InitHandlers(userService)
	restAPI.InitRouter(app)

	app.Listen(PORT)
}
