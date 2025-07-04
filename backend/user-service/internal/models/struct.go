package models

import "github.com/golang-jwt/jwt/v5"

// Форма обратной связи
type Feedback struct {
	ID          int    `json:"ID" gorm:"primaryKey"`
	PhoneNumber string `json:"PhoneNumber"`
	TitleTheme  string `json:"TitleTheme"`
	Message     string `json:"Message"`
	ReadStatus  bool   `json:"Status" gorm:"default:false"`
}

// Админ
type Admin struct {
	ID       uint   `json:"ID" gorm:"primaryKey"`
	Username string `json:"Username" gorm:"unique"`
	Password string `json:"Password"`
}

type Claims struct {
	Username string `json:"Username"`
	jwt.RegisteredClaims
}
