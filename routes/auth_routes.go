package routes

import (
	// controllers "github.com/Pratchaya0/auth-api-gofiber-jwt/controllers"
	"github.com/Pratchaya0/auth-api-gofiber-jwt/entities"
	"github.com/Pratchaya0/auth-api-gofiber-jwt/interfaces/handlers"
	"github.com/Pratchaya0/auth-api-gofiber-jwt/repositories"
	"github.com/Pratchaya0/auth-api-gofiber-jwt/usecases"
	"github.com/gofiber/fiber/v2"
)

func AuthRouteSetup(app *fiber.App) {

	userRepo := repositories.NewGormUserRepository(entities.DB())
	userUseCase := usecases.NewUserUseCase(userRepo)
	authHandler := handlers.NewAuthHandler(userUseCase)

	authRoute := app.Group("/auth")
	{
		authRoute.Get("/current-user", authHandler.CurrentUser)
		authRoute.Post("/register", authHandler.Register)
		authRoute.Post("/login", authHandler.Login)
		authRoute.Post("/logout", authHandler.Logout)
	}
}
