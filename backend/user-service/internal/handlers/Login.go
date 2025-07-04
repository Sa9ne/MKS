package handlers

import (
	"log"
	"net/http"
	"os"
	"time"
	"user-service/internal/databases"
	"user-service/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Читаем JWT токен
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func Login(ctx *gin.Context) {
	var Input models.Admin

	// Парсим данные
	if err := ctx.ShouldBindJSON(&Input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parsing error"})
		return
	}

	// Поиск администратора по username
	var Admin models.Admin
	if err := databases.DB.Where("username = ?", Input.Username).First(&Admin).Error; err != nil {
		ctx.JSON(http.StatusNotFound, "Admin with this username was not found")
		return
	}

	// Проверяем пароль, который ввел пользователь
	if err := bcrypt.CompareHashAndPassword([]byte(Admin.Password), []byte(Input.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	// Время работы JWT токена
	JWTLifeTime := time.Now().Add(30 * 24 * time.Hour)

	// Создаём token
	claims := &models.Claims{
		Username: Admin.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(JWTLifeTime),
		},
	}

	// Подписываем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Token creation error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	// Отдаем токен пользователю
	ctx.JSON(http.StatusOK, tokenString)
}
