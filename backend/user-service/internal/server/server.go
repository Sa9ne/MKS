package server

import (
	"log"
	"user-service/internal/handlers"
	"user-service/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	// Используем фреймворк GIN
	s := gin.Default()

	// Подключаем корс для корректной работы с фронтендом
	s.Use(cors.Default())

	// Маршрутизация путей
	s.POST("/MakeMessage", handlers.MakeMessage)
	s.POST("/Login", handlers.Login)

	// Настройка администратора
	AdminRoutes := s.Group("/admin")
	AdminRoutes.Use(middleware.AuthMiddleware())

	// Маршрутизация админов
	AdminRoutes.GET("/CheckMessage", handlers.LoadMessage)
	AdminRoutes.POST("/ReadMessage/:id", handlers.ReadMessage)

	// Выбираем порт работы сервера
	err := s.Run(":8081")
	if err != nil {
		log.Fatal("Failed to create server")
	}
}
