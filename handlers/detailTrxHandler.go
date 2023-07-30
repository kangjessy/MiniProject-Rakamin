package handlers

import (
	"miniproject/models"

	"github.com/gofiber/fiber/v2"
)

func CreateDetailTrx(c *fiber.Ctx) error {
	var newDetailTrx models.DetailTrx

	if err := c.BodyParser(&newDetailTrx); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	result := db.Create(&newDetailTrx)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create detail transaksi",
		})
	}

	return c.JSON(fiber.Map{
		"message":     "Detail transaksi created successfully",
		"detail_trx": newDetailTrx,
	})
}
