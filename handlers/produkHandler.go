package handlers

import (
	"miniproject/models"

	"github.com/gofiber/fiber/v2"
)

// Handler untuk membuat produk baru
func CreateProduk(c *fiber.Ctx) error {
	var newProduk models.Produk

	// Baca data produk baru dari body permintaan
	if err := c.BodyParser(&newProduk); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Simpan produk baru ke database
	if err := db.Create(&newProduk).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create produk",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Produk created successfully",
		"produk":  newProduk,
	})
}

// Handler untuk mendapatkan daftar produk berdasarkan IDToko
func GetProdukByTokoID(c *fiber.Ctx) error {
	// Ambil IDToko dari parameter URL
	idToko := c.Params("idToko")

	// Cari daftar produk berdasarkan IDToko
	var produkList []models.Produk
	if err := db.Where("id_toko = ?", idToko).Find(&produkList).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get produk list",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Produk list retrieved successfully",
		"produk":  produkList,
	})
}