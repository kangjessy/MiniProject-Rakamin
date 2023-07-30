package main

import (
	handlers "miniproject/handlers"
	"miniproject/models"
	"miniproject/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main(){
	// Menghubungkan ke database
	dsn := "root:@tcp(127.0.0.1:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// Mengecek apabila koneksi gagal
	if err != nil{
		panic("Failed to connect to database")
	}

	// Memanggil fungsi AutoMigrate
	db.AutoMigrate(
		&models.User{}, 
		&models.Toko{}, 
		&models.Produk{}, 
		&models.LogProduk{},
		&models.Alamat{}, 
		&models.Category{},
		&models.Trx{},
		&models.DetailTrx{},
		&models.FotoProduk{},
	)
	if err != nil {
		panic("Failed to run database migration")
	}

	// Memanggil fungsi Inisialiasi DB
	handlers.InitDB(db)

	// Fungsi fiber
	app := fiber.New()

	// Menambahkan custom middleware untuk logging
	// app.Use(middleware.LoggingMiddleware)

	routes.APIRoutes(app)
	routes.WebRoutes(app)

	// Endpoint untuk pesan selamat datang dan informasi login
	app.Post("/api/v1", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Selamat Datang! Silakan login atau daftar untuk melanjutkan.",
		})
	})

	app.Static("/", "./views", fiber.Static{
		Index: "index.html",
	})

	err = app.Listen(":3000")
	if err != nil{
		panic("Failed to start Fiber server")
	}
}