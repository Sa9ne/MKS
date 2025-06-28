package handlers

import (
	"net/http"
	"text-service/internal/databases"
	"text-service/internal/models"

	"github.com/gin-gonic/gin"
)

func LoadArea(ctx *gin.Context) {
	// Инициализируем стек технологий
	var Area []models.StackTech

	// Поиск в базе данных и обработка ошибок
	if err := databases.DB.Find(&Area).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Area was not found"})
		return
	}

	// Вывод стека технологий
	ctx.JSON(http.StatusOK, Area)
}
