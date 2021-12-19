package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type User struct {
	Id         int
	Age        int
	First_Name string
	Last_Name  string
	Framework  string
}

func getUsers() {
	var users [1000]User

	for i := 1; i < 1001; i++ {
		var stringIndex = strconv.Itoa(i)
		users[i-1] = User{
			Id:         i,
			Age:        25,
			First_Name: "First_name" + stringIndex,
			Last_Name:  "Last_Name" + stringIndex,
			Framework:  "Go fiber",
		}
	}
}

func main() {
	app := fiber.New()

	app.Get("/api/v1/users", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
