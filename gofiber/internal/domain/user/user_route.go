package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/risoluzumaki/jwt/go/pkg/middleware"
)

func UserRoute(r fiber.Router, h *UserHandler) {

	r.Use(middleware.AuthMiddleware)
	r.Get("/profile", h.GetProfile)
}
