package server

import (
	"log"
	"text-service/internal/databases"
	"text-service/internal/handlers"

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
	s.GET("/LoadProject", handlers.LoadProject)
	s.GET("/LoadStack", handlers.LoadStack)
	s.GET("/LoadArea", handlers.LoadArea)
	s.GET("/LoadContactInfo", handlers.LoadContactInfo)

	// Выбираем порт работы сервера
	err := s.Run(":8080")
	if err != nil {
		log.Fatal("Failed to create server")
	}
}
