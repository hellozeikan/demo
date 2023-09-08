package main

import (
	"context"
	"fmt"
	"lock/interf"
	"lock/lua"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx context.Context
var redisClient *redis.Client

func Init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "1.15.248.7:6379", // Redis服务器地址
		Password: "自己的",
		DB:       0,
		PoolSize: 2,
	})
	ctx = context.Background()
}

func main() {
	// 初始化Redis客户端
	Init()
	defer redisClient.Close()
	// interF := interf.Register(&inc.Inc{Ctx: ctx, RedisCli: redisClient})
	interF := interf.Register(&lua.Lua{Ctx: ctx, RedisCli: redisClient})
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		ip := c.ClientIP()
		fmt.Println(ip)
		if interF.AcquireLock(ip) {
			defer interF.ReleaseLock(ip)
			c.JSON(http.StatusOK, gin.H{
				"message": "成功获取锁，执行操作...",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "未成功获取锁，请重试...",
			})
		}

	})
	r.Run()
}
