package models

import "github.com/jinzhu/gorm"

type GoodsType struct {
	gorm.Model
	Name string
}
