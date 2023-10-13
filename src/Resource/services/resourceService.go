package resources

import (
	"mms-api/db"
	"mms-api/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Resource struct {
	Name                   string `json:"name"`
	IsResourceAvailability bool   `json:"is_available"`
}

type CreateResourceRequest struct {
	Name                   string `json:"name"`
	IsResourceAvailability bool   `json:"is_available"`
}

func createResponseResource(resourceModel models.Resource) Resource {
	return Resource{
		Name:                   resourceModel.Name,
		IsResourceAvailability: resourceModel.IsResourceAvailability,
	}
}

// create a new resource

func CreateResource(c *fiber.Ctx) error {
	var request CreateResourceRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if request.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "resources name is required",
		})
	}

	var existingResources models.Resource

	if result := db.Database.Db.Where("name=?", request.Name).First(&existingResources); result.RowsAffected != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "resources already exists",
		})
	}

	resources := models.Resource{
		Model:                  gorm.Model{},
		Name:                   request.Name,
		IsResourceAvailability: request.IsResourceAvailability,
	}

	if err := db.Database.Db.Create(&resources).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "resources creation failed",
		})
	}

	responseResources := createResponseResource(resources)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "resources created successfully",
		"data":    responseResources,
	})
}

// get all resources
func GetAllResources(c *fiber.Ctx) error {
	var resources []models.Resource
	if err := db.Database.Db.Find(&resources).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "resources not found",
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"data":    resources,
	})
}

// get a single resources
func GetServiceById(c *fiber.Ctx) error {
	id := c.Params("id")
	var resources models.Resource

	if err := db.Database.Db.First(&resources, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "resources not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve resources",
		})
	}
	return c.JSON(resources)
}

// update a resources
func UpdateResources(c *fiber.Ctx) error {
	id := c.Params("id")
	var resources models.Resource

	if err := db.Database.Db.First(&resources, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   "resources not found",
				"success": false,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrieve resources",
			"success": false,
		})
	}

	if err := c.BodyParser(&resources); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"success": false,
		})
	}

	if err := db.Database.Db.Save(&resources).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update resources",
			"success": false,
		})
	}
	return c.JSON(resources)
}

// delete a resources

func DeleteResources(c *fiber.Ctx) error {
	id := c.Params("id")

	var resources models.Resource

	if err := db.Database.Db.First(&resources, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   "resources not found",
				"success": false,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrieve resources",
			"success": false,
		})
	}

	if err := db.Database.Db.Delete(&resources).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to delete resources",
			"success": false,
		})
	}
	return c.SendStatus(fiber.StatusOK)
}
