package memcache_test

import (
	"testing"

	"gitoa.ru/go-4devs/cache"
	"gitoa.ru/go-4devs/cache/provider/memcache"
	"gitoa.ru/go-4devs/cache/test"
)

func TestProvider(t *testing.T) {
	t.Parallel()
	test.RunSute(t, memcache.New(test.MemcacheClient()), test.WithExpire(cache.ErrCacheMiss))
}
