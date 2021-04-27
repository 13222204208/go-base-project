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
	admin.Name = dto.Name
	admin.Phone = dto.Phone
	err := database.DB.Create(&admin).Error

	return err
}

//GetUserByUsername 通过用户名查询用户
func GetAdminByUsername(username string) models.Admin {
	user := models.Admin{}
	database.DB.Where("username = ?", username).First(&user)

	return user
}

/* func AdminList(username string, page , limit int) models.Admin{

} */
