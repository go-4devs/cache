package cache_test

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gitoa.ru/go-4devs/cache"
	"gitoa.ru/go-4devs/cache/test"
)

func TestCache_Get(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	pro := test.NewProviderMock(t, test.WithGet(func(t *testing.T) func(ctx context.Context, item *cache.Item) error {
		t.Helper()

		return func(ctx context.Context, d *cache.Item) error {
			u := test.NewUser(1)

			return cache.TypeAssert(u, d.Value)
		}
	}))
	cache := cache.New(pro)

	var user test.User

	require.Nil(t, cache.Get(ctx, 1, &user))
	require.Equal(t, test.NewUser(1), user)
}

func TestCache_Set(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	pro := test.NewProviderMock(t, test.WithSet(func(t *testing.T) func(ctx context.Context, item *cache.Item) error {
		t.Helper()

		return func(ctx context.Context, d *cache.Item) error {
			require.Equal(t, 1, d.Key.Key)
			require.Equal(t, test.NewUser(1), d.Value)
			require.Equal(t, "1", d.Key.String())

			return nil
		}
	}))
	cache := cache.New(pro)

	require.Nil(t, cache.Set(ctx, 1, test.NewUser(1)))
}

func TestCache_Delete(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	pro := test.NewProviderMock(t, test.WithDelete(func(t *testing.T) func(ctx context.Context, item *cache.Item) error {
		t.Helper()

		return func(ctx context.Context, d *cache.Item) error {
			require.Equal(t, 1, d.Key.Key)
			require.Empty(t, d.Value)
			require.Equal(t, "1", d.Key.String())

			return nil
		}
	}))
	cache := cache.New(pro)

	require.Nil(t, cache.Delete(ctx, 1))
}

func TestCache_Get_withMiddleware(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	pro := test.NewProviderMock(t,
		test.WithSet(func(t *testing.T) func(ctx context.Context, item *cache.Item) error {
			t.Helper()

			return func(ctx context.Context, d *cache.Item) error {
				require.Equal(t, 2, d.Key.Key)
				require.Equal(t, test.NewUser(2), d.Value)
				require.Equal(t, "mw_prefix::_2", d.Key.String())
				require.Equal(t, time.Minute, d.TTL)

				return nil
			}
		}),
		test.WithGet(func(t *testing.T) func(ctx context.Context, item *cache.Item) error {
			t.Helper()

			return func(ctx context.Context, d *cache.Item) error {
				require.Equal(t, 2, d.Key.Key)
				require.Equal(t, "mw_prefix----2", d.Key.String())

				return nil
			}
		}),
	)
	c := cache.New(pro,
		cache.WithDataOption(
			func(i *cache.Item) {
				i.Key.Prefix = "mw_prefix"
			},
			cache.WithTTL(time.Hour),
		),
		cache.WithHandleSet(
			func(ctx context.Context, op string, d *cache.Item, n cache.Provider) error {
				d.Key.Separator = "::"

				return n(ctx, op, d)
			}, func(ctx context.Context, op string, d *cache.Item, n cache.Provider) error {
				d.Key.Separator += "_"

				return n(ctx, op, d)
			}),
		cache.WithHandleGet(func(ctx context.Context, op string, d *cache.Item, n cache.Provider) error {
			d.Key.Separator = "----"

			return n(ctx, op, d)
		}),
	)

	var user test.User

	require.Nil(t, c.Set(ctx, 2, test.NewUser(2), cache.WithTTL(time.Minute)))
	require.Nil(t, c.Get(ctx, 2, &user))
}

func TestCacheWith(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	pro := test.NewProviderMock(t,
		test.WithGet(func(t *testing.T) func(ctx context.Context, item *cache.Item) error {
			t.Helper()

			return func(ctx context.Context, d *cache.Item) error {
				switch d.Key.Key.(int) {
				case 1:
					require.Equal(t, "ns:1", d.Key.String())
				case 2:
					require.Equal(t, "new_ns_2", d.Key.String())
				default:
					t.Errorf("key %v no allowed", d.Key.Key)
				}

				return nil
			}
		}),
		test.WithSet(func(t *testing.T) func(ctx context.Context, item *cache.Item) error {
			t.Helper()

			return func(ctx context.Context, d *cache.Item) error {
				switch d.Key.Key.(int) {
				case 1:
					require.Equal(t, time.Hour, d.TTL)
				case 2:
					require.Equal(t, time.Minute, d.TTL)
				default:
					t.Errorf("key %v no allowed", d.Key.Key)
				}

				return nil
			}
		}),
	)

	var (
		cntSetCache2, cntGetCache1 int32
		val1, val2                 string
	)

	cache1 := cache.New(pro,
		cache.WithHandleGet(func(ctx context.Context, operation string, item *cache.Item, next cache.Provider) error {
			atomic.AddInt32(&cntGetCache1, 1)

			return next(ctx, operation, item)
		}),
		cache.WithDataOption(
			cache.WithNamespace("ns", ":"),
			cache.WithTTL(time.Hour),
		),
	)
	cache2 := cache1.With(
		cache.WithHandleSet(func(ctx context.Context, operation string, item *cache.Item, next cache.Provider) error {
			atomic.AddInt32(&cntSetCache2, 1)

			return next(ctx, operation, item)
		}),
		cache.WithDataOption(
			cache.WithNamespace("new_ns", "_"),
			cache.WithTTL(time.Minute),
		),
	)

	require.NoError(t, cache1.Get(ctx, 1, &val1))
	require.NoError(t, cache2.Get(ctx, 2, &val2))

	require.NoError(t, cache1.Set(ctx, 1, val1))
	require.NoError(t, cache2.Set(ctx, 2, val2))

	require.Equal(t, int32(1), cntSetCache2)
	require.Equal(t, int32(2), cntGetCache1)
}
