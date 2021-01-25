package api

import (
	"context"
	"github.com/tomiok/shopping-cart/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Services struct {
	BuySVC BuyService
}

type BuyService struct {
	Buy Buy
}

type Buy interface {
	BuyItem(i Item) (string, error)
}

type BuyShopping struct {
	*storage.ShoppingStorage
}

func (b *BuyShopping) BuyItem(i Item) (string, error) {
	client := b.ClientMongo

	res, err := client.Database(storage.MongoDBName).
		Collection("buyCollection").
		InsertOne(context.Background(), i)

	if err != nil {
		return "", err
	}
	oid := res.InsertedID.(primitive.ObjectID)
	return oid.Hex(), nil
}

type Item struct {
	ID       string `json:"id,omitempty" bson:"id,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Price    int    `json:"price,omitempty" bson:"usd,omitempty"`
	ObjectID string `json:"object_id,omitempty" bson:"-"`
}

type StockService struct {
	Stock
}

type Stock interface {
	AddItem(i Item) (*Item, error)
}
