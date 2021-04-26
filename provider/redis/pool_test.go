package redis_test

import (
	"testing"

	"gitoa.ru/go-4devs/cache"
	"gitoa.ru/go-4devs/cache/provider/redis"
	"gitoa.ru/go-4devs/cache/test"
)

func TestRedisPool(t *testing.T) {
	t.Parallel()
	test.RunSute(t, redis.New(test.RedisClient()), test.WithExpire(cache.ErrCacheMiss))
}
