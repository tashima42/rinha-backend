package main

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	dbHostname := os.Getenv("DB_HOSTNAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	port := os.Getenv("PORT")
	skipPopulate := os.Getenv("SKIP_POPULATE")

	mongoURI := "mongodb://" + dbUsername + ":" + dbPassword + "@" + dbHostname
	ctx := context.Background()
	client, err := mongoConnect(ctx, mongoURI)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if skipPopulate != "1" {
		if err := populateCustomers(ctx, client); err != nil {
			panic(err)
		}
	}

	app := fiber.New()

	if err != nil {
		panic(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello, world")
	})

	app.Listen(":" + port)
}
