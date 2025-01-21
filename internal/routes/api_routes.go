package routes

import (
	"github.com/camilorigueros/api-jwt/internal/controller"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas las rutas de la API
func SetupRoutes(router *gin.Engine, jwtHandler controller.JWTHandler) {
	// Rutas de autenticación
	authRoutes := router.Group("/api")
	{
		authRoutes.POST("/login", controller.Login(jwtHandler))
		authRoutes.GET("/validate", controller.Validate(jwtHandler))
	}
}
