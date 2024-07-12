package routes

import (
	"github.com/Pratchaya0/auth-api-gofiber-jwt/entities"
	"github.com/Pratchaya0/auth-api-gofiber-jwt/interfaces/handlers"
	"github.com/Pratchaya0/auth-api-gofiber-jwt/middlewares"
	"github.com/Pratchaya0/auth-api-gofiber-jwt/repositories"
	"github.com/Pratchaya0/auth-api-gofiber-jwt/usecases"
	"github.com/gofiber/fiber/v2"
)

func UserRouteSetup(app *fiber.App) {

	userRepo := repositories.NewGormUserRepository(entities.DB())
	userUseCase := usecases.NewUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(userUseCase)

	// Private controller
	userRoute := app.Group("/users")
	{
		userRoute.Use(middlewares.Authorize)
		{
			userRoute.Get("/list", userHandler.List)
		}
	}
}
