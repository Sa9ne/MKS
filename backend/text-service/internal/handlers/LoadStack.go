package handlers

import (
	"net/http"
	"text-service/internal/databases"
	"text-service/internal/models"

	"github.com/gin-gonic/gin"
)

func LoadStack(ctx *gin.Context) {
	// Инициализируем стек технологий
	var Stack []models.StackTech

	// Поиск в базе данных и обработка ошибок
	if err := databases.DB.Find(&Stack).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Stack was not found"})
		return
	}

	// Вывод стека технологий
	ctx.JSON(http.StatusOK, Stack)
}
