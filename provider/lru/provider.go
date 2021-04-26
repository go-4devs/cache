package lru

import (
	"context"
	"fmt"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"gitoa.ru/go-4devs/cache"
)

// New create new lru cache provider.
func New(client *lru.Cache) cache.Provider {
	return func(ctx context.Context, operation string, item *cache.Item) error {
		switch operation {
		case cache.OperationGet:
			val, ok := client.Get(item.Key)
			if !ok {
				return wrapErr(cache.ErrCacheMiss)
			}

			it, _ := val.(expired)

			if !it.ex.IsZero() {
				item.TTL = time.Until(it.ex)
			}

			if item.IsExpired() {
				return wrapErr(cache.ErrCacheExpired)
			}

			return wrapErr(cache.TypeAssert(it.value, item.Value))
		case cache.OperationSet:
			it := expired{
				value: item.Value,
				ex:    time.Time{},
			}
			if item.TTL > 0 {
				it.ex = item.Expired()
			}

			_ = client.Add(item.Key, it)

			return nil
		case cache.OperationDelete:
			_ = client.Remove(item.Key)

			return nil
		}

		return wrapErr(cache.ErrOperationNotAllwed)
	}
}

type expired struct {
	ex    time.Time
	value interface{}
}

func wrapErr(err error) error {
	if err != nil {
		return fmt.Errorf("%w: lru", err)
	}

	return nil
}
