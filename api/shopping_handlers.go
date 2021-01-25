package api

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func (b *BuyService) buyItemHandler(ctx *fiber.Ctx) error {
	var item Item
	err := ctx.BodyParser(&item)

	if err != nil {
		log.Println(err)
		return err
	}

	id, err := b.Buy.BuyItem(item)

	if err != nil {
		log.Println(err)
		return err
	}

	item.ObjectID = id

	return ctx.JSON(&item)
}

type BuySummary struct {
	Summary     string `json:"summary"`
	TotalAmount int    `json:"total_amount"`
}
