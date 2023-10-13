package services

import (
	"mms-api/db"
	"mms-api/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Catering struct {
	Name        string  `json:"name"`
	IsAvailable bool    `json:"is_available"`
	Price       float64 `json:"price"`
}

type CreateCateringRequest struct {
	Name        string  `json:"name"`
	IsAvailable bool    `json:"is_available"`
	Price       float64 `json:"price"`
}

func createResponseServices(servicesModel models.Catering) Catering {
	return Catering{
		Name:        servicesModel.Name,
		IsAvailable: servicesModel.IsAvailable,
		Price:       servicesModel.Price,
	}
}

// create a new services
func CreateServices(c *fiber.Ctx) error {
	var request CreateCateringRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if request.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "services name is required",
		})
	}

	var existingservices models.Catering

	if result := db.Database.Db.Where("name=?", request.Name).First(&existingservices); result.RowsAffected != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "services already exists",
		})
	}

	services := models.Catering{
		Model:       gorm.Model{},
		Name:        request.Name,
		IsAvailable: request.IsAvailable,
		Price:       request.Price,
	}

	if err := db.Database.Db.Create(&services).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "services creation failed",
		})
	}

	responseservices := createResponseServices(services)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "services created successfully",
		"data":    responseservices,
	})
}

// get all servicess
func GetAllServices(c *fiber.Ctx) error {
	var servicess []models.Catering
	if err := db.Database.Db.Find(&servicess).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve servicess",
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": "servicess retrieved successfully",
		"data":    servicess,
	})
}

// get a services by id
func GetServiceById(c *fiber.Ctx) error {
	id := c.Params("id")
	var services models.Catering

	if err := db.Database.Db.First(&services, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "services not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve services",
		})
	}
	return c.JSON(services)
}

// update servicess
func UpdateServices(c *fiber.Ctx) error {
	id := c.Params("id")
	var services models.Catering

	if err := db.Database.Db.First(&services, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   "services not found",
				"success": false,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrieve services",
			"success": false,
		})
	}

	if err := c.BodyParser(&services); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"success": false,
		})
	}

	if err := db.Database.Db.Save(&services).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update services",
			"success": false,
		})
	}
	return c.JSON(services)
}

// delete services
func DeleteServices(c *fiber.Ctx) error {
	id := c.Params("id")

	var services models.Catering

	if err := db.Database.Db.First(&services, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   "services not found",
				"success": false,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrieve services",
			"success": false,
		})
	}

	if err := db.Database.Db.Delete(&services).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to delete services",
			"success": false,
		})
	}
	return c.SendStatus(fiber.StatusOK)
}
