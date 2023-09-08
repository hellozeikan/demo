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
	lockExpire = 10 * time.Second
)

func (i *Inc) AcquireLock(lockKey string) bool {
	// 尝试获取锁
	success, err := i.RedisCli.SetNX(i.Ctx, lockKey, "locked", lockExpire).Result()
	if err != nil || !success {
		return false
	}
	return true
}

func (i *Inc) ReleaseLock(lockKey string) bool {
	// 释放锁
	i.RedisCli.Del(i.Ctx, lockKey)
	return true
}
