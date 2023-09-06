package inc

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Inc struct {
	Ctx      context.Context
	RedisCli *redis.Client
}

const (
	lockKey    = "mylock"
	lockExpire = 10 * time.Second
)

func (i *Inc) AcquireLock() bool {
	// 尝试获取锁
	success, err := i.RedisCli.SetNX(i.Ctx, lockKey, "locked", lockExpire).Result()
	if err != nil || !success {
		return false
	}
	return true
}

func (i *Inc) ReleaseLock() bool {
	// 释放锁
	i.RedisCli.Del(i.Ctx, lockKey)
	return true
}
