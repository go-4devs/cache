package memory

import (
	"context"
	"fmt"
	"sync"
	"time"

	"gitoa.ru/go-4devs/cache"
)

// NewEncoding create new provider.
func NewEncoding() cache.Provider {
	items := make(map[cache.Key]encodedEntry)
	mu := sync.RWMutex{}

	return func(ctx context.Context, operation string, item *cache.Item) error {
		switch operation {
		case cache.OperationSet:
			i, err := newEncodedEntry(item)
			if err != nil {
				return err
			}

			mu.Lock()
			items[item.Key] = i
			mu.Unlock()

			return nil
		case cache.OperationDelete:
			mu.Lock()
			delete(items, item.Key)
			mu.Unlock()

			return nil
		case cache.OperationGet:
			mu.RLock()
			i, ok := items[item.Key]
			mu.RUnlock()

			if !ok {
				return wrapErr(cache.ErrCacheMiss)
			}

			return resolveEncodedEntry(i, item)
		}

		return wrapErr(cache.ErrOperationNotAllwed)
	}
}

type encodedEntry struct {
	data    []byte
	expired time.Time
}

func wrapErr(err error) error {
	if err != nil {
		return fmt.Errorf("%w: encoding", err)
	}

	return nil
}

func newEncodedEntry(item *cache.Item) (encodedEntry, error) {
	var (
		e   encodedEntry
		err error
	)

	e.data, err = item.Marshal()
	if err != nil {
		return e, wrapErr(err)
	}

	if item.TTL > 0 {
		e.expired = item.Expired()
	}

	return e, nil
}

func resolveEncodedEntry(e encodedEntry, item *cache.Item) error {
	if !e.expired.IsZero() {
		item.TTL = time.Until(e.expired)
	}

	if item.IsExpired() {
		return wrapErr(cache.ErrCacheExpired)
	}

	return wrapErr(item.Unmarshal(e.data))
}
