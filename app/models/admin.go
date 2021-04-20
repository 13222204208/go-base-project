package models

import (
	"firstProject/database"
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	gorm.Model
	Username string
	Password string
}

func Register(username string, password string) error {

	//database.DB.AutoMigrate(&Admin{})
	database.DB.Where("name = ?", username).First(&Admin{})
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理

	if err != nil {
		fmt.Println(err)
	}
	encodePWD := string(hash)
	database.DB.Create(&Admin{Username: username, Password: encodePWD})

	return nil
}