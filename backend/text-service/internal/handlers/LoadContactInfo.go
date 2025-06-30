package handlers

import (
	"net/http"
	"text-service/internal/databases"
	"text-service/internal/models"

	"github.com/gin-gonic/gin"
)

func LoadContactInfo(ctx *gin.Context) {
	var ContactInformation []models.ContactInformation

	if err := databases.DB.Find(&ContactInformation).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Contact information was not found"})
		return
	}

	ctx.JSON(http.StatusOK, ContactInformation)
}
