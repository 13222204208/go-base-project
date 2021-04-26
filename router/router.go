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
		web.POST("register", admin.RegisterHandle)
	}
	web.Use(jwt.JWTAuth())
	{
		web.GET("user", admin.User)
		web.GET("info", admin.Info)
	}

	router.Run()
}

//404
func HandleNotFound(c *gin.Context) {
	result.NewResult(c).Error("路径或方法错误")
}
