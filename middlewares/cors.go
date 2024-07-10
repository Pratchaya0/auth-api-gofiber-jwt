package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORSMiddleware(app *fiber.App) {
	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     "*",
			AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
			AllowCredentials: false,
			AllowMethods:     "POST, OPTIONS, GET, PUT",
		}))
}
