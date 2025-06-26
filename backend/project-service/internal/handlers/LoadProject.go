package handlers

import (
	"net/http"
	"project-service/internal/database"
	"project-service/internal/models"

	"github.com/gin-gonic/gin"
)

func LoadProject(ctx *gin.Context) {
	// Инициализируем переменную project
	var Project []models.Project

	// Ищем в базе данных проекты, если поиск выдал ошибку выводим ошибку поиска
	if err := database.DB.Find(&Project).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Project was not found"})
		return
	}

	// Выводим найденные проекты
	ctx.JSON(http.StatusOK, Project)
}
