package handlers

import (
	"miniproject/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Handler untuk mengubah data toko
func UpdateToko(c *fiber.Ctx, db *gorm.DB) error {
	// Ambil data user yang telah diotentikasi dari context
	user := c.Locals("user").(models.User)

	// Baca data toko yang ingin diubah dari body permintaan
	var tokoData models.Toko
	if err := c.BodyParser(&tokoData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Pastikan hanya pemilik toko yang dapat mengubah data toko
	if tokoData.IDUser != user.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbidden. You are not the owner of this toko.",
		})
	}

	// Cari data toko terkini dari database berdasarkan IDUser
	var currentToko models.Toko
	if err := db.Where("id_user = ?", user.ID).First(&currentToko).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get current toko data",
		})
	}

	// Lakukan perubahan yang diperlukan pada data toko
	currentToko.NamaToko = tokoData.NamaToko
	currentToko.UrlFoto = tokoData.UrlFoto

	// Simpan perubahan ke dalam database
	if err := db.Save(&currentToko).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update toko data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Toko data updated successfully",
		"toko":    currentToko,
	})
}
