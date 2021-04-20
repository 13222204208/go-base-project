package admin

import (
	"firstProject/app/dto"
	"firstProject/app/http/result"
	"firstProject/app/requests"
	"firstProject/app/services/admin"
	"fmt"

	"github.com/gin-gonic/gin"
)

/* type Admin struct {
	Username string `form:"username" json:"username" validate:"required" label:"用户名"`
	Password string `form:"password" json:"password" validate:"required" label:"密码"`
}

func Register(c *gin.Context) {
	var adminAccount Admin
	returnData := result.NewResult(c)

	erra := c.ShouldBind(&adminAccount)
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
	err = validate.Struct(adminAccount)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			result.NewResult(c).Error(err.Translate(trans))
			return
		}
	}

	if erra == nil {
		err := model.Register(adminAccount.Username, adminAccount.Password)
		if err == nil {
			returnData.Success("注册成功")
		} else {
			returnData.Error("注册失败" + err.Error())
		}
	} else {
		returnData.Error("解析数据失败")
	}

} */

//RegisterHandle 注册
func RegisterHandle(c *gin.Context) (interface{}, error) {
	request := requests.AdminRegisterRequest{}
	returnData := result.NewResult(c)
	err := c.ShouldBind(&request)
	if err != nil {
		fmt.Println(err)
	}

	userDto := dto.AdminDto{
		Username: request.Username,
		Password: request.Password,
	}

	service := admin.AdminService{}
	err = service.Register(userDto)
	if err != nil {

		returnData.Error("注册失败" + err.Error())

	}

	return "注册成功", nil
}
