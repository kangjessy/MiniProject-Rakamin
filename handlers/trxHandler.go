package handlers

import (
	"miniproject/models"

	"github.com/gofiber/fiber/v2"
)

// Handler untuk membuat transaksi baru
func CreateTrx(c *fiber.Ctx) error {
	var newTrx models.Trx

	// Baca data transaksi baru dari body permintaan
	if err := c.BodyParser(&newTrx); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Simpan transaksi baru ke dalam database
	result := db.Create(&newTrx)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create transaksi",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Transaksi created successfully",
		"trx":     newTrx,
	})
}