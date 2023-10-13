package web

import "github.com/gofiber/fiber/v2"

func welcome(c *fiber.Ctx) error {

	return c.Render("index", fiber.Map{"Title": "Server is running 👋!"})

	// return c.SendString("Server is running 👋!")

}

func SetupWebRoutes(app *fiber.App) {

	app.Get("/", welcome)

}
