package main

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/tomiok/shopping-cart/api"
)

func main() {
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

	api.RoutesUp(app)

	app.Listen(":3000")
}
