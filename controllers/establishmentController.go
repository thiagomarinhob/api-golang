package controllers

import (
	"net/http"

	"go-api/database"
	"go-api/models"

	"github.com/gin-gonic/gin"
)

func CreateEstablishment(c *gin.Context) {
	var establishment models.Establishment
	if err := c.ShouldBindJSON(&establishment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Conexão com o banco de dados não inicializada"})
		return
	}

	// Validar se o AdminUserID (id do usuário) existe
	var user models.User
	if err := database.DB.First(&user, establishment.AdminUserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Se o usuário for válido, criamos o estabelecimento
	if err := database.DB.Create(&establishment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, establishment)
}

func GetEstablishments(c *gin.Context) {
	var establishments []models.Establishment
	if err := database.DB.Find(&establishments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, establishments)
}

func GetEstablishmentByID(c *gin.Context) {
	id := c.Param("id")
	var establishments models.Establishment
	if err := database.DB.Find(&establishments, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Estabelecimento não encontrado": err.Error()})
		return
	}

	c.JSON(http.StatusOK, establishments)
}

func UpdateEstablishment(c *gin.Context) {
	id := c.Param("id")
	var establishments models.Establishment

	if err := database.DB.Find(&establishments, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Estabelecimento não encontrado": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&establishments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, establishments)
}

func DeleteEstablishment(c *gin.Context) {
	id := c.Param("id")
	var establishments []models.Establishment

	if err := database.DB.Find(&establishments, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Estabelecimento não encontrado": err.Error()})
		return
	}

	if err := database.DB.Delete(&establishments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estabelecimenti excluído com suceeso"})
}
