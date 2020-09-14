package main

import (
	"os"

	"github.com/Kmiet/fides/internal/contracts/events"

	"github.com/Kmiet/fides/internal/net/amqp"
	"github.com/Kmiet/fides/internal/storage/mongo"
	"github.com/Kmiet/fides/internal/storage/redis"
	"github.com/Kmiet/fides/services/users"
	amqpAPI "github.com/Kmiet/fides/services/users/api/amqp"
	restAPI "github.com/Kmiet/fides/services/users/api/rest"
	"github.com/Kmiet/fides/services/users/repo/cache"
	"github.com/Kmiet/fides/services/users/repo/db"

	"github.com/gofiber/fiber"
)

var (
	MONGO_URI     = os.Getenv("MONGO_URI")
	PORT          = os.Getenv("PORT")
	RABBIT_MQ_URI = os.Getenv("RABBIT_MQ_URI")
	REDIS_URI     = os.Getenv("REDIS_URI")
)

const (
	EVENT_EXCHANGE          = "events"
	USER_SERVICE_QUEUE_NAME = "users"
)

func main() {
	dbClient := mongo.InitClient(MONGO_URI)
	mongo.TestConnection(dbClient)
	defer mongo.Disconnect(dbClient)

	dbRepo := db.InitRepository(dbClient)

	redisClient := redis.InitClient(REDIS_URI)
	redis.TestConnection(redisClient)
	defer redis.Disconnect(redisClient)

	cacheRepo := cache.InitRepository(redisClient)

	amqp.Connect(RABBIT_MQ_URI)
	defer amqp.Disconnect()
	amqpConsumer := amqp.InitConsumer(
		EVENT_EXCHANGE,
		USER_SERVICE_QUEUE_NAME,
		[]string{events.TOPICS.USER_SERVICE},
	)
	defer amqpConsumer.Close()
	amqpProducer := amqp.InitProducer(EVENT_EXCHANGE, events.TOPICS.USER_SERVICE)
	defer amqpProducer.Close()

	userService := users.InitService(cacheRepo, dbRepo, amqpProducer)

	amqpAPI.InitHandlers(userService)
	go amqpAPI.Run(amqpConsumer)

	app := fiber.New()

	restAPI.InitHandlers(userService)
	restAPI.InitRouter(app)

	app.Listen(PORT)
}
