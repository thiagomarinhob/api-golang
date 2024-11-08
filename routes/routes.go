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
	router.POST("/clients", controllers.CreateClient)

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

		auth.POST("/establishment", controllers.CreateEstablishment)
		auth.GET("/establishment", controllers.GetEstablishments)
		auth.GET("/establishment/:id", controllers.GetEstablishmentByID)
		auth.PUT("/establishment/:id", controllers.UpdateEstablishment)
		auth.DELETE("/establishment/:id", controllers.DeleteEstablishment)

		auth.GET("/product-types", controllers.GetProductTypes)
		auth.GET("/product-types/:id", controllers.GetProductTypeByID)
		auth.POST("/product-types", controllers.CreateProductType)
		auth.PUT("/product-types/:id", controllers.UpdateProductType)
		auth.DELETE("/product-types/:id", controllers.DeleteProductType)

		auth.POST("/order", controllers.CreateOrder)
	}
}
