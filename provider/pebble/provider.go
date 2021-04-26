package pebble

import (
	"context"
	"errors"
	"fmt"

	"github.com/cockroachdb/pebble"
	"gitoa.ru/go-4devs/cache"
	"gitoa.ru/go-4devs/cache/item"
)

func New(db *pebble.DB) cache.Provider {
	return func(ctx context.Context, operation string, i *cache.Item) error {
		key := []byte(i.Key.String())

		switch operation {
		case cache.OperationGet:
			val, cl, err := db.Get([]byte(i.Key.String()))
			if err != nil {
				if errors.Is(err, pebble.ErrNotFound) {
					return wrapErr(cache.ErrCacheMiss)
				}

				return wrapErr(err)
			}

			defer func() {
				_ = cl.Close()
			}()

			return wrapErr(item.UnmarshalExpired(i, val))
		case cache.OperationSet:
			b, err := item.MarshalExpired(i)
			if err != nil {
				return wrapErr(err)
			}

			return wrapErr(db.Set(key, b, pebble.Sync))
		case cache.OperationDelete:
			return wrapErr(db.Delete(key, pebble.Sync))
		}

		return wrapErr(cache.ErrOperationNotAllwed)
	}
}

func wrapErr(err error) error {
	if err != nil {
		return fmt.Errorf("%w: pebble", err)
	}

	return nil
}
