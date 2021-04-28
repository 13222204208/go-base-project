package admin

import (
	"firstProject/app/dto"
	"firstProject/app/http/result"
	"firstProject/app/models"
	goodsTypeRep "firstProject/app/repositories/goodsType"
	"firstProject/app/requests"
	"firstProject/database"
	"fmt"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func CreateGoodsType(c *gin.Context) {
	request := requests.GoodsTypeRequest{}
	returnData := result.NewResult(c)

	erra := c.ShouldBind(&request)

	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
	//验证器注册翻译器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}
	err = validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			result.NewResult(c).Error(err.Translate(trans))
			return
		}
	}

	if erra != nil {
		fmt.Println(err)
	}

	goodsTypeDto := dto.GoodsTypeDto{
		Name: request.Name,
	}

	model := goodsTypeRep.CreateGoodsType(goodsTypeDto)
	if model != nil {
		returnData.Error("创建失败：" + err.Error())
		return
	}

	returnData.Success("创建成功")
}

func GoodsTypeList(c *gin.Context) {
	returnData := result.NewResult(c)
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	goodsType := make([]models.GoodsType, 0)
	database.DB.Limit(limit).Offset((page - 1) * limit).Find(&goodsType)
	result := database.DB.Find(&goodsType)
	var data struct {
		Item  []models.GoodsType `json:"item"`
		Total int                `json:"total"`
	}

	data.Item = goodsType
	data.Total = int(result.RowsAffected)
	returnData.Success(data)
}

func OneGoodsType(c *gin.Context) {
	returnData := result.NewResult(c)
	id := c.Param("id")
	model := database.DB.First(&models.GoodsType{}, id)
	if model.Error == nil {
		returnData.Success(model.Value)
		return
	}

	returnData.Error("获取失败")
}

func UpdateGoodsType(c *gin.Context) {
	returnData := result.NewResult(c)
	id := c.Param("id")
	var goodsType models.GoodsType
	database.DB.First(&goodsType, id)
	goodsType.Name = c.Query("name")
	err := database.DB.Save(&goodsType)
	if err.Error == nil {
		returnData.Success(goodsType)
		return
	}

	returnData.Error("修改失败")
}

func DeleteGoodsType(c *gin.Context) {
	returnData := result.NewResult(c)
	id := c.Param("id")
	model := database.DB.Delete(&models.GoodsType{}, id)

	if model.Error == nil {
		returnData.Success("删除成功")
		return
	}

	returnData.Error("删除失败")
}
