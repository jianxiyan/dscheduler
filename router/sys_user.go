package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jianxiyan/dscheduler/controller"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("register", controller.Register)
	}
}
