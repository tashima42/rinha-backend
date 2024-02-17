package main

import "github.com/gofiber/fiber/v2"

type Customer struct {
	ID      int `bson:"id"`
	Limit   int `bson:"limit"`
	Balance int `bson:"balance"`
}

func handleCreateTransaction(c *fiber.Ctx) error {

	return c.Status(200).JSON(map[string]interface{}{})
}
