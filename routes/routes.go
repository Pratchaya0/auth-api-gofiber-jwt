package routes

import (
	controllers "github.com/Pratchaya0/auth-api-gofiber-jwt/controllers"
	users "github.com/Pratchaya0/auth-api-gofiber-jwt/controllers/users"
	"github.com/Pratchaya0/auth-api-gofiber-jwt/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RouteSetup(app *fiber.App) {

	authRoute := app.Group("/auth")
	{
		authRoute.Get("/currentUser", controllers.CurrentUser)
		authRoute.Post("/register", controllers.Register)
		authRoute.Post("/login", controllers.Login)
		authRoute.Post("/logout", controllers.Logout)
	}

	// Private controller
	userRoute := app.Group("/users")
	{
		userRoute.Use(middlewares.Authorize)
		{
			userRoute.Get("/hello-world", users.HelloWorld)
			userRoute.Get("/list", users.List)
			userRoute.Get("/select", users.Select)
		}
	}
}
