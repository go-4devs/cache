package ristretto

import (
	"context"
	"errors"
	"fmt"

	"github.com/dgraph-io/ristretto"
	"gitoa.ru/go-4devs/cache"
)

var ErrSetValue = errors.New("failed set value")

type Option func(*setting)

func WithCost(cost int64) Option {
	return func(s *setting) {
		s.cost = cost
	}
}

type setting struct {
	cost int64
}

func New(retto *ristretto.Cache, opts ...Option) cache.Provider {
	s := setting{
		cost: 1,
	}

	for _, opt := range opts {
		opt(&s)
	}

	return func(ctx context.Context, operation string, item *cache.Item) error {
		var key interface{}
		if item.Key.Prefix != "" {
			key = item.Key.String()
		} else {
			key = item.Key.Key
		}

		switch operation {
		case cache.OperationGet:
			res, ok := retto.Get(key)
			if !ok {
				return fmt.Errorf("%w: ristretto", cache.ErrCacheMiss)
			}

			if err := cache.TypeAssert(res, item.Value); err != nil {
				return fmt.Errorf("failed assert type: %w", err)
			}

			return nil
		case cache.OperationDelete:
			retto.Del(key)

			return nil
		case cache.OperationSet:
			if ok := retto.SetWithTTL(key, item.Value, s.cost, item.TTL); !ok {
				return ErrSetValue
			}

			return nil
		}

		return cache.ErrOperationNotAllwed
	}
}
