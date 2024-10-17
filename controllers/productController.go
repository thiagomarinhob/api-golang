package controllers

import (
	"net/http"

	"go-api/database"
	"go-api/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	establishmentID := c.GetHeader("Establishment-ID")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.EstablishmentID = establishmentID

	// Verifique se models.DB não é nil
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Conexão com o banco de dados não inicializada"})
		return
	}

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func GetProducts(c *gin.Context) {
	establishmentID := c.GetHeader("Establishment-ID")
	var products []models.Product
	if err := database.DB.Where("establishment_id = ?", establishmentID).Preload("ProductType").Preload("ProductPhoto").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	var product models.Product
	productID := c.Param("id")
	establishmentID := c.GetHeader("Establishment-ID") // Assumindo que o Establishment-ID vem no cabeçalho

	// Buscar o produto pelo ID e verificar se pertence ao estabelecimento
	if err := database.DB.Where("id = ? AND establishment_id = ?", productID, establishmentID).
		Preload("ProductType").
		Preload("ProductPhoto").
		First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	productID := c.Param("id")
	establishmentID := c.GetHeader("Establishment-ID") // Assumindo que o Establishment-ID vem no cabeçalho

	// Buscar o produto pelo ID e verificar se pertence ao estabelecimento
	if err := database.DB.Where("id = ? AND establishment_id = ?", productID, establishmentID).
		First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Bind JSON body para o modelo Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Atualizar o produto no banco de dados
	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	productID := c.Param("id")
	establishmentID := c.GetHeader("Establishment-ID") // Assumindo que o Establishment-ID vem no cabeçalho

	// Verificar se o produto pertence ao estabelecimento e excluí-lo
	if err := database.DB.Where("id = ? AND establishment_id = ?", productID, establishmentID).
		Delete(&models.Product{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto excluído com sucesso"})
}
