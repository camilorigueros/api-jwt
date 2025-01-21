package service

import (
	"errors"
	"github.com/camilorigueros/api-jwt/config"
)

// AuthService gestiona la autenticación
type AuthService struct {
	users map[string]string
}

func NewAuthService() *AuthService {
	return &AuthService{
		users: config.AppConfig.Security.Users,
	}
}

func (s *AuthService) Authenticate(username, password string) error {
	// Validar las credenciales del usuario
	if storedPassword, exists := s.users[username]; !exists || storedPassword != password {
		return errors.New("credenciales inválidas")
	}
	return nil
}
