package main

import (
	callparser "mia/CallParser"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Console struct {
	Code string `json: "code"`
}

var Call *callparser.CallParser

func main() {
	Call = &callparser.CallParser{}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	app.Post("/interpreter", parser)

	app.Listen(":8080")
}

func parser(c *fiber.Ctx) error {
	var console Console
	if err := c.BodyParser(&console); err != nil {
		return err
	}

	result := Call.ExecutionParser(console.Code)

	return c.JSON(fiber.Map{
		"message": result,
	})
}
