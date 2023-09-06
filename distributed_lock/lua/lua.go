package lua

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Lua struct {
	Ctx      context.Context
	RedisCli *redis.Client
}

const (
	lockKey    = "mylock"
	lockExpire = 10 * time.Second
)

func (l *Lua) AcquireLock() bool {
	// 尝试获取锁
	success, err := l.RedisCli.SetNX(l.Ctx, lockKey, "locked", lockExpire).Result()
	if err != nil || !success {
		return false
	}
	return true
}

func (l *Lua) ReleaseLock() {
	// 释放锁
	l.RedisCli.Del(l.Ctx, lockKey)
}
