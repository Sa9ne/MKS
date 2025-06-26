package handlers

import (
	"net/http"
	"project-service/internal/database"
	"project-service/internal/models"

	"github.com/gin-gonic/gin"
)

func LoadProject(ctx *gin.Context) {
	var Project []models.Project

	if err := database.DB.Find(&Project).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Project was not found"})
		return
	}

	ctx.JSON(http.StatusOK, Project)
}
