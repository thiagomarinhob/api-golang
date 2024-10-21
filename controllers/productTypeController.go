package controllers

import (
	"go-api/database"
	"go-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateProductType(c *gin.Context) {
	var productType models.ProductType
	establishmentID := c.GetHeader("Establishment-ID")

	// Bind JSON body para o modelo ProductType
	if err := c.ShouldBindJSON(&productType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productType.EstablishmentID = establishmentID

	// Verifique se models.DB não é nil
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Conexão com o banco de dados não inicializada"})
		return
	}

	var establishment models.Establishment
	if err := database.DB.First(&establishment, "id = ?", productType.EstablishmentID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Establishment not found"})
		return
	}

	productType.Establishment = establishment

	// Criar o tipo de produto no banco de dados
	if err := database.DB.Create(&productType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, productType)
}

func GetProductTypes(c *gin.Context) {
	var productTypes []models.ProductType
	establishmentID := c.GetHeader("Establishment-ID") // Pegando o Establishment-ID do cabeçalho

	// Verificar se o EstablishmentID é um UUID válido
	if _, err := uuid.Parse(establishmentID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Establishment ID"})
		return
	}

	// Verificar se o estabelecimento existe
	var establishment models.Establishment
	if err := database.DB.First(&establishment, "id = ?", establishmentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Establishment not found"})
		return
	}

	// Buscar todos os tipos de produtos para o estabelecimento
	if err := database.DB.Where("establishment_id = ?", establishmentID).Preload("Products").Find(&productTypes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, productTypes)
}

func GetProductTypeByID(c *gin.Context) {
	var productType models.ProductType
	productTypeID := c.Param("id")
	establishmentID := c.GetHeader("Establishment-ID")

	// Verificar se o EstablishmentID é um UUID válido
	if _, err := uuid.Parse(establishmentID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Establishment ID"})
		return
	}

	// Verificar se o estabelecimento existe
	var establishment models.Establishment
	if err := database.DB.First(&establishment, "id = ?", establishmentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Establishment not found"})
		return
	}

	// Buscar o tipo de produto pelo ID e pelo estabelecimento
	if err := database.DB.Where("id = ? AND establishment_id = ?", productTypeID, establishmentID).
		Preload("Products").First(&productType).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product type not found"})
		return
	}

	c.JSON(http.StatusOK, productType)
}

func UpdateProductType(c *gin.Context) {
	var productType models.ProductType
	productTypeID := c.Param("id")
	establishmentID := c.GetHeader("Establishment-ID")

	// Verificar se o EstablishmentID é um UUID válido
	if _, err := uuid.Parse(establishmentID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Establishment ID"})
		return
	}

	// Verificar se o estabelecimento existe
	var establishment models.Establishment
	if err := database.DB.First(&establishment, "id = ?", establishmentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Establishment not found"})
		return
	}

	// Buscar o tipo de produto pelo ID e verificar se pertence ao estabelecimento
	if err := database.DB.Where("id = ? AND establishment_id = ?", productTypeID, establishmentID).First(&productType).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product type not found"})
		return
	}

	// Bind JSON body para o modelo ProductType
	if err := c.ShouldBindJSON(&productType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Atualizar o tipo de produto no banco de dados
	if err := database.DB.Save(&productType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, productType)
}

func DeleteProductType(c *gin.Context) {
	productTypeID := c.Param("id")
	establishmentID := c.GetHeader("Establishment-ID")

	// Verificar se o EstablishmentID é um UUID válido
	if _, err := uuid.Parse(establishmentID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Establishment ID"})
		return
	}

	// Verificar se o estabelecimento existe
	var establishment models.Establishment
	if err := database.DB.First(&establishment, "id = ?", establishmentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Establishment not found"})
		return
	}

	// Verificar se o tipo de produto existe e pertence ao estabelecimento antes de deletar
	if err := database.DB.Where("id = ? AND establishment_id = ?", productTypeID, establishmentID).Delete(&models.ProductType{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product type deleted successfully"})
}
