package ristretto_test

import (
	"testing"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/stretchr/testify/require"
	"gitoa.ru/go-4devs/cache"
	provider "gitoa.ru/go-4devs/cache/provider/ristretto"
	"gitoa.ru/go-4devs/cache/test"
)

func TestRistretto(t *testing.T) {
	t.Parallel()

	retto, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	require.Nil(t, err)
	test.RunSute(t, provider.New(retto), test.WithWaitGet(func() {
		time.Sleep(10 * time.Millisecond)
	}), test.WithExpire(cache.ErrCacheMiss))
}
