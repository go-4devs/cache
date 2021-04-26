package test

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"time"

	gom "github.com/bradfitz/gomemcache/memcache"
	"github.com/cockroachdb/pebble"
	"github.com/dgraph-io/ristretto"
	redigo "github.com/gomodule/redigo/redis"
	"gitoa.ru/go-4devs/cache/provider/redis"
)

// RedisClient created redis client.
func RedisClient() func(ctx context.Context) (redis.Conn, error) {
	host, ok := os.LookupEnv("FDEVS_CACHE_REDIS_HOST")
	if !ok {
		host = ":6379"
	}

	client := &redigo.Pool{
		DialContext: func(ctx context.Context) (redigo.Conn, error) {
			return redigo.DialContext(ctx, "tcp", host)
		},
	}

	return redis.NewPool(client)
}

// MemcacheClient created memcached client.
func MemcacheClient() *gom.Client {
	host, ok := os.LookupEnv("FDEVS_CACHE_MEMCACHE_HOST")
	if !ok {
		host = "localhost:11211"
	}

	return gom.New(host)
}

// RistrettoClient creates ristretto client.
func RistrettoClient() *ristretto.Cache {
	cache, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})

	return cache
}

// PebbleDB creates pebble DB.
func PebbleDB() (*pebble.DB, func()) {
	path := "demo.test"
	if _, filename, _, ok := runtime.Caller(0); ok {
		path = filepath.Dir(filename) + "/" + path
	}

	db, _ := pebble.Open(path, &pebble.Options{})

	return db, func() {
		os.RemoveAll(path)
	}
}

// User tested user data.
type User struct {
	ID        int
	Name      string
	UpdateAt  time.Time
	CreatedAt time.Time
}

// NewUser create mocks data user.
func NewUser(id int) User {
	return User{
		ID:        id,
		Name:      "andrey",
		UpdateAt:  time.Date(2020, 2, 1, 1, 2, 3, 4, time.UTC),
		CreatedAt: time.Date(1999, 2, 1, 1, 2, 3, 4, time.UTC),
	}
}
