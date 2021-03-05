package main

import (
	"os"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func main() {
	user := os.Getenv("MYSQL_USER")
  pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
  dbname := os.Getenv("MYSQL_DATABASE")
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, dbname)

	_, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8000")
}
