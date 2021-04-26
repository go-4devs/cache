package cache

import (
	"context"
)

// Provider for different types of cache, in memory, lru, redis.
type Provider func(ctx context.Context, operation string, item *Item) error

// Handle middleware before/after provider.
type Handle func(ctx context.Context, operation string, item *Item, next Provider) error

// ChainHandle chain handle middleware.
func ChainHandle(handle ...Handle) Handle {
	if n := len(handle); n > 1 {
		lastI := n - 1

		return func(ctx context.Context, operation string, item *Item, next Provider) error {
			var (
				chainHandler func(context.Context, string, *Item) error
				curI         int
			)

			chainHandler = func(currentCtx context.Context, currentOperation string, currentData *Item) error {
				if curI == lastI {
					return next(currentCtx, currentOperation, currentData)
				}
				curI++
				err := handle[curI](currentCtx, currentOperation, currentData, chainHandler)
				curI--

				return err
			}

			return handle[0](ctx, operation, item, chainHandler)
		}
	}

	return handle[0]
}

func chain(init Provider, handleFunc ...Handle) Provider {
	if len(handleFunc) > 0 {
		return func(ctx context.Context, operation string, item *Item) error {
			return ChainHandle(handleFunc...)(ctx, operation, item, init)
		}
	}

	return init
}
