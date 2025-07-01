package handlers

import (
	"net/http"
	"regexp"
	"user-service/internal/databases"
	"user-service/internal/models"

	"github.com/gin-gonic/gin"
)

// Сама проверка номера телефона
func IsValidPhone(PhoneNumber string) bool {
	re := regexp.MustCompile(`^(?:\+7|8)\d{10}$`)
	return re.MatchString(PhoneNumber)
}

func MakeMessage(ctx *gin.Context) {
	var FeedBack models.Feedback

	// Парсим запрос пользователя
	if err := ctx.ShouldBindJSON(&FeedBack); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Проверяем номер телефона
	if !IsValidPhone(FeedBack.PhoneNumber) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong phone number"})
		return
	}

	// Проверка пустых полей
	if len(FeedBack.TitleTheme) == 0 || len(FeedBack.Message) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Message or Title was empty"})
		return
	}

	// Создание заказа
	if err := databases.DB.Create(&FeedBack).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Feedback was not created"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Successful feedback create"})
}
