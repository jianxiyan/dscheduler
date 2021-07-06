package main

import (
	"github.com/jianxiyan/dscheduler/core"
	"github.com/jianxiyan/dscheduler/global"
	"github.com/jianxiyan/dscheduler/initialize"
)

func main() {
	global.G_VP = core.Viper() //初始化viper
	global.G_LOG = core.Zap()  //初始化日志库

	global.G_DB = initialize.Gorm()     // gorm链接数据库
	initialize.MysqlTables(global.G_DB) //初始化表
	//关闭程序前关闭数据库链接
	db, _ := global.G_DB.DB()
	defer db.Close()

	core.RunServer()
}
