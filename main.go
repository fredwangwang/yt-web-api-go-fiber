package main

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/semaphore"
)

type User struct {
	Id         int
	Age        int
	First_Name string
	Last_Name  string
	Framework  string
}

// func getUsers() [1000]User {
func getUsers() []User {
	var users [1000]User

	for i := 0; i < 1000; i++ {
		var stringIndex = strconv.Itoa(i)
		users[i] = User{
			Id:         i,
			Age:        25,
			First_Name: "First_name" + stringIndex,
			Last_Name:  "Last_Name" + stringIndex,
			Framework:  "Golang (fiber)  ",
		}
	}

	// return users
	return users[:] // https://go.dev/ref/spec#Slice_expressions
}

func main() {
	app := fiber.New()

	go func() {
		for {
			fmt.Println("num go routine", runtime.NumGoroutine())
			time.Sleep(time.Second)
		}
	}()

	ctx := context.TODO()
	sem := semaphore.NewWeighted(int64(runtime.GOMAXPROCS(0)))

	app.Get("/api/v1/users", func(c *fiber.Ctx) error {
		sem.Acquire(ctx, 1)
		defer sem.Release(1)
		return c.JSON(getUsers())
	})

	app.Listen(":3000")
}
