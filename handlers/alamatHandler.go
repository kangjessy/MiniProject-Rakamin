package handlers

import (
	"miniproject/models"

	"github.com/gofiber/fiber/v2"
)

// Handler untuk membuat alamat baru
func CreateAlamat(c *fiber.Ctx) error {
	var newAlamat models.Alamat

	// Baca data alamat baru dari body permintaan
	if err := c.BodyParser(&newAlamat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Simpan alamat baru ke database
	result := db.Create(&newAlamat)
	if result.Error != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create alamat",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Alamat created successfully",
		"alamat":  newAlamat,
	})
}

// Handler untuk mendapatkan daftar alamat berdasarkan IDUser
// func GetAlamatByUserID(c *fiber.Ctx) error {
// 	// Ambil data user yang telah diotentikasi dari context
// 	user := c.Locals("user").(models.User)

// 	// Cari daftar alamat berdasarkan IDUser
// 	var alamatList []models.Alamat
// 	if err := db.Where("id_user = ?", user.ID).Find(&alamatList).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Failed to get alamat list",
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Alamat list retrieved successfully",
// 		"alamat":  alamatList,
// 	})
// }