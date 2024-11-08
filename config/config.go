package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	JWTSecret  string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Erro ao carregar o arquivo .env, usando variáveis de ambiente do sistema")
	}

	AppConfig = Config{
		DBUser:     getEnv("DB_USER", "docker"),
		DBPassword: getEnv("DB_PASSWORD", "docker"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "postgres"),
		JWTSecret:  getEnv("JWT_SECRET", "251099thiago"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
