package admin

import (
	"firstProject/app/dto"
	"firstProject/app/models"
	"firstProject/database"
)

//CreateUser 创建用户
func CreateAdmin(dto dto.AdminDto) error {
	admin := models.Admin{}
	admin.Username = dto.Username
	admin.Password = dto.Password

	err := database.DB.Create(&admin).Error

	return err
}
