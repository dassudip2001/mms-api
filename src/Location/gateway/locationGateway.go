package locationGateway

import (
	locationController "mms-api/src/Location/controllers"

	"github.com/gofiber/fiber/v2"
)

func LocationRoute(c *fiber.App) {
	c.Post("/api/v1/locations", locationController.CreateLocation)
	c.Get("/api/v1/locations", locationController.GetAllLocations)
	c.Get("/api/v1/locations/:id", locationController.GetLocationById)
	c.Put("/api/v1/locations/:id", locationController.UpdateLocation)
	c.Delete("/api/v1/locations/:id", locationController.DeleteLocation)
}
