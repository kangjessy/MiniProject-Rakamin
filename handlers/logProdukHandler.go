package handlers

import (
	"miniproject/models"

	"github.com/gofiber/fiber/v2"
)

func CreateLogProduk(c *fiber.Ctx) error {
	var newLogProduk models.LogProduk

	if err := c.BodyParser(&newLogProduk); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	result := db.Create(&newLogProduk)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create log produk",
		})
	}

	return c.JSON(fiber.Map{
		"message":   "Log produk created successfully",
		"log_produk": newLogProduk,
	})
}
