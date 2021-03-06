package router

import (
	admin "firstProject/app/http/controller/admin"
	minapp "firstProject/app/http/controller/minapp"
	"firstProject/app/http/middleware/cors"
	"firstProject/app/http/middleware/handler"
	"firstProject/app/http/middleware/jwt"

	//"firstProject/app/http/middleware/logger"
	"firstProject/app/http/result"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	router.Use(handler.Recover)
	//router.Use(logger.LogerMiddleware())
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(cors.Cors())
	api := router.Group("api")
	{
		api.GET("/test", minapp.Ping)
	}

	web := router.Group("admin")
	{
		web.POST("login", admin.AdminLogin)
		web.POST("admin", admin.RegisterHandle) //注册后台用户
		web.POST("upload/img", admin.UploadImg) //图片上传
	}
	web.Use(jwt.JWTAuth())
	{
		web.GET("admin", admin.List)
		web.DELETE("admin/:id", admin.DeleteAdmin)
		web.PATCH("admin/:id", admin.UpdateAdminPassword)
		web.GET("info", admin.Info)
		web.POST("logout", admin.Logout)
		web.POST("goods_type", admin.CreateGoodsType)
		web.GET("goods_type", admin.GoodsTypeList)
		web.GET("goods_type/:id/edit", admin.OneGoodsType)
		web.PATCH("goods_type/:id", admin.UpdateGoodsType)
		web.DELETE("goods_type/:id", admin.DeleteGoodsType)

		web.GET("goods", admin.GoodsList)
	}

	router.Run()
}

//404
func HandleNotFound(c *gin.Context) {
	result.NewResult(c).Error("路径或方法错误")
}
