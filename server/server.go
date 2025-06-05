package server

import (
	"github.com/gin-gonic/gin"
	"staging/logger"
	"staging/pkg/settings"
	"staging/service"
	"strconv"
)

var svc *service.Service

func initRouter() *gin.Engine {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// todo 注册使用的中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	return r
}

func Init(app *settings.AppConfig) {
	svc = service.InitService(app)
	router := initRouter()
	router.GET("/test", test)

	router.Run(":" + strconv.Itoa(app.Port))
}
