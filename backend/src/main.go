package main

import (
	callparser "mia/CallParser"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Console struct {
	Code string `json:"code"`
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
	app.Get("/files", getFiles)

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

func getFiles(c *fiber.Ctx) error {
	directorio := "/home/jefferson/Escritorio/SO"

	archivos, err := obtenerArchivosEnDirectorio(directorio)
	if err != nil {
		return err
	}

	var nombres []string
	for _, archivo := range archivos {
		nombres = append(nombres, archivo.Name())
	}

	return c.JSON(fiber.Map{
		"files": nombres,
	})
}

func obtenerArchivosEnDirectorio(directorio string) ([]os.FileInfo, error) {
	entradas, err := os.ReadDir(directorio)
	if err != nil {
		return nil, err
	}

	var archivos []os.FileInfo
	for _, entrada := range entradas {
		info, err := entrada.Info()
		if err != nil {
			return nil, err
		}
		archivos = append(archivos, info)
	}

	return archivos, nil
}
