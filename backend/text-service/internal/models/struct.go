package models

// Направление деятельности
type AreaOfActivity struct {
	ID              int    `json:"ID" gorm:"primaryKey"`
	AreaTitle       string `json:"AreaTitle"`
	AreaDescription string `json:"AreaDescription"`
}

// Структура проекта
type Project struct {
	ID               int    `json:"ID" gorm:"primaryKey"`
	Title            string `json:"Title"`
	InfoAboutProject string `json:"InfoAboutProject"`
	Img              string `json:"Img"`
}

// Стек технологий
type StackTech struct {
	ID        int    `json:"ID" gorm:"primaryKey"`
	StackName string `json:"StackName"`
}

// Контактная информация
type ContactInformation struct {
	ID          int    `json:"ID" gorm:"primaryKey"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	PhoneNumber string `json:"PhoneNumber"`
	Email       string `json:"Email"`
	Address     string `json:"Address"`
}
