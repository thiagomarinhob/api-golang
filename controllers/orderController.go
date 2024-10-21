package controllers

import (
	"net/http"
	"time"

	"go-api/database"
	"go-api/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order

	// Bind JSON body para o modelo Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Definir o status como "carrinho"
	order.Status = "carrinho"

	// Criar o pedido (carrinho) no banco de dados
	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func UpdateOrderStatus(c *gin.Context) {
	var order models.Order
	orderID := c.Param("id")
	var input struct {
		Status string `json:"status"` // O novo status que será atualizado
	}

	// Bind JSON body para capturar o novo status
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Buscar o pedido
	if err := database.DB.First(&order, "id = ?", orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Atualizar o status do pedido
	order.Status = input.Status

	// Adicionar o novo status ao histórico
	newHistoryEntry := models.OrderHistory{
		Status:    input.Status,
		Timestamp: time.Now(),
	}

	// Adicionar o novo evento ao campo HistoryLog (append na lista)
	order.HistoryLog = append(order.HistoryLog, newHistoryEntry)

	// Salvar o pedido atualizado com o novo status e o histórico
	if err := database.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}
