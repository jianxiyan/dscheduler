package core

import (
	"fmt"
	"time"

	"github.com/jianxiyan/dscheduler/global"
	"github.com/jianxiyan/dscheduler/initialize"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	if global.G_CONFING.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
		defer global.G_REDISPOOL.Close()
	}

	//初始化路由
	Router := initialize.Router()
	address := fmt.Sprintf("%s:%d", global.G_CONFING.System.Addr, global.G_CONFING.System.Port)
	//初始化服务
	s := initServer(address, Router)

	time.Sleep(10 * time.Microsecond)
	global.G_LOG.Info("server run success on ", zap.String("address", address))
	global.G_LOG.Error(s.ListenAndServe().Error())
}
