package handlers

import (
	"miniproject/models"

	"github.com/gofiber/fiber/v2"
)

// Handler untuk menambahkan foto produk baru
func CreateFotoProduk(c *fiber.Ctx) error {
	var newFotoProduk models.FotoProduk

	// Baca data foto produk baru dari body permintaan
	if err := c.BodyParser(&newFotoProduk); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Simpan foto produk baru ke dalam database
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