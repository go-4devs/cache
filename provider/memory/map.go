package memory

import (
	"context"
	"fmt"
	"hash/crc64"
	"sync"
	"time"

	"gitoa.ru/go-4devs/cache"
)

const defaultShards = 255

// NewMap creates new map cache.
func NewMap() cache.Provider {
	m := sync.Map{}

	return func(ctx context.Context, op string, item *cache.Item) error {
		switch op {
		case cache.OperationDelete:
			m.Delete(item.Key)

			return nil
		case cache.OperationSet:
			m.Store(item.Key, newEntry(item))

			return nil
		case cache.OperationGet:
			e, ok := m.Load(item.Key)
			if !ok {
				return fmt.Errorf("%w: map", cache.ErrCacheMiss)
			}

			return resolveEntry(e.(entry), item)
		}

		return fmt.Errorf("%w: map", cache.ErrOperationNotAllwed)
	}
}

func resolveEntry(e entry, item *cache.Item) error {
	if !e.expired.IsZero() {
		item.TTL = time.Until(e.expired)
	}

	if item.IsExpired() {
		return fmt.Errorf("%w: map", cache.ErrCacheExpired)
	}

	if err := cache.TypeAssert(e.data, item.Value); err != nil {
		return fmt.Errorf("%w: map", err)
	}

	return nil
}

func newEntry(item *cache.Item) entry {
	e := entry{data: item.Value}

	if item.TTL > 0 {
		e.expired = item.Expired()
	}

	return e
}

type entry struct {
	data    interface{}
	expired time.Time
}

type settings struct {
	numShards  uint64
	hashString func(in cache.Key) uint64
}
type Option func(*settings)

func WithNumShards(num uint64) Option {
	return func(s *settings) {
		s.numShards = num
	}
}

func WithHashKey(f func(in cache.Key) uint64) Option {
	return func(s *settings) {
		s.hashString = f
	}
}

//nolint: gochecknoglobals
var table = crc64.MakeTable(crc64.ISO)

func hashString(in cache.Key) uint64 {
	switch k := in.Key.(type) {
	case int64:
		return uint64(k)
	case int32:
		return uint64(k)
	case int:
		return uint64(k)
	case uint64:
		return k
	case uint32:
		return uint64(k)
	case uint:
		return uint64(k)
	default:
		return crc64.Checksum([]byte(in.String()), table)
	}
}

func NewMapShard(opts ...Option) cache.Provider {
	s := settings{
		numShards:  defaultShards,
		hashString: hashString,
	}

	for _, opt := range opts {
		opt(&s)
	}

	items := make([]*sync.Map, s.numShards)
	for i := range items {
		items[i] = &sync.Map{}
	}

	return func(ctx context.Context, operation string, item *cache.Item) error {
		idx := s.hashString(item.Key)

		switch operation {
		case cache.OperationDelete:
			items[idx%s.numShards].Delete(item.Key)

			return nil
		case cache.OperationSet:
			items[idx%s.numShards].Store(item.Key, newEntry(item))

			return nil
		case cache.OperationGet:
			e, ok := items[idx%s.numShards].Load(item.Key)
			if !ok {
				return wrapShardErr(cache.ErrCacheMiss)
			}

			return resolveEntry(e.(entry), item)
		}

		return wrapShardErr(cache.ErrOperationNotAllwed)
	}
}

func wrapShardErr(err error) error {
	return fmt.Errorf("%w: memory shards", err)
}
