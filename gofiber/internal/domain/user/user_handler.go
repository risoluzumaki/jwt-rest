package user

import (
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service UserServiceInterface
}

func NewUserHandler(s UserServiceInterface) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	id := c.Locals("user_id").(int)
	res, err := h.service.GetUserProfile(id)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(&ResponseUser{
		ID:       res.ID,
		Name:     res.Name,
		Email:    res.Email,
		Username: res.Username,
	})
}
