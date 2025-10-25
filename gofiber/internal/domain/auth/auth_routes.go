package auth

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(r fiber.Router, h *AuthHandler) {
	r.Post("/login", h.Login)
}
