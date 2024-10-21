package database

import (
	"fmt"
	"log"

	"go-api/config"
	"go-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.AppConfig.DBHost,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
		config.AppConfig.DBPort,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados:", err)
	}

	// Migrar os modelos
	DB.AutoMigrate(
		&models.User{},
		&models.Token{},
		&models.Establishment{},
		&models.Client{},
		&models.Address{},
		&models.Product{},
		&models.ProductType{},
		&models.ProductPhoto{},
		&models.OrderItem{},
		&models.Order{},
	)
}
