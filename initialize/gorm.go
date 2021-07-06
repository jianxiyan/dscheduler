package initialize

import (
	"os"

	"github.com/jianxiyan/dscheduler/global"
	"github.com/jianxiyan/dscheduler/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.G_CONFING.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// mysqlTables 注册数据库表
func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
		model.Project{},
		model.ProjectPermission{},
		model.Processor{},
		model.ProcessorVersion{},
		model.Menu{},
		model.Role{},
		model.UserRole{},
		model.RoleParent{},
		model.SynData{},
	)
	if err != nil {
		global.G_LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.G_LOG.Info("register table success")
}

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	// dirverName := "mysql"
	m := global.G_CONFING.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		global.G_LOG.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		global.G_LOG.Info("MySQL启动成功")
		return db
	}
}
