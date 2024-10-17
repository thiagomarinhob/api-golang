package routes

import (
	"go-api/controllers"
	"go-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Rotas públicas
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.Login)
	router.POST("/refresh", controllers.RefreshToken)

	// Grupo de rotas protegidas
	auth := router.Group("/")
	auth.Use(middlewares.Auth())
	{
		// Rotas de produtos
		auth.POST("/products", controllers.CreateProduct)
		auth.GET("/products", controllers.GetProducts)
		auth.GET("/products/:id", controllers.GetProductByID)
		auth.PUT("/products/:id", controllers.UpdateProduct)
		auth.DELETE("/products/:id", controllers.DeleteProduct)
	}
}
