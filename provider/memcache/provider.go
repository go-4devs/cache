package memcache

import (
	"context"
	"errors"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	"gitoa.ru/go-4devs/cache"
)

// New memcache provider.
func New(client *memcache.Client) cache.Provider {
	return func(ctx context.Context, operation string, item *cache.Item) error {
		key := item.Key.String()

		switch operation {
		case cache.OperationGet:
			ci, err := client.Get(item.Key.String())

			switch {
			case errors.Is(err, memcache.ErrCacheMiss):
				return wrapErr(cache.ErrCacheMiss)
			case errors.Is(err, memcache.ErrMalformedKey):
				return wrapErr(cache.ErrKeyNotValid)
			case err != nil:
				return wrapErr(err)
			}

			return wrapErr(item.Unmarshal(ci.Value))
		case cache.OperationSet:
			data, err := item.Marshal()
			if err != nil {
				return wrapErr(err)
			}

			return wrapErr(client.Set(&memcache.Item{Key: key, Flags: 0, Value: data, Expiration: int32(item.TTL.Seconds())}))
		case cache.OperationDelete:
			return wrapErr(client.Delete(key))
		}

		return wrapErr(cache.ErrOperationNotAllwed)
	}
}

func wrapErr(err error) error {
	if err != nil {
		return fmt.Errorf("%w: memcache", err)
	}

	return nil
}
