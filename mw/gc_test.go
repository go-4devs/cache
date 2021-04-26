package mw_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gitoa.ru/go-4devs/cache"
	"gitoa.ru/go-4devs/cache/mw"
	"gitoa.ru/go-4devs/cache/provider/memory"
)

func TestWithClearByTTL(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	gcMap := cache.New(memory.NewMap(), mw.WithClearByTTL())
	cacheMap := cache.New(memory.NewMap())

	var (
		value string
		err   error
	)

	require.NoError(t, gcMap.Set(ctx, "keys", "value", cache.WithTTL(time.Second/3)))
	require.NoError(t, cacheMap.Set(ctx, "keys", "value", cache.WithTTL(time.Second/3)))
	time.Sleep(time.Second)

	err = gcMap.Get(ctx, "keys", &value)
	require.EqualError(t, err, cache.ErrCacheMiss.Error()+": map")

	err = cacheMap.Get(ctx, "keys", &value)
	require.EqualError(t, err, cache.ErrCacheExpired.Error()+": map")

	require.NoError(t, gcMap.Set(ctx, "keys", "value", cache.WithTTL(time.Second/2)))
	time.AfterFunc(time.Second/3, func() {
		require.NoError(t, gcMap.Set(ctx, "keys", "value", cache.WithTTL(time.Second)))
	})
	time.Sleep(time.Second / 2)
	require.NoError(t, gcMap.Get(ctx, "keys", &value))
	require.Equal(t, value, "value")
}

func TestWithClearByContext(t *testing.T) {
	t.Parallel()

	type ctxKey int

	var (
		requestID ctxKey = 1
		data      string
	)

	ctx1, cancel1 := context.WithCancel(context.WithValue(context.Background(), requestID, "request1"))
	ctx2, cancel2 := context.WithCancel(context.WithValue(context.Background(), requestID, "request2"))

	cacheMap := cache.New(memory.NewMap(), mw.WithClearByContext(requestID))

	require.NoError(t, cacheMap.Set(ctx1, "key", "value"))
	require.EqualError(t, cacheMap.Get(ctx2, "key", &data), "cache miss: map")
	require.NoError(t, cacheMap.Get(ctx1, "key", &data))
	cancel1()

	time.Sleep(time.Millisecond)
	require.EqualError(t, cacheMap.Get(ctx1, "key", &data), "cache miss: map")
	cancel2()
}
