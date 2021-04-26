package mw

import (
	"context"
	"errors"
	"time"

	"gitoa.ru/go-4devs/cache"
)

// Metrics interface for middleware.
type Metrics interface {
	Hit(label string)
	Miss(label string)
	Expired(label string)
	Err(label string, operation string)
	Observe(label string, operation string, start time.Time)
}

// LabelName sets static name.
func LabelName(name string) func(ctx context.Context, item *cache.Item) string {
	return func(_ context.Context, _ *cache.Item) string {
		return name
	}
}

// LabelPreficKey gets lebale by item prefix.
func LabelPreficKey(ctx context.Context, item *cache.Item) string {
	return item.Key.Prefix
}

// WithMetrics cache middleware metrics.
func WithMetrics(m Metrics, labelCallback func(ctx context.Context, item *cache.Item) string) cache.Configure {
	return cache.WithMiddleware(
		func(ctx context.Context, op string, item *cache.Item, next cache.Provider) error {
			label := labelCallback(ctx, item)
			start := time.Now()
			err := next(ctx, op, item)
			m.Observe(label, op, start)
			if err != nil {
				switch {
				case errors.Is(err, cache.ErrCacheMiss):
					m.Miss(label)
				case errors.Is(err, cache.ErrCacheExpired):
					m.Expired(label)
				default:
					m.Err(label, op)
				}

				return err
			}

			if op == cache.OperationGet {
				m.Hit(label)
			}

			return nil
		},
	)
}
