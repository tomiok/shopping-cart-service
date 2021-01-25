package api

import "github.com/gofiber/fiber/v2"

func RoutesUp(services Services, app *fiber.App) {
	grp := app.Group("/shopping")
	grp.Post("/buy", services.BuySVC.buyItemHandler)
}
