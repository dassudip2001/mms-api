package resourceGateway

import (
	"mms-api/src/Resource/controllers"

	"github.com/gofiber/fiber/v2"
)

func ResourceRoute(c *fiber.App) {
	c.Get("/api/v1/resources", controllers.GetAllResources)
	c.Get("/api/v1/resources/:id", controllers.GetResourcesById)
	c.Post("/api/v1/resources", controllers.CreateResources)
	c.Put("/api/v1/resources/:id", controllers.UpdateResources)
	c.Delete("/api/v1/resources/:id", controllers.DeleteResources)
}
