package main

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/tomiok/shopping-cart/api"
	"github.com/tomiok/shopping-cart/storage"
	"time"
)

func main() {
	ctx, f := context.WithTimeout(context.Background(), 10*time.Second)
	defer f()
	mongoClient := storage.MongoClient(ctx)

	defer mongoClient.ClientMongo.Disconnect(ctx)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {

			if _, ok := err.(*fiber.Error); ok {
				return errors.New("this is a fiber error")
			}

			return errors.New("this is a managed error")
		},
	})

	// middlewares
	app.Use(recover.New())
	app.Use(logger.New())

	services := newServices(mongoClient)

	api.RoutesUp(services, app)

	app.Listen(":3000")
}

func newServices(mongoClient *storage.ShoppingStorage) api.Services {
	svcs := api.Services{
		BuySVC: api.BuyService{
			Buy: &api.BuyShopping{
				ShoppingStorage: mongoClient,
			},
		}}

	return svcs
}
