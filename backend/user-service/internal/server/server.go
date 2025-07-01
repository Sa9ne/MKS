package server

import (
	"log"
	"user-service/internal/databases"
	"user-service/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	// Используем фреймворк GIN
	s := gin.Default()

	// Подключаем корс для корректной работы с фронтендом
	s.Use(cors.Default())

	// Подключаемся в бд
	databases.ConnectDB()

	// Маршрутизация путей
	s.POST("/MakeMessage", handlers.MakeMessage)

	// Выбираем порт работы сервера
	err := s.Run(":8081")
	if err != nil {
		log.Fatal("Failed to create server")
	}
}
