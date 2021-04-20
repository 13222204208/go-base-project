package admin

import (
	"errors"
	"firstProject/app/dto"
	adminRep "firstProject/app/repositories/admin"

	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
}

func (admin AdminService) Register(adminDto dto.AdminDto) error {
	//密码加密
	bytes, err := bcrypt.GenerateFromPassword([]byte(adminDto.Password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return errors.New(err.Error())
	}
	adminDto.Password = string(bytes)
	return adminRep.CreateAdmin(adminDto)
}
