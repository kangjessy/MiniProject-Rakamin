// handlers/toko_handler.go
package handlers

import (
	"miniproject/models"

	"github.com/gofiber/fiber/v2"
)

// Handler untuk mendapatkan data toko pengguna yang sudah login
func GetTokoData(c *fiber.Ctx) error {
	// Ambil data user dari context
	user := c.Locals("user").(models.User)

	// Query data toko berdasarkan ID user
	var toko models.Toko
	if err := db.Where("id_user = ?", user.ID).First(&toko).Error; err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(toko)
}
