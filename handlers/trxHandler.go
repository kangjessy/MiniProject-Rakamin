package handlers

import (
	"miniproject/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTrx(c *fiber.Ctx) error {
	var newTrx models.Trx

	if err := c.BodyParser(&newTrx); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

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