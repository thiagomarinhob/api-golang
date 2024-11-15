package controllers

import (
	"fmt"
	"net/http"
	"time"

	"go-api/config"
	"go-api/database"
	"go-api/models"
	"go-api/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// Não hashear a senha aqui! Apenas comparar.
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		fmt.Println("Erro ao comparar senhas:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// Gerar Token JWT
	accessToken, err := utils.GenerateToken(user.ID, time.Hour*24)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Não foi possível gerar o token de acesso"})
		return
	}

	// Gerar Refresh Token
	refreshTokenString, err := utils.GenerateToken(user.ID, time.Hour*24*7) // Refresh Token expira em 7 dias
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Não foi possível gerar o refresh token"})
		return
	}

	// Armazenar o Refresh Token no banco de dados
	refreshToken := models.Token{
		Token:     refreshTokenString,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	if err := database.DB.Create(&refreshToken).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Não foi possível salvar o refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
		"access_token":  accessToken,
		"refresh_token": refreshTokenString,
	})
}

func RefreshToken(c *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar se o Refresh Token é válido
	token, err := jwt.Parse(input.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token inválido"})
		return
	}

	// Verificar se o Refresh Token existe no banco de dados
	var storedToken models.Token
	if err := database.DB.Where("token = ?", input.RefreshToken).First(&storedToken).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token inválido"})
		return
	}

	// Verificar se o Refresh Token expirou
	if storedToken.ExpiresAt < time.Now().Unix() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token expirado"})
		return
	}

	// Gerar novo Access Token
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	accessToken, err := utils.GenerateToken(userID, time.Minute*15)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Não foi possível gerar o token de acesso"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}

func Logout(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	// Opcional: Receber o refresh token no corpo da requisição
	var input struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Excluir o Refresh Token do banco de dados
	if err := database.DB.Where("user_id = ? AND token = ?", userID, input.RefreshToken).Delete(&models.Token{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Não foi possível revogar o token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout realizado com sucesso"})
}
