package admin

import (
	"firstProject/app/dto"
	myjwt "firstProject/app/http/middleware/jwt"
	"firstProject/app/http/result"
	"firstProject/app/models"
	adminRep "firstProject/app/repositories/admin"
	"firstProject/app/requests"
	"firstProject/app/services/admin"
	"fmt"

	"net/http"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func AdminLogin(c *gin.Context) {
	loginReq := requests.AdminLoginRequest{}
	returnData := result.NewResult(c)
	err := c.ShouldBind(&loginReq)

	if err != nil {
		returnData.Error("解析失败")
		return
	}

	model := adminRep.GetAdminByUsername(loginReq.Username)
	fmt.Println(model)
	if model.ID == 0 {
		returnData.Error("帐号错误")
		return
	}

	erra := bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(loginReq.Password))

	if erra != nil {
		returnData.Error("密码错误")
		return
	}

	generateToken(c, model)
}

// 生成令牌
func generateToken(c *gin.Context, user models.Admin) {
	j := &myjwt.JWT{
		[]byte("newtrekWang"),
	}
	claims := myjwt.CustomClaims{
		user.ID,
		user.Username,
		user.Password,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),     // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 36000000), // 过期时间 一小时
			Issuer:    "newtrekWang",                       //签名的发行者
		},
	}
	returnData := result.NewResult(c)
	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	returnData.Success(token)

}

// GetDataByTime 一个需要token认证的测试接口
/* func GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
} */
