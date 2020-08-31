package repo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type database struct {
	_database *mongo.Database
	users     *mongo.Collection
}

const TIMEOUT_DURATION = 3 * time.Second
const USERS_COLLECTION = "users"

func InitDatabase(client *mongo.Client) UserRepository {
	db := client.Database(USERS_COLLECTION)
	db.CreateCollection(context.Background(), USERS_COLLECTION)
	collection := db.Collection(USERS_COLLECTION)
	return &database{
		_database: db,
		users:     collection,
	}
}

func (db *database) findUserByID(id string) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_DURATION)
	defer cancel()
	db.users.FindOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": id,
		},
	})
}
