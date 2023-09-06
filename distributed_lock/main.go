package main

import (
	"context"
	"fmt"
	"lock/inc"
	"lock/interf"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx context.Context
var redisClient *redis.Client

func Init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis服务器地址
	})
	ctx = context.Background()
}

func main() {
	// 初始化Redis客户端
	Init()
	defer redisClient.Close()
	interF := interf.Register(&inc.Inc{Ctx: ctx, RedisCli: redisClient})
	// 尝试获取锁
	if interF.AcquireLock() {
		defer interF.ReleaseLock() // 确保在函数结束时释放锁

		// 执行需要加锁的操作
		fmt.Println("成功获取锁，执行操作...")
		// 这里可以执行需要保护的代码块

		// 模拟长时间运行的任务
		time.Sleep(5 * time.Second)

		fmt.Println("操作完成，释放锁")
	} else {
		fmt.Println("未能获取锁，退出...")
	}
}
