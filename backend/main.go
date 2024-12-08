package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Phone int `json:"phone"`
}

func main() {
	fmt.Println("Empezamos a desarrollar")
	app := fiber.New()

	userRegister := []User{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello, World!"})
	})

	app.Post("/api/users", func(c *fiber.Ctx) error {
		user := &User{}

		if err := c.BodyParser(user); err != nil {
			return err
		}

		if user.ID == 0 {
			return c.Status(400).JSON(fiber.Map{"error":"ID de usuario es requerido"})
		}

		user.ID = len(userRegister) + 1 
		userRegister = append(userRegister, *user)

		return c.Status(201).JSON(user)
	})

	log.Fatal(app.Listen(":4000"))
}