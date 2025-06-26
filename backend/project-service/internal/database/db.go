package database

import (
	"log"
	"os"
	"project-service/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Объявляем переменную DB для работы с ORM (В моем случае в GORM)
var DB *gorm.DB

func ConnectDB() {
	// Пока что не абсолютный путь, в последствии при подключении docker, либо же на хост надо будет менять
	err := godotenv.Load("/Users/user/important/MKS/.env")
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	// Читаем пути в .env файлах
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Database_url not found")
	}

	// Подключаемся пока что к postgreSQL, так как не знаю какую СУБД используем пока postgreSQL
	var errOpen error
	DB, errOpen = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errOpen != nil {
		log.Fatalf("Failed connect database: %v", errOpen)
	}

	// Проводим автомиграцию таблицы в бд
	DB.AutoMigrate(&models.Project{})
}
