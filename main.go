package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rubiagatra/coba-github-action/calc"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello, World 👋! Add(3 + 3 = %d)", calc.Add(3, 3)))
	})

	port := os.Getenv("PORT")
	fmt.Println(port)
	app.Listen(":" + port)
}
