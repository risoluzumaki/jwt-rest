package auth

import (
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	services AuthServiceInterface
}

func NewAuthHandler(s AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		services: s,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req CreateUserDTO
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		},
		)
	}

	email := req.Email
	password := req.Password

	token, err := h.services.LoginUser(email, password)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(200).JSON(&ResponseUserDTO{
		Token: token,
	})

}
