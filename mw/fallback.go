package mw

import (
	"context"
	"fmt"
	"sync"

	"gitoa.ru/go-4devs/cache"
)

type Fallback func(ctx context.Context, key, value interface{}) error

type Getter func(ctx context.Context, key interface{}) (interface{}, error)

// HandleByErr checks if cache return err.
func HandleByErr(_ *cache.Item, err error) bool {
	return err != nil
}

// LockFallback locks run fallback by item key.
func LockFallback(fallback Fallback) Fallback {
	var mu sync.Mutex

	type entry struct {
		item interface{}
		err  error
	}

	keys := make(map[interface{}]chan entry)

	return func(ctx context.Context, key, value interface{}) error {
		mu.Lock()
		if _, ok := keys[key]; !ok {
			keys[key] = make(chan entry, 1)
			mu.Unlock()

			err := fallback(ctx, key, value)
			keys[key] <- entry{
				item: value,
				err:  err,
			}

			defer func() {
				close(keys[key])
				delete(keys, key)
			}()

			return err
		}
		mu.Unlock()

		entry := <-keys[key]
		if entry.err != nil {
			return entry.err
		}

		if err := cache.TypeAssert(entry.item, value); err != nil {
			return fmt.Errorf("%w: assert value", err)
		}

		return nil
	}
}

// WithFallback sets fallback when cache handle success and set result in cache.
func WithFallback(fallback Fallback, isHandleFallback func(*cache.Item, error) bool) cache.Configure {
	return cache.WithHandleGet(func(ctx context.Context, op string, item *cache.Item, next cache.Provider) error {
		err := next(ctx, op, item)
		if isHandleFallback(item, err) {
			if ferr := fallback(ctx, item.Key.Key, item.Value); ferr != nil {
				return ferr
			}

			return next(ctx, cache.OperationSet, item)
		}

		return err
	})
}

// WithLockGetter sets values from getter when cache handle success and set result in cache.
func WithLockGetter(getter Getter, isHandle func(*cache.Item, error) bool) cache.Configure {
	var mu sync.Mutex

	type entry struct {
		value interface{}
		err   error
	}

	keys := make(map[cache.Key]chan entry)

	return cache.WithHandleGet(func(ctx context.Context, op string, item *cache.Item, next cache.Provider) error {
		if err := next(ctx, op, item); !isHandle(item, err) {
			return err
		}

		mu.Lock()
		if _, ok := keys[item.Key]; !ok {
			keys[item.Key] = make(chan entry, 1)
			mu.Unlock()
			value, gerr := getter(ctx, item.Key.Value())
			keys[item.Key] <- entry{
				value: value,
				err:   gerr,
			}

			defer func() {
				close(keys[item.Key])
				delete(keys, item.Key)
			}()
			if gerr != nil {
				return gerr
			}

			if err := cache.TypeAssert(value, item.Value); err != nil {
				return fmt.Errorf("lock failed assert type: %w", err)
			}

			return nil
		}
		mu.Unlock()

		entry := <-keys[item.Key]
		if entry.err != nil {
			return entry.err
		}

		if err := cache.TypeAssert(entry.value, item.Value); err != nil {
			return fmt.Errorf("lock failed assert type: %w", err)
		}

		return next(ctx, cache.OperationSet, item)
	})
}
