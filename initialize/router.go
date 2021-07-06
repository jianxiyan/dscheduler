package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/jianxiyan/dscheduler/global"
	"github.com/jianxiyan/dscheduler/middleware"
	"github.com/jianxiyan/dscheduler/router"
)

//初始化总路由
func Router() *gin.Engine {
	Router := gin.Default()

	global.G_LOG.Info("use middleware logger")
	//跨域
	Router.Use(middleware.Cors())
	global.G_LOG.Info("use middleware cors")

	PublicGroup := Router.Group("")
	{
		router.InitUserRouter(PublicGroup) //注册功能
	}

	global.G_LOG.Info("router register success")

	return Router

}
