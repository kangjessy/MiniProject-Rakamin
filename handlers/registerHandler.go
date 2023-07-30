// handlers/user_handler.go
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

// Handler untuk menambahkan pengguna baru
func RegisterUser(c *fiber.Ctx) error {
    var newUser models.User

    // Baca data dari body permintaan dan bind ke struct newUser
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

    // Simpan pengguna baru ke database
    result := db.Create(&newUser)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to create user",
        })
    }

    // Setelah pengguna berhasil dibuat, buat toko baru untuk pengguna ini
	newToko := models.Toko{
		IDUser: newUser.ID,
        NamaToko: newUser.Nama + "'s Toko",
		// Tambahkan field lain yang relevan untuk toko baru
	}

	// Simpan toko baru ke database
	result = db.Create(&newToko)
	if result.Error != nil {
		// Jika ada masalah dalam membuat toko, Anda dapat memutuskan bagaimana menanganinya,
		// misalnya, menghapus user yang sudah dibuat sebelumnya atau memberikan respons error.
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


