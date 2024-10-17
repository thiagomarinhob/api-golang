package main

import (
	"go-api/config"
	"go-api/database"
	"go-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Carregar as configurações
	config.LoadConfig()

	// Conectar ao banco de dados
	database.ConnectDatabase()

	// Iniciar o roteador
	server := gin.Default()

	// Configurar as rotas
	routes.SetupRoutes(server)

	// Iniciar o servidor
	server.Run(":8080")
}
