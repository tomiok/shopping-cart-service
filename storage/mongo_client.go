package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	MongoDBName = "mongoTest"
)

type ShoppingStorage struct {
	ClientMongo *mongo.Client
}

func MongoClient(ctx context.Context) *ShoppingStorage {
	opt := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, opt)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return &ShoppingStorage{
		ClientMongo: client,
	}
}















