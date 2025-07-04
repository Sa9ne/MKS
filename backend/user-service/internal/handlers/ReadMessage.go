package handlers

import (
	"net/http"
	"user-service/internal/databases"
	"user-service/internal/models"

	"github.com/gin-gonic/gin"
)

func ReadMessage(ctx *gin.Context) {
	// Понимаем какое ID у выбранного сообщения
	ID := ctx.Param("id")

	// Ищем сообщение по ID
	var Message models.Feedback
	if err := databases.DB.First(&Message, ID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		return
	}

	// Отмечаем статус прочитанным
	Message.ReadStatus = true

	if err := databases.DB.Save(&Message).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update message"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Message successful read"})
}
