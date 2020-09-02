package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Kmiet/fides/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	RegisterNewUser(email string) (string, error)
	FindUserByID(id string)
}

type database struct {
	_database *mongo.Database
	users     *mongo.Collection
}

const TIMEOUT_DURATION = 3 * time.Second

func InitRepository(client *mongo.Client) Repository {
	db := client.Database(storage.DATABASE_NAME)
	db.CreateCollection(context.Background(), storage.DATA_TYPES.USERS)
	collection := db.Collection(storage.DATA_TYPES.USERS)
	return &database{
		_database: db,
		users:     collection,
	}
}

func (db *database) RegisterNewUser(email string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_DURATION)
	defer cancel()
	res, err := db.users.InsertOne(ctx, bson.M{
		"email": email,
	})
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (db *database) FindUserByID(id string) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_DURATION)
	defer cancel()
	db.users.FindOne(ctx, bson.M{
		"_id": bson.M{
			"$eq": id,
		},
	})
}
