package main

import (
	"fmt"
	"os"

	callparser "mia/CallParser"
	env "mia/Classes/Env"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Console struct {
	Code string `json:"code"`
}

type PartitionForm struct {
	DiskName string `json:"disk"`
}

var (
	Call           *callparser.CallParser
	Conversaciones [][]string
	Partitions     [][]string
)

func main() {
	Call = &callparser.CallParser{}
	Conversaciones = [][]string{}
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
	app.Get("/messages", getMessages)
	app.Post("/partitions", getPartitions)

	app.Listen(":8080")
}

func parser(c *fiber.Ctx) error {
	var console Console
	if err := c.BodyParser(&console); err != nil {
		return err
	}

	result := Call.ExecutionParser(console.Code)

	conversation := []string{console.Code, result}
	Conversaciones = append(Conversaciones, conversation)

	return c.JSON(fiber.Map{
		"message": result,
	})
}

func verChat(chat [][]string) {
	for i := 0; i < len(chat); i++ {
		for j := 0; j < len(chat[i]); j++ {
			fmt.Printf("-> %s\n", chat[i][j])
		}
	}
}

func getFiles(c *fiber.Ctx) error {
	directory := "/home/jefferson/Escritorio/SO"

	files, err := obtenerArchivosEnDirectorio(directory)
	if err != nil {
		return err
	}

	var names []string
	for _, file := range files {
		names = append(names, file.Name())
	}

	return c.JSON(fiber.Map{
		"files": names,
	})
}

func obtenerArchivosEnDirectorio(directory string) ([]os.FileInfo, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var files []os.FileInfo
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}
		files = append(files, info)
	}

	return files, nil
}

func getMessages(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"messages": Conversaciones,
	})
}

func getPartitions(c *fiber.Ctx) error {
	var partForm PartitionForm
	if err := c.BodyParser(&partForm); err != nil {
		return err
	}

	partitions(partForm.DiskName)
	_Partitions := Partitions
	Partitions = [][]string{}

	return c.JSON(fiber.Map{
		"partitions": _Partitions,
	})
}

func partitions(diskName string) {
	if disk, ok := env.Disks[diskName]; ok {
		for id, name := range disk.Ids {
			partition := []string{id, name.Name}
			Partitions = append(Partitions, partition)
		}
	}
}
