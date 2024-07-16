package handlers

import (
	"github.com/Pratchaya0/auth-api-gofiber-jwt/DTOs/responses"
	"github.com/Pratchaya0/auth-api-gofiber-jwt/entities"
	"github.com/Pratchaya0/auth-api-gofiber-jwt/usecases"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUseCase *usecases.UserUseCase
}

func NewUserHandler(userUseCase *usecases.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (h *UserHandler) Save(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	user, err := h.userUseCase.Save(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": user})
}

// @Summary แสดงข้อมูลผู้ใช้งานทั้งหมดในระบบ
// @Tag User
// @Security bearerToken
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /users/list [get]
func (h *UserHandler) List(c *fiber.Ctx) error {
	users, err := h.userUseCase.List()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	webResponse := responses.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   users,
	}

	return c.Status(fiber.StatusOK).JSON(webResponse)
}
