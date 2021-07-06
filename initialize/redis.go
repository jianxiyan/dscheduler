package initialize

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/jianxiyan/dscheduler/global"

	"go.uber.org/zap"
)

func Redis() {
	global.G_REDISPOOL = &redis.Pool{ //实例化一个连接池
		MaxIdle: global.G_CONFING.Redis.MaxIdle, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   global.G_CONFING.Redis.MaxActive,                  //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: time.Duration(global.G_CONFING.Redis.IdleTimeout), //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			c, err := redis.Dial("tcp", global.G_CONFING.Redis.Addr)
			if err != nil {
				global.G_LOG.Error("get redis connect failed, err:", zap.Any("err", err))
				return nil, err
			}
			if global.G_CONFING.Redis.Password != "" {
				if _, err := c.Do("AUTH", global.G_CONFING.Redis.Password); err != nil {
					c.Close()
					global.G_LOG.Error("redis connect Password failed, err:", zap.Any("err", err))
					return nil, err
				}
			}

			return c, err
		},
	}
}
