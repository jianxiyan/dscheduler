package controller

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jianxiyan/dscheduler/global"
	"github.com/jianxiyan/dscheduler/model"
)

func Register(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")

	if len(phone) != 11 {
		model.FailWithMessage("电话号码错误", c)
		return
	}

	if len(password) < 6 {
		model.FailWithMessage("密码不小于6位", c)
		return
	}

	if len(name) == 0 {
		name = randomString(10)
	}

	c1 := global.G_REDISPOOL.Get()
	defer c1.Close()
	_, err := c1.Do("set", "abc", 300)
	if err != nil {
		fmt.Println("get abc faild :", err)
		return
	}
	r, _ := redis.Int(c1.Do("get", "abc"))
	fmt.Println(r)
	model.OkWithData(map[string]string{
		"name":     name,
		"password": password,
		"phone":    phone,
	}, c)
}

func randomString(n int) string {
	temp := "abcdefghijklmnopqistuvwxyzABCDEFGHIJKLMNOPQISTUVWXYZ"
	sta := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for k, _ := range sta {
		sta[k] = temp[rand.Intn(len(temp))]
	}
	return string(sta)
}
