package handlers

import (
	"net/http"
	"user-service/internal/databases"
	"user-service/internal/models"

	"github.com/gin-gonic/gin"
)

func CheckMessage(ctx *gin.Context) {
	var Feedback []models.Feedback

	if err := databases.DB.Find(&Feedback).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Feedback was not found"})
		return
	}

	ctx.JSON(http.StatusOK, Feedback)
}
