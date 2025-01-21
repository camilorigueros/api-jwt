package main

import (
	"github.com/camilorigueros/api-jwt/config"
	"github.com/camilorigueros/api-jwt/internal/controller"
	"github.com/camilorigueros/api-jwt/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Cargar configuración
	config.LoadConfig()

	// Inicializar controlador
	jwtHandler := controller.NewJWTHandler()

	// Crear la instancia del servidor
	r := gin.Default()

	// Configurar las rutas de la API
	routes.SetupRoutes(r, jwtHandler)

	// Iniciar servidor en el puerto 8080
	r.Run(":8080")
}
