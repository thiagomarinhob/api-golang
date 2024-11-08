package controllers

import (
	"net/http"

	"go-api/database"
	"go-api/models"

	"github.com/gin-gonic/gin"
)

func CreateClient(c *gin.Context) {
	var client models.Client

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// // Verifique se models.DB não é nil
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Conexão com o banco de dados não inicializada"})
		return
	}

	if err := database.DB.Create(&client).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}