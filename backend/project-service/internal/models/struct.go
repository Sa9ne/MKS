package models

// Структура проекта
type Project struct {
	ID               int    `json:"id" gorm:"primaryKey"`
	Title            string `json:"Title"`
	InfoAboutProject string `json:"InfoAboutProject"`
}
