package models

// Форма обратной связи
type Feedback struct {
	ID          int    `json:"ID" gorm:"primaryKey"`
	PhoneNumber string `json:"PhoneNumber"`
	TitleTheme  string `json:"TitleTheme"`
	Message     string `json:"Message"`
}
