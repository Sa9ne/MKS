package databases

import (
	"fmt"
	"log"
	"os"
	"user-service/internal/models"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateAdmin() {
	// Пока что не абсолютный путь, в последствии при подключении docker, либо же на хост надо будет менять
	err := godotenv.Load("/Users/user/important/MKS/.env")
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	// Чтение из env файла
	username := os.Getenv("ADMIN_NAME")
	password := os.Getenv("ADMIN_PASSWORD")

	// Проверка на пустой ввод ввод
	if username == "" || password == "" {
		log.Println("ADMIN_NAME or ADMIN_PASSWORD are null in .env file")
		return
	}

	var Admin models.Admin

	// Ищем есть ли такой админ
	Find := DB.Where("username = ?", username).First(&Admin)
	// Если ошибок в поиске нет, значит запись была найдена, что означает что админ уже создан
	if Find.Error == nil {
		log.Println("Administrator with this username was created early ✅")
		return
	}
	// Проверка других ошибок
	if Find.Error != nil && Find.Error != gorm.ErrRecordNotFound {
		log.Fatalf("Error administrator: %v", Find.Error)
	}

	// Хэшируем пароль и ставим дефолтную "стоимость" шифрации
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Password was not hashed", err)
	}

	// Создание структуры админа с хэшированным паролем
	CreatedAdmin := models.Admin{
		Username: username,
		Password: string(HashedPassword),
	}

	// Внесение администрации в БД
	if err := DB.Create(&CreatedAdmin).Error; err != nil {
		log.Fatal("New admin was not created", err)
	}

	fmt.Println("Admin created successful ✅")
}
