package routes

import (
	handlers "miniproject/handlers"

	"github.com/gofiber/fiber/v2"
)

func APIRoutes(app *fiber.App) {
    app.Post("/api/login", handlers.LoginUser)
    app.Post("/api/register", handlers.RegisterUser)
    
    app.Post("/api/alamat", handlers.CreateAlamat)

    app.Post("/api/produk", handlers.CreateProduk)
	app.Get("/api/produk/:idToko", handlers.GetProdukByTokoID)
	
    app.Post("/categories", handlers.CreateCategory)
    app.Get("/categories", handlers.GetCategories)

    app.Post("/trx", handlers.CreateTrx)
    app.Post("/detail-trx", handlers.CreateDetailTrx)

    app.Post("/log-produk", handlers.CreateLogProduk)

    app.Post("/foto-produk", handlers.CreateFotoProduk)

    // app.Get("/api/alamat", handlers.GetAlamatByUserID)
	// app.Get("/api/v1/toko", handlers.GetTokoData)
}
