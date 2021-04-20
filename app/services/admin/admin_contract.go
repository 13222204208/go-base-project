package admin

import "firstProject/app/dto"

type AdminContract interface {
	Register(dto dto.AdminDto) error
	Login(dto dto.AdminDto) error
}
