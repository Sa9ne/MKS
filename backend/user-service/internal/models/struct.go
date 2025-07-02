package models

// Форма обратной связи
type Feedback struct {
	ID          int    `json:"ID" gorm:"primaryKey"`
	PhoneNumber string `json:"PhoneNumber"`
	TitleTheme  string `json:"TitleTheme"`
	Message     string `json:"Message"`
}

// Админ
type Admin struct {
	ID       uint   `json:"ID" gorm:"primaryKey"`
	Username string `json:"Username" gorm:"unique"`
	Password string `json:"Password"`
}
