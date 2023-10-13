package gateway

import (
	"mms-api/src/Services/controllers"

	"github.com/gofiber/fiber/v2"
)

func ServicesRoute(c *fiber.App) {
	c.Get("/api/v1/services", controllers.GetAllServices)
	c.Get("/api/v1/services/:id", controllers.GetServiceById)
	c.Post("/api/v1/services", controllers.CreateCatering)
	c.Put("/api/v1/services/:id", controllers.UpdateServices)
	c.Delete("/api/v1/services/:id", controllers.DeleteServices)
}
