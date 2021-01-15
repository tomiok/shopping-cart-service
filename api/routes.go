package api

import "github.com/gofiber/fiber/v2"

func RoutesUp(app *fiber.App) {
	grp := app.Group("/shopping")

	grp.Get("/buy", buyItemHandler)
}
