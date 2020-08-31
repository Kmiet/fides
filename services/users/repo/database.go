package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type database struct {
	client *mongo.Client
}

func InitDatabase(client *mongo.Client) UserRepository {
	return &database{client}
}

func (db *database) findUserById() {

}
