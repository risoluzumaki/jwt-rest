package auth

import (
	"github.com/risoluzumaki/jwt/go/internal/domain/user"
	appErr "github.com/risoluzumaki/jwt/go/pkg/app"
	appUtils "github.com/risoluzumaki/jwt/go/pkg/utils"
)

type AuthServiceInterface interface {
	LoginUser(email string, password string) (string, error)
}

type AuthService struct {
	repo user.UserRepositoryInterface
}

func NewAuthService(ur user.UserRepositoryInterface) AuthServiceInterface {
	return &AuthService{
		repo: ur,
	}
}

func (s *AuthService) LoginUser(email string, password string) (string, error) {
	result, err := s.repo.GetByEmail(email)

	if err != nil {
		return "", appErr.NewError(404, "user not found")
	}

	if result.Password != password {
		return "", appErr.NewError(400, "Invalid Password")
	}

	token, err := appUtils.GenerateToken(result.ID, result.Email)
	if err != nil {
		return "", appErr.NewError(500, "Internal Server Error")
	}

	return token, nil
}
