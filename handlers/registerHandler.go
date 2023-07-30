package handlers

import (
	"miniproject/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(database *gorm.DB){
	db = database
}

func RegisterUser(c *fiber.Ctx) error {
    var newUser models.User

    if err := c.BodyParser(&newUser); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid request body",
        })
    }

	// Periksa apakah email sudah digunakan atau belum
	var existingUser models.User
    if err := db.Where("email = ?", newUser.Email).First(&existingUser).Error; err == nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Email sudah digunakan",
        })
    }

	// Periksa apakah no telepon sudah digunakan atau belum
	var notelpUser models.User
	if err := db.Where("notelp = ?", newUser.Notelp).First(&notelpUser).Error; err == nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "No telepon sudah digunakan",
		})
	}

	// Set nilai default IsAdmin menjadi false
    newUser.IsAdmin = true

    // Simpan user baru ke database
    result := db.Create(&newUser)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to create user",
        })
    }

    // Membuat toko setelah register
	newToko := models.Toko{
		IDUser: newUser.ID,
        NamaToko: newUser.Nama + "'s Toko",
	}

	// save to database
	result = db.Create(&newToko)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create toko",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Selamat! Anda telah mendaftar!",
		"user":    newUser,
		"toko":    newToko,
	})
}


