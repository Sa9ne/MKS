package server

import (
	"log"
	"main-service/internal/database"

	"github.com/gin-gonic/gin"
)

func Start() {
	// Используем фреймворк GIN
	s := gin.Default()

	// Подключаемся в бд
	database.ConnectDB()

	// Выбираем порт работы сервера
	err := s.Run(":8080")
	if err != nil {
		log.Fatal("Failed to create server")
	}
}
