package service

import (
	"github.com/camilorigueros/api-jwt/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService struct {
	secret     string
	expiration time.Duration
	issuer     string
}

func NewJWTService(cfg config.JWTConfig) *JWTService {
	return &JWTService{
		secret:     cfg.Secret,
		expiration: time.Duration(cfg.Expiration) * time.Second,
		issuer:     cfg.Issuer,
	}
}

func (s *JWTService) GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"sub": username,
		"iss": s.issuer,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(s.expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
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
