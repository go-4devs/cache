package mw_test

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitoa.ru/go-4devs/cache"
	"gitoa.ru/go-4devs/cache/mw"
	"gitoa.ru/go-4devs/cache/test"
)

var (
	errFallback = errors.New("fallback error")
	errKey      = errors.New("unexpected key")
)

func TestWithFallback(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	key1 := "fallback:key1"
	key2 := "fb:key2"

	prov := test.NewProviderMock(t,
		test.WithGet(cacheGetMiss),
		test.WithSet(cacheSetMiss(map[interface{}]test.User{
			key1: test.NewUser(1),
			key2: test.NewUser(2),
		})),
	)
	c := cache.New(prov, mw.WithFallback(
		func(ctx context.Context, key, value interface{}) error {
			switch key.(string) {
			case key1:
				*value.(*test.User) = test.NewUser(1)
			case key2:
				*value.(*test.User) = test.NewUser(2)
			default:
				t.Errorf("unexpected key: %s", key)
			}

			return nil
		},
		mw.HandleByErr,
	))

	var user test.User

	require.Nil(t, c.Get(ctx, key1, &user))
	require.Equal(t, test.NewUser(1), user)

	require.Nil(t, c.Get(ctx, key2, &user, cache.WithNamespace("namespace", ":")))
	require.Equal(t, test.NewUser(2), user)
}

func TestLockFallback(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	var (
		val1, val2, val3 string
		wg               sync.WaitGroup
		cnt              int32
	)

	fallback := mw.LockFallback(func(ctx context.Context, key, value interface{}) error {
		time.Sleep(time.Second)
		atomic.AddInt32(&cnt, 1)
		*value.(*string) = fmt.Sprintf("value:%v", cnt)

		return nil
	})

	wg.Add(2)

	go func() {
		assert.Nil(t, fallback(ctx, 1, &val1))
		wg.Done()
	}()
	go func() {
		assert.Nil(t, fallback(ctx, 1, &val2))
		wg.Done()
	}()
	wg.Wait()

	require.Equal(t, "value:1", val1)
	require.Equal(t, "value:1", val2)

	assert.Nil(t, fallback(ctx, 1, &val3))
	require.Equal(t, "value:2", val3)
}

func TestLockFallback_Error(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	var (
		val1, val2, val3 string
		wg               sync.WaitGroup
		cnt              int32
	)

	fallback := mw.LockFallback(func(ctx context.Context, key, value interface{}) error {
		time.Sleep(time.Second)
		atomic.AddInt32(&cnt, 1)

		return fmt.Errorf("%w:%v", errFallback, cnt)
	})

	wg.Add(2)

	go func() {
		assert.EqualError(t, fallback(ctx, 1, &val1), "fallback error:1")
		wg.Done()
	}()
	go func() {
		assert.EqualError(t, fallback(ctx, 1, &val2), "fallback error:1")
		wg.Done()
	}()
	wg.Wait()

	require.Empty(t, val1)
	require.Empty(t, val2)

	assert.EqualError(t, fallback(ctx, 1, val3), "fallback error:2")
	require.Empty(t, val3)
}

func TestWithLockGetter(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	key1 := "getter:key1"

	var cnt int32

	prov := test.NewProviderMock(t,
		test.WithGet(cacheGetMiss),
		test.WithSet(cacheSetMiss(
			map[interface{}]test.User{
				key1: test.NewUser(1),
			},
		)),
	)
	c := cache.New(prov,
		cache.WithDataOption(
			cache.WithNamespace("gn", ":"),
		),
		mw.WithLockGetter(
			func(ctx context.Context, key interface{}) (interface{}, error) {
				atomic.AddInt32(&cnt, 1)
				time.Sleep(time.Second / 2)
				switch key.(string) {
				case key1:
					return test.NewUser(1), nil
				case "key2":
					return test.NewUser(2), nil
				}

				return nil, fmt.Errorf("%w: key '%v'", errKey, key)
			},
			mw.HandleByErr,
		))

	var (
		user1, user2 test.User
		wg           sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		require.Nil(t, c.Get(ctx, key1, &user1))
		wg.Done()
	}()
	go func() {
		require.Nil(t, c.Get(ctx, key1, &user2))
		wg.Done()
	}()

	wg.Wait()

	require.Equal(t, test.NewUser(1), user1)
	require.Equal(t, test.NewUser(1), user2)
	require.Equal(t, int32(1), cnt)
}

func cacheGetMiss(t *testing.T) func(ctx context.Context, item *cache.Item) error {
	t.Helper()

	return func(ctx context.Context, item *cache.Item) error {
		return cache.ErrCacheMiss
	}
}

func cacheSetMiss(items map[interface{}]test.User) func(t *testing.T) func(ctx context.Context, item *cache.Item) error {
	return func(t *testing.T) func(ctx context.Context, item *cache.Item) error {
		t.Helper()

		return func(ctx context.Context, item *cache.Item) error {
			if value, ok := items[item.Key.Key]; ok {
				require.Equal(t, &value, item.Value)

				return nil
			}

			t.Errorf("unexpected key %v", item.Key.String())

			return nil
		}
	}
}
