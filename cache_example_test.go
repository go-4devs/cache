package cache_test

import (
	"context"
	"fmt"
	"time"

	glru "github.com/hashicorp/golang-lru"
	prom "github.com/prometheus/client_golang/prometheus"
	"gitoa.ru/go-4devs/cache"
	"gitoa.ru/go-4devs/cache/mw"
	"gitoa.ru/go-4devs/cache/mw/prometheus"
	"gitoa.ru/go-4devs/cache/provider/lru"
	"gitoa.ru/go-4devs/cache/provider/memcache"
	"gitoa.ru/go-4devs/cache/provider/memory"
	"gitoa.ru/go-4devs/cache/provider/redis"
	"gitoa.ru/go-4devs/cache/test"
	"gitoa.ru/go-4devs/encoding/gob"
)

func ExampleCache_map() {
	ctx := context.Background()
	c := cache.New(memory.NewMap())

	var cached string

	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, "not found key", &cached), cached)
	fmt.Printf("err: %v\n", c.Set(ctx, "key", "some value"))
	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, "key", &cached), cached)
	// Output:
	// err: cache miss: map, value: ''
	// err: <nil>
	// err: <nil>, value: 'some value'
}

func ExampleCache_encoding() {
	ctx := context.Background()
	c := cache.New(memory.NewEncoding(), cache.WithDataOption(cache.WithMarshal(gob.Unmarshal, gob.Marshal)))

	var cached string

	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, "not found key", &cached), cached)
	fmt.Printf("err: %v\n", c.Set(ctx, "key", "some value"))
	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, "key", &cached), cached)
	// Output:
	// err: cache miss: encoding, value: ''
	// err: <nil>
	// err: <nil>, value: 'some value'
}

func ExampleCache_redis() {
	ctx := context.Background()
	c := cache.New(redis.New(test.RedisClient()))

	var cached string

	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, "not found redis key", &cached), cached)
	fmt.Printf("err: %v\n", c.Set(ctx, "key", "some redis value", cache.WithNamespace("redis", ":")))
	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, "redis:key", &cached), cached)
	// Output:
	// err: cache miss: redis pool, value: ''
	// err: <nil>
	// err: <nil>, value: 'some redis value'
}

func ExampleCache_memacache() {
	ctx := context.Background()
	c := cache.New(memcache.New(test.MemcacheClient()), cache.WithDataOption(cache.WithNamespace("memcache", ":")))

	var cached string

	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, "not found memcached key", &cached), cached)
	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, "not:found:memcached:key", &cached), cached)
	fmt.Printf("err: %v\n", c.Set(ctx, "key", "some mamcache value"))
	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, "key", &cached), cached)
	// Output:
	// err: key is not valid: memcache, value: ''
	// err: cache miss: memcache, value: ''
	// err: <nil>
	// err: <nil>, value: 'some mamcache value'
}

func ExampleCache_lru() {
	ctx := context.Background()
	client, _ := glru.New(10)

	c := cache.New(lru.New(client), cache.WithDataOption(cache.WithTTL(time.Hour)))

	var cached string

	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, "not found lru key", &cached), cached)
	fmt.Printf("err: %v\n", c.Set(ctx, "key", "some lru value"))
	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, "key", &cached), cached)
	fmt.Printf("deleted err: %v\n", c.Delete(ctx, "key"))
	// Output:
	// err: cache miss: lru, value: ''
	// err: <nil>
	// err: <nil>, value: 'some lru value'
	// deleted err: <nil>
}

func ExampleCache_withNamespace() {
	ctx := context.Background()
	c := cache.New(provider(), cache.WithDataOption(
		cache.WithNamespace("prefix", ":"),
		cache.WithTTL(time.Hour),
	))

	var cached, cached2 string

	fmt.Printf("prefix  err: %v, value: '%v'\n", c.Get(ctx, "key", &cached), cached)
	fmt.Printf("prefix  err: %v\n", c.Set(ctx, "key", "some value", cache.WithTTL(time.Minute)))
	fmt.Printf("prefix2 err: %v\n", c.Set(ctx, "key", "some value2", cache.WithNamespace("prefix2", ":")))
	fmt.Printf("prefix  err: %v, value: '%v'\n", c.Get(ctx, "key", &cached), cached)
	fmt.Printf("prefix2 err: %v, value: '%v'\n", c.Get(ctx, "key", &cached2, cache.WithNamespace("prefix2", ":")), cached2)
	// Output:
	// prefix  err: cache miss: map, value: ''
	// prefix  err: <nil>
	// prefix2 err: <nil>
	// prefix  err: <nil>, value: 'some value'
	// prefix2 err: <nil>, value: 'some value2'
}

func ExampleCache_withFallback() {
	ctx := context.Background()
	c := cache.New(provider(), mw.WithFallback(
		func(ctx context.Context, key, value interface{}) error {
			fmt.Printf("loaded key: %#v\n", key)

			return cache.TypeAssert("some loaded data", value)
		},
		func(i *cache.Item, e error) bool {
			return e != nil
		},
	))

	var cached, cached2 string

	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, 1, &cached), cached)
	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, 1, &cached2), cached2)
	// Output:
	// loaded key: 1
	// err: <nil>, value: 'some loaded data'
	// err: <nil>, value: 'some loaded data'
}

func ExampleCache_clearByContext() {
	type ctxKey int

	var (
		requestID       ctxKey = 1
		cached, cached2 string
	)

	ctx, cancel := context.WithCancel(context.WithValue(context.Background(), requestID, "unique ctx key"))
	ctx2 := context.WithValue(context.Background(), requestID, "unique ctx key2")
	c := cache.New(provider(),
		mw.WithClearByContext(requestID),
		cache.WithDataOption(cache.WithNamespace("clear_by_ctx", "")),
	)

	fmt.Printf("err: %v\n", c.Set(ctx, 1, "some ctx loaded data", cache.WithTTL(time.Hour)))
	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, 1, &cached), cached)
	cancel()
	time.Sleep(time.Millisecond)
	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx2, 1, &cached2), cached2)
	// Output:
	// err: <nil>
	// err: <nil>, value: 'some ctx loaded data'
	// err: cache miss: map, value: ''
}

func ExampleCache_clearByTTL() {
	ctx := context.Background()
	c := cache.New(provider(),
		mw.WithClearByTTL(),
		cache.WithDataOption(cache.WithNamespace("clear_by_ttl", "")),
	)

	var cached, cached2 string

	fmt.Printf("err: %v\n", c.Set(ctx, 1, "some ttl loaded data", cache.WithTTL(time.Microsecond*200)))
	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, 1, &cached), cached)
	time.Sleep(time.Second)
	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, 1, &cached2), cached2)
	// Output:
	// err: <nil>
	// err: <nil>, value: 'some ttl loaded data'
	// err: cache miss: map, value: ''
}

func ExampleCache_withMetrics() {
	ctx := context.Background()
	cacheLabel := "cache_label"
	c := cache.New(provider(),
		mw.WithMetrics(prometheus.Metrics{}, mw.LabelName(cacheLabel)),
		cache.WithDataOption(cache.WithNamespace("metrics", ":")),
	)

	var cached, cached2 string

	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, 1, &cached), cached)
	fmt.Printf("err: %v\n", c.Set(ctx, 1, "cached"))
	fmt.Printf("err: %v, value: '%v'\n", c.Get(ctx, 1, &cached2), cached2)

	mfs, _ := prom.DefaultGatherer.Gather()
	for _, mf := range mfs {
		for _, metric := range mf.GetMetric() {
			label := metric.GetLabel()
			if len(label) > 0 && metric.Counter != nil {
				fmt.Printf("name:%s, label:%s, value: %.0f\n", *mf.Name, *label[0].Value, mf.GetMetric()[0].Counter.GetValue())
			}
		}
	}

	// Output:
	// err: cache miss: map, value: ''
	// err: <nil>
	// err: <nil>, value: 'cached'
	// name:cache_hit_total, label:cache_label, value: 1
	// name:cache_miss_total, label:cache_label, value: 1
}

func provider() cache.Provider {
	return memory.NewMap()
}
