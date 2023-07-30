// middleware/logging.go
package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	// Catat waktu permintaan
	start := time.Now()

	// Lanjutkan permintaan ke handler selanjutnya
	err := c.Next()

	// Catat waktu selesai permintaan
	end := time.Now()

	// Cetak log
	fmt.Printf("Request: %s %s | Status: %d | Duration: %v\n", c.Method(), c.Path(), c.Response().StatusCode(), end.Sub(start))

	// Lanjutkan error jika ada
	return err
}
