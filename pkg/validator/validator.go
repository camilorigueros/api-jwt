package validator

import (
	"github.com/camilorigueros/api-jwt/config"
	"github.com/golang-jwt/jwt/v4"
)

type JWTService struct {
	secret string
}

func NewJWTService(cfg config.JWTConfig) *JWTService {
	return &JWTService{
		secret: cfg.Secret,
	}
}
func (s *JWTService) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", jwt.ErrTokenInvalidId
	}

	username, ok := claims["sub"].(string)
	if !ok {
		return "", jwt.ErrInvalidKey
	}

	return username, nil
}
