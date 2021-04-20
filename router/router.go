package router

import (
	admin "firstProject/app/http/controller/admin"
	minapp "firstProject/app/http/controller/minapp"
	"firstProject/app/http/middleware/cors"
	"firstProject/app/http/middleware/handler"
	"firstProject/app/http/middleware/logger"
	"firstProject/app/http/result"
	"firstProject/pkg/e"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	router.Use(handler.Recover)
	router.Use(logger.LogerMiddleware())
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(cors.Cors())
	api := router.Group("api")
	{
		api.GET("/test", minapp.Ping)
	}

	web := router.Group("web")
	{
		web.GET("user", admin.User)
		web.POST("register", e.ErrorWrapper(admin.RegisterHandle))
		web.POST("login", admin.UserLogin)
	}

	router.Run()
}

//404
func HandleNotFound(c *gin.Context) {
	result.NewResult(c).Error("路径或方法错误")
}
