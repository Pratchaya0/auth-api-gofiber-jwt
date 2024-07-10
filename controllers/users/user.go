package users

import (
	"net/http"

	"github.com/Pratchaya0/auth-api-gofiber-jwt/models"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/log"
)

// Handler functions
// helloWorld godoc
// @Summary Test Swagger with hello world
// @Description Return Hello world
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /hello-world [get]
func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func List(c *fiber.Ctx) error {
	var users []models.User

	// get list user
	if err := models.DB().Raw("SELECT id, created_at, updated_at, deleted_at, name, email, password FROM users;").Scan(&users).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(users)
}

func Select(c *fiber.Ctx) error {
	var user models.User
	var params map[string]int

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if tx := models.DB().Where("id = ?", params["id"]).First(&user); tx.RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found."})
	}

	log.Info(user)

	return c.Status(http.StatusOK).JSON(user)
}
