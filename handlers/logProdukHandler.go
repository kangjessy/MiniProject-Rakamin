// handlers/log_produk_handler.go
package handlers

import (
	"miniproject/models"

	"github.com/gofiber/fiber/v2"
)

// Handler untuk menambahkan log produk baru
func CreateLogProduk(c *fiber.Ctx) error {
	var newLogProduk models.LogProduk

	// Baca data log produk baru dari body permintaan
	if err := c.BodyParser(&newLogProduk); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Simpan log produk baru ke dalam database
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
