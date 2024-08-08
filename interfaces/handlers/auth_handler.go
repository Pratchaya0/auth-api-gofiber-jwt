package handlers

import (
	"os"
	"strconv"
	"time"

	"github.com/Pratchaya0/auth-api-gofiber-jwt/dtos/responses"
	"github.com/Pratchaya0/auth-api-gofiber-jwt/entities"
	"github.com/Pratchaya0/auth-api-gofiber-jwt/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userUseCase *usecases.UserUseCase
}

func NewAuthHandler(userUseCase *usecases.UserUseCase) *AuthHandler {
	return &AuthHandler{userUseCase: userUseCase}
}

// @Summary แสดงข้อมูลของผู้ใช้
// @Tag Auth
// @Security bearerToken
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /auth/current-user [get]
func (h *AuthHandler) CurrentUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("AUTH_SECRET_KEY")), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	user, err := h.userUseCase.GetUserById(claims.Issuer)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found."})
	}

	webResponse := responses.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   user,
	}

	return c.Status(fiber.StatusOK).JSON(webResponse)
}

// @Summary สมัครเข้าใช้งานระบบ
// @Tag Auth
// @Param register query requests.RegisterRequest true "RegisterRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := entities.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	user, err := h.userUseCase.Save(user)
	if err != nil {
		return err
	}

	webResponse := responses.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   user,
	}

	return c.Status(fiber.StatusOK).JSON(webResponse)
}

// @Summary ลงชื่อเข้าใช้งานระบบ
// @Tag Auth
// @Param loginRequest query requests.LoginRequest true "LoginRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user, err := h.userUseCase.GetUserByEmail(data["email"])
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found."})
	}

	if user.ID == 0 { //If the ID return is '0' then there is no such email present in the DB
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("AUTH_SECRET_KEY")))

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	webResponse := responses.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data: fiber.Map{
			"message": "login success",
			"token":   "Bearer " + token,
		},
	}

	return c.Status(fiber.StatusOK).JSON(webResponse)
}

// @Summary ออกจากการใช้งานระบบ
// @Tag Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), //Sets the expiry time an hour ago in the past.
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	webResponse := responses.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data: fiber.Map{
			"message": "logout success",
		},
	}

	return c.Status(fiber.StatusOK).JSON(webResponse)
}
