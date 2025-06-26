package server

import (
	"log"
	"project-service/internal/database"
	"project-service/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	// Используем фреймворк GIN
	s := gin.Default()

	// Подключаем корс для корректной работы с фронтендом
	s.Use(cors.Default())

	// Подключаемся в бд
	database.ConnectDB()

	// Маршрутизация путей
	s.GET("/LoadProject", handlers.LoadProject)

	// Выбираем порт работы сервера
	err := s.Run(":8080")
	if err != nil {
		log.Fatal("Failed to create server")
	}
}
