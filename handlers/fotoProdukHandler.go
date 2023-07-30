package handlers

import (
	"miniproject/models"

	"github.com/gofiber/fiber/v2"
)

func CreateFotoProduk(c *fiber.Ctx) error {
	var newFotoProduk models.FotoProduk

	if err := c.BodyParser(&newFotoProduk); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	result := db.Create(&newFotoProduk)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create foto produk",
		})
	}

	return c.JSON(fiber.Map{
		"message":    "Foto produk created successfully",
		"foto_produk": newFotoProduk,
	})
}