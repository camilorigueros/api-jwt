package controller

import (
	"github.com/camilorigueros/api-jwt/config"
	service2 "github.com/camilorigueros/api-jwt/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWTHandler define la interfaz pública para manejar JWT.
type JWTHandler interface {
	GenerateToken(username string) (string, error)
	ValidateToken(token string) (string, error)
}

// jwtHandler es la implementación de JWTHandler que usa el servicio interno.
type jwtHandler struct {
	jwtService *service2.JWTService
}

// NewJWTHandler crea una nueva instancia de jwtHandler.
func NewJWTHandler() JWTHandler {
	return &jwtHandler{
		jwtService: service2.NewJWTService(config.AppConfig.Security.JWT),
	}
}

// GenerateToken implementa la generación de tokens.
func (j *jwtHandler) GenerateToken(username string) (string, error) {
	return j.jwtService.GenerateToken(username)
}

// ValidateToken implementa la validación de tokens.
func (j *jwtHandler) ValidateToken(token string) (string, error) {
	return j.jwtService.ValidateToken(token)
}

// Login es un controlador público para autenticar y generar un token JWT.
func Login(handler JWTHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		// Crear instancia del AuthService
		authService := service2.NewAuthService()

		// Validar credenciales
		err := authService.Authenticate(username, password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Generar token JWT
		token, err := handler.GenerateToken(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

// Validate es un controlador público para validar un token JWT.
func Validate(handler JWTHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")

		// Validar el token utilizando el manejador
		username, err := handler.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Token válido", "username": username})
	}
}
