package models

import "github.com/jinzhu/gorm"

type Goods struct {
	gorm.Model
	Title       string
	Summary     string
	Cover       string
	Content     string
	GoodsTypeId uint
	Sort        uint
	Price       string
	IsRecommend uint8
}
