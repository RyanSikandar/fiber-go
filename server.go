package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:       "First Fiber App",
		CaseSensitive: true,
	})

	micro := fiber.New()
	app.Group("/rawr")
	app.Mount("/micro", micro)
	//print the config
	// cfg := app.Config()
	// cfgData, _ := json.MarshalIndent(cfg, "", "  ")
	// fmt.Println(string(cfgData))


	micro.Get("/ryan", func(c *fiber.Ctx) error {
		return c.SendString("meow")
	})

	if fiber.IsChild() {
		println("This is a child process")
	} else {
		println("This is the parent process")
	}
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	match := fiber.RoutePatternMatch("/api/lmao/rawr/", "/api/:version/:chuss/")
	fmt.Println("Match:", match) // Match: map[version:lmao chuss:rawr]
	app.Get("/value/:id", func(c *fiber.Ctx) error {
		return c.SendString("params is " + c.Params("id"))
	}).Name("Hello 125")

	app.Get("/name/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("params is " + c.Params("name"))
		}
		return fiber.NewError(400, "no params")
	})

	data, _ := json.MarshalIndent(app.GetRoute("index"), "", "  ")
	fmt.Print(string(data))
	app.Use("/api", func(c *fiber.Ctx) error {
		c.Set("x-lmao-header", fmt.Sprint(rand.Int()))

		return c.Next()
	})
	// Simple GET handler
	app.Get("/api/list", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})
	
	data, _ = json.MarshalIndent(app.Stack(), "", "  ")
    fmt.Println(string(data))

	app.Listen(":3000")
}
