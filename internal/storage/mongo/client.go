package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitClient(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client
}

func TestConnection(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
}

func Disconnect(client *mongo.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client.Disconnect(ctx)
}
