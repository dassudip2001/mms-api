package main

import (
	"mms-api/db"
	"mms-api/routes/apis"
	"mms-api/routes/web"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// initialize the database connection
	db.ConnectDb()
	// load the api routes
	apis.ServicesApi(app)

	// load the web routes
	web.SetupWebRoutes(app)

	// load the static files
	app.Static("/", "./public")

	app.Listen(":8000")
}
