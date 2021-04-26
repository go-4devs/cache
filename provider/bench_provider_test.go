package provider_test

import (
	"context"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
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
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type provider struct {
	name string
	prov cache.Provider
}

func providers() []provider {
	client, _ := glru.New(10000)
	db, cl := test.PebbleDB()

	defer cl()

	return []provider{
		{"encoding", memory.NewEncoding()},
		{"map", memory.NewMap()},
		{"shard", memory.NewMapShard()},
		{"lru", lru.New(client)},
		{"ristretto", ristretto.New(test.RistrettoClient())},
		{"memcache", memcache.New(test.MemcacheClient())},
		{"redis", redis.New(test.RedisClient())},
		{"pebble", pebble.New(db)},
	}
}

func randStringBytes(n int64) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = letterBytes[randInt64(int64(len(letterBytes)))]
	}

	return string(b)
}

func randInt64(max int64) int64 {
	m := big.NewInt(max)
	nBig, _ := rand.Int(rand.Reader, m)

	return nBig.Int64()
}

func BenchmarkCacheGetRandomKeyString(b *testing.B) {
	ctx := context.Background()
	keysLen := 10000

	for _, p := range providers() {
		prov := p.prov
		items := make([]*cache.Item, keysLen)

		for i := 0; i < keysLen; i++ {
			var val string

			key := randStringBytes(55)
			items[i] = cache.NewItem(key, &val)
			require.Nil(b, prov(ctx, cache.OperationSet, cache.NewItem(key, "value: "+p.name, cache.WithTTL(time.Minute))))
		}

		b.Run(p.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = prov(ctx, cache.OperationGet, items[i%keysLen])
			}
		})
	}
}

func BenchmarkCacheGetRandomKeyInt(b *testing.B) {
	ctx := context.Background()
	keysLen := 10000

	for _, p := range providers() {
		prov := p.prov
		items := make([]*cache.Item, keysLen)

		for i := 0; i < keysLen; i++ {
			var val int64

			key := randInt64(math.MaxInt64)

			items[i] = cache.NewItem(key, &val)
			require.Nil(b, prov(ctx, cache.OperationSet, cache.NewItem(key, randInt64(math.MaxInt64), cache.WithTTL(time.Minute))))
		}

		b.Run(p.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = prov(ctx, cache.OperationGet, items[i%keysLen])
			}
		})
	}
}

func BenchmarkCacheGetStruct(b *testing.B) {
	ctx := context.Background()

	type testStruct struct {
		Key string
		val string
	}

	var val testStruct
	item := cache.NewItem("key", &val)

	for _, p := range providers() {
		prov := p.prov
		require.Nil(b, prov(ctx, cache.OperationSet, cache.NewItem("key", testStruct{Key: "key", val: ""}, cache.WithTTL(time.Minute))))

		b.Run(p.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = prov(ctx, cache.OperationGet, item)
			}
		})
	}
}

func BenchmarkCacheSetStruct(b *testing.B) {
	ctx := context.Background()

	type testStruct struct {
		Key string
		Val int
	}

	for _, p := range providers() {
		prov := p.prov
		b.Run(p.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				item := cache.NewItem(i, testStruct{"k", i}, cache.WithTTL(time.Hour))
				_ = prov(ctx, cache.OperationSet, item)
			}
		})
	}
}

func BenchmarkCacheGetParallel(b *testing.B) {
	ctx := context.Background()

	for _, p := range providers() {
		prov := p.prov
		key := fmt.Sprintf("key_%s", p.name)
		val := fmt.Sprintf("value_%s", p.name)
		item := cache.NewItem(key, &val, cache.WithTTL(time.Minute))
		require.Nil(b, prov(ctx, cache.OperationSet, item))

		b.Run(p.name, func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = prov(ctx, cache.OperationGet, item)
				}
			})
		})
	}
}

func BenchmarkCacheSetParallel(b *testing.B) {
	ctx := context.Background()

	for _, p := range providers() {
		prov := p.prov
		key := fmt.Sprintf("key: %v", prov)
		val := fmt.Sprintf("value: %v", prov)

		b.Run(p.name, func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					item := cache.NewItem(key, val, cache.WithTTL(time.Hour))
					_ = prov(ctx, cache.OperationSet, item)
				}
			})
		})
	}
}
