package middlewares

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

// use for check service.
// not use yet.

func HealthCheck(app *fiber.App) {
	// Provide a minimal config
	app.Use(healthcheck.New())

	// Or extend your config for customization
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/live",
		// ReadinessProbe: func(c *fiber.Ctx) bool {
		// 	return serviceA.Ready() && ...
		// },
		ReadinessEndpoint: "/ready",
	}))

}
