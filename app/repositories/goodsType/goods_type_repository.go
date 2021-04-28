package goodsType

import (
	"firstProject/app/dto"
	"firstProject/app/models"
	"firstProject/database"
)

func CreateGoodsType(dto dto.GoodsTypeDto) error {
	goodsType := models.GoodsType{}
	//database.DB.AutoMigrate(&goodsType)
	goodsType.Name = dto.Name

	err := database.DB.Create(&goodsType).Error
	return err
}
