// handlers/category_handler.go
package handlers

import (
	"miniproject/models"

	"github.com/gofiber/fiber/v2"
)

// Handler untuk membuat kategori baru
func CreateCategory(c *fiber.Ctx) error {
	var newCategory models.Category

	// Baca data kategori baru dari body permintaan
	if err := c.BodyParser(&newCategory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Simpan kategori baru ke database
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

// Handler untuk mendapatkan daftar kategori
func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category

	// Cari semua kategori dari database
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
