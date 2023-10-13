package locationController

import (
	locationService "mms-api/src/Location/services"

	"github.com/gofiber/fiber/v2"
)

func CreateLocation(c *fiber.Ctx) error {
	return locationService.CreateLocation(c)
}

func GetAllLocations(c *fiber.Ctx) error {
	return locationService.GetLocations(c)
}

func GetLocationById(c *fiber.Ctx) error {
	return locationService.GetLocationById(c)
}

func UpdateLocation(c *fiber.Ctx) error {
	return locationService.UpdateLocation(c)
}

func DeleteLocation(c *fiber.Ctx) error {
	return locationService.DeleteLocation(c)
}
