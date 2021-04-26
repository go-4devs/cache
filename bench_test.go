package cache_test

import (
	"context"
	"testing"
	"time"

	glru "github.com/hashicorp/golang-lru"
	"github.com/stretchr/testify/require"
	"gitoa.ru/go-4devs/cache"
	"gitoa.ru/go-4devs/cache/provider/lru"
	"gitoa.ru/go-4devs/cache/provider/memcache"
	"gitoa.ru/go-4devs/cache/provider/memory"
	"gitoa.ru/go-4devs/cache/provider/pebble"
	"gitoa.ru/go-4devs/cache/provider/redis"
	"gitoa.ru/go-4devs/cache/provider/ristretto"
	"gitoa.ru/go-4devs/cache/test"
	"gitoa.ru/go-4devs/encoding/gob"
)

type cacheBench struct {
	name  string
	cache *cache.Cache
}

func cacheBenchList() []cacheBench {
	client, _ := glru.New(10000)
	db, cl := test.PebbleDB()

	defer cl()

	return []cacheBench{
		{"encoding json", cache.New(memory.NewEncoding())},
		{"encoding gob", cache.New(memory.NewEncoding(), cache.WithDataOption(cache.WithMarshal(gob.Unmarshal, gob.Marshal)))},
		{"map", cache.New(memory.NewMap())},
		{"map shards", cache.New(memory.NewMapShard())},
		{"ristretto", cache.New(ristretto.New(test.RistrettoClient()))},
		{"lru", cache.New(lru.New(client))},
		{"redis json", cache.New(redis.New(test.RedisClient()))},
		{"redis gob", cache.New(redis.New(test.RedisClient()), cache.WithDataOption(cache.WithMarshal(gob.Unmarshal, gob.Marshal)))},
		{"memcache json", cache.New(memcache.New(test.MemcacheClient()))},
		{"memcache gob", cache.New(memcache.New(test.MemcacheClient()), cache.WithDataOption(cache.WithMarshal(gob.Unmarshal, gob.Marshal)))},
		{"pebble json", cache.New(pebble.New(db))},
	}
}

type testStruct struct {
	Key string
	Val string
}

func BenchmarkCacheGetStruct(b *testing.B) {
	ctx := context.Background()

	var val testStruct

	for _, c := range cacheBenchList() {
		current := c.cache
		require.Nil(b, current.Set(ctx, "key", testStruct{"key", c.name}, cache.WithTTL(time.Minute)))

		b.Run(c.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = current.Get(ctx, "key", &val)
			}
		})
	}
}
