package lua

import (
	"context"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

type Lua struct {
	Ctx      context.Context
	RedisCli *redis.Client
}

const (
	lockExpire = 10 * time.Second
)

var nodeIdentifier = rand.Int63n(10000)

func (l *Lua) AcquireLock(lockKey string) bool {
	// 尝试获取锁
	success, err := l.RedisCli.SetNX(l.Ctx, lockKey, nodeIdentifier, lockExpire).Result()
	if err != nil || !success {
		return false
	}

	// 设置一个过期时间，防止锁被无限持有
	l.RedisCli.Expire(l.Ctx, lockKey, lockExpire)

	return true
}

func (l *Lua) ReleaseLock(lockKey string) bool {
	// 释放锁
	luaScript := `
        if redis.call("get", KEYS[1]) == ARGV[1] then
            return redis.call("del", KEYS[1])
        else
            return 0
        end
    `
	result, err := l.RedisCli.Eval(l.Ctx, luaScript, []string{lockKey}, nodeIdentifier).Result()
	if err != nil {
		return false
	}
	return result == int64(1)
}
