package main

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {
	PORTSERVER := ":3000"
	app := fiber.New()
	app.Listen(PORTSERVER)
	fmt.Println("server work in " + PORTSERVER + "port")
}
