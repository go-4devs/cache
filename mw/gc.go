package mw

import (
	"context"
	"fmt"
	"sync"
	"time"

	"gitoa.ru/go-4devs/cache"
)

type key struct {
	key       interface{}
	ctxPrefix string
}

func (k key) Value() interface{} {
	return k.key
}

func (k key) String() string {
	return fmt.Sprint(k.ctxPrefix, k.key)
}

// WithClearByContext clear cache if context done.
func WithClearByContext(ctxKey interface{}) cache.Configure {
	operation := func(ctx context.Context, op string, item *cache.Item, next cache.Provider) error {
		ctxPrefix, ok := ctx.Value(ctxKey).(string)
		if !ok {
			return fmt.Errorf("%w: must be unique ctx key", cache.ErrKeyNotValid)
		}

		k := item.Key.Key
		item.Key.Key = key{
			key:       k,
			ctxPrefix: ctxPrefix,
		}

		return next(ctx, op, item)
	}

	return cache.WithMiddleware(
		func(ctx context.Context, op string, item *cache.Item, next cache.Provider) error {
			if op == cache.OperationSet {
				go func(ctx context.Context, item *cache.Item) {
					<-ctx.Done()
					_ = next(ctx, cache.OperationDelete, item)
				}(ctx, item)
			}

			return operation(ctx, op, item, next)
		})
}

// WithClearByTTL clear cache by key after ttl.
func WithClearByTTL() cache.Configure {
	keys := make(map[cache.Key]*time.Timer)
	mu := sync.Mutex{}

	return cache.WithHandleSet(func(ctx context.Context, op string, item *cache.Item, next cache.Provider) error {
		if item.TTL > 0 {
			go func() {
				mu.Lock()
				defer mu.Unlock()
				if t, ok := keys[item.Key]; ok {
					t.Reset(item.TTL)
				} else {
					keys[item.Key] = time.AfterFunc(item.TTL, func() {
						_ = next(ctx, cache.OperationDelete, item)
						delete(keys, item.Key)
					})
				}
			}()
		}

		return next(ctx, op, item)
	})
}
