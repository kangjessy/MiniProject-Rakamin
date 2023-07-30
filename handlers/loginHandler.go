package handlers

import (
	"miniproject/models"

	"github.com/gofiber/fiber/v2"
)

func LoginUser(c *fiber.Ctx) error {
    var loginData struct {
		Email    string `json:"email"`
		Katasandi string `json:"katasandi"`
	}
    if err := c.BodyParser(&loginData); err != nil {
		return err
	}

    // Cari user berdasarkan email
	var user models.User
	if err := db.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		return fiber.ErrUnauthorized
	}

    // Periksa password
	if user.KataSandi != loginData.Katasandi {
		return fiber.ErrUnauthorized
	}

    // Jika login berhasil, buat token JWT
	// token, err := utils.CreateToken(user.ID)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": "Failed to create token",
	// 	})
	// }

    // mencari Toko berdasarkan ID
    // var toko models.Toko
    // if err := db.Where("id_user = ?", user.ID).First(&toko).Error; err !=nil{
    //     return fiber.ErrInternalServerError
    // }

    return c.JSON(fiber.Map{
        "message": "Login berhasil",
        // "token": token,
        "user": user,
        // "toko": toko,
    })
}

