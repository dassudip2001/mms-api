package locationService

import (
	"mms-api/db"
	"mms-api/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CreateLocationRequest struct {
	Name     string `json:"name"`
	ParentID *uint  `json:"parent_id"`
}

type Location struct {
	Name string `json:"name"`
	// `ParentID *uint  `json:"parent_id"` is a field in the `CreateLocationRequest` struct. It is used to
	// represent the parent ID of a location when creating a new location.
	ParentID *uint `json:"parent_id"`
}

func createResponseLocation(locationModel models.Location) Location {
	return Location{
		Name:     locationModel.Name,
		ParentID: locationModel.ParentID,
	}
}

// create a new location
func CreateLocation(c *fiber.Ctx) error {
	var request CreateLocationRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if request.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Location name is required",
		})
	}

	// Check if the name is already in use
	var existingLocation models.Location

	result := db.Database.Db.Where("name = ?", request.Name).First(&existingLocation)
	if result.Error == nil {
		// Name already exists, return an error
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": "Location name must be unique",
		})
	}

	location := models.Location{
		Name:     request.Name,
		ParentID: request.ParentID,
	}

	// Create the location
	if err := db.Database.Db.Create(&location).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Could not create location",
		})
	}

	responseLocation := createResponseLocation(location)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Location created successfully",
		"data":    responseLocation,
	})
}

// Get all locations
func GetLocations(c *fiber.Ctx) error {
	var locations []models.Location

	db.Database.Db.Preload("Children").Find(&locations)
	return c.JSON(fiber.Map{
		"success": true,
		"data":    locations,
	})
}

// get location by id
func GetLocationById(c *fiber.Ctx) error {
	id := c.Params("id")

	var location models.Location

	if db.Database.Db.Preload("Children").First(&location, id).Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "failed to retrieve location",
			"success": false,
		})
	}

	if err := db.Database.Db.Preload("Children").Find(&location, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   "Location not found",
				"Success": false,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrieve Location",
			"success": false,
		})

	}
	return c.Status(fiber.StatusOK).JSON(location)
}

// update location

func UpdateLocation(c *fiber.Ctx) error {
	id := c.Params("id")

	var location models.Location

	if db.Database.Db.First(&location, id).Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrive Location",
			"success": false,
		})

	}

	if err := db.Database.Db.First(&location, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"errr":    "location not found",
				"success": false,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update location",
		})
	}

	if err := c.BodyParser(&location); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid body request",
			"success": false,
		})
	}

	if err := db.Database.Db.Save(&location).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "failed to update location",
			"success": false,
		})
	}
	return c.Status(fiber.StatusOK).JSON(location)
}

// delete location

func DeleteLocation(c *fiber.Ctx) error {
	id := c.Params("id")

	var location models.Location

	if db.Database.Db.First(&location, id).Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrive location",
			"success": false,
		})
	}

	if err := db.Database.Db.First(&location, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   "location not found",
				"Success": false,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrive Location",
			"success": false,
		})
	}

	if err := db.Database.Db.Delete(&location).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to delete location",
			"success": false,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Location deleted successfully",
	})
}
