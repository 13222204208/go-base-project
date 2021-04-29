package admin

import (
	"firstProject/app/http/result"
	"firstProject/app/models"
	"firstProject/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GoodsList(c *gin.Context) {
	returnData := result.NewResult(c)
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	goods := make([]models.Goods, 0)
	//database.DB.AutoMigrate(&goods)
	database.DB.Limit(limit).Offset((page - 1) * limit).Find(&goods)
	result := database.DB.Find(&goods)
	var data struct {
		Item  []models.Goods `json:"item"`
		Total int            `json:"total"`
	}

	data.Item = goods
	data.Total = int(result.RowsAffected)
	returnData.Success(data)
}
