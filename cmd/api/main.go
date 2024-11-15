package main

import (
	"go-api/config"
	"go-api/database"
	"go-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"time"
)

func main() {
	// Carregar as configurações
	config.LoadConfig()

	// Conectar ao banco de dados
	database.ConnectDatabase()

	// Iniciar o roteador
	server := gin.Default()

	server.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"}, // Domínios permitidos
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},     // Métodos permitidos
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Establishment-ID"},    // Cabeçalhos permitidos
        ExposeHeaders:    []string{"Content-Length"},                             // Cabeçalhos expostos
        AllowCredentials: true,                                                  // Permitir credenciais (cookies)
        MaxAge:           12 * time.Hour,                                        // Tempo de cache do CORS
    }))

	// Configurar as rotas
	routes.SetupRoutes(server)

	// Iniciar o servidor
	server.Run(":8080")
}
