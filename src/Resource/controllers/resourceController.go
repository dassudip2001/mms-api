package controllers

import (
	resources "mms-api/src/Resource/services"

	"github.com/gofiber/fiber/v2"
)

func CreateResources(c *fiber.Ctx) error {
	return resources.CreateResource(c)
}

func GetAllResources(c *fiber.Ctx) error {
	return resources.GetAllResources(c)
}

func GetResourcesById(c *fiber.Ctx) error {
	return resources.GetServiceById(c)
}

func UpdateResources(c *fiber.Ctx) error {
	return resources.UpdateResources(c)
}

func DeleteResources(c *fiber.Ctx) error {
	return resources.DeleteResources(c)
}
