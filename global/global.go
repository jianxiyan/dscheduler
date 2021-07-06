package global

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jianxiyan/dscheduler/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	G_CONFING   config.Server
	G_VP        *viper.Viper
	G_DB        *gorm.DB
	G_LOG       *zap.Logger
	G_REDISPOOL *redis.Pool
)
