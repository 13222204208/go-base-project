package admin

import (
	"firstProject/app/http/result"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"

	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func User(c *gin.Context) {
	//var slice = []int{1, 2, 3, 4, 5}
	//slice[6] = 6
	returnData := result.NewResult(c)
	returnData.Success(true)
}

type Users struct {
	Name   string `form:"name" json:"name" validate:"required" label:"姓名"`
	Age    uint8  `form:"age" json:"age" validate:"required,gt=18,lt=128" label:"年龄"`
	Passwd string `form:"passwd" json:"passwd" validate:"max=20,min=6" label:"密码"`
	Code   string `form:"code" json:"code" validate:"required,len=6" label:"验证码"`
}

func UserLogin(c *gin.Context) {
	var logins Users
	returnData := result.NewResult(c)

	erra := c.ShouldBind(&logins)
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
	err = validate.Struct(logins)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			result.NewResult(c).Error(err.Translate(trans))
			return
		}
	}

	if erra == nil {

		returnData.Success(logins)
		return
	} else {
		returnData.Error("error")
		return
	}

}
