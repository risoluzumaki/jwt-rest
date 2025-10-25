package internal

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	// "github.com/risoluzumaki/jwt/go/internal/domain/auth"
	"github.com/risoluzumaki/jwt/go/internal/domain/auth"
	"github.com/risoluzumaki/jwt/go/internal/domain/user"
	"github.com/risoluzumaki/jwt/go/pkg/app"
)

func NewApp() *fiber.App {

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Default err
			code := fiber.StatusInternalServerError
			msg := err.Error()

			// cek AppError
			var aErr *app.AppError
			if ok := errors.As(err, &aErr); ok {
				code = aErr.Code
				msg = aErr.Msg
			}

			// cek Fiber error
			if fErr, ok := err.(*fiber.Error); ok {
				code = fErr.Code
				msg = fErr.Message
			}

			return c.Status(code).JSON(fiber.Map{
				"code":  code,
				"error": msg,
			})
		},
	})

	// GLOBAL MIDDLEWARE
	// Recover
	app.Use(recover.New())
	// logging
	// app.Use()

	// Wiring dependency
	// User
	userRepository := user.NewUserRepository() // Sharing depend

	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	// Auth
	authService := auth.NewAuthService(userRepository) // inject sharing depend on user repo
	authHandler := auth.NewAuthHandler(authService)

	// Route grouping
	api := app.Group("/api/v1")
	authRouterGrup := api.Group("/auth")
	userRouterGroup := api.Group("/user")

	auth.AuthRoute(authRouterGrup, authHandler)
	user.UserRoute(userRouterGroup, userHandler)

	return app
}
