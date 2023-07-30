package handlers

import (
	"miniproject/models"

	"github.com/gofiber/fiber/v2"
)

func CreateCategory(c *fiber.Ctx) error {
	var newCategory models.Category

	if err := c.BodyParser(&newCategory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := db.Create(&newCategory).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create category",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Category created successfully",
		"category":  newCategory,
	})
}

func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category

	if err := db.Find(&categories).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get categories",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Categories retrieved successfully",
		"categories":  categories,
	})
}
