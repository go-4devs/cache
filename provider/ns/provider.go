package ns

import (
	"context"
	"errors"

	"gitoa.ru/go-4devs/cache"
)

var ErrProviderNotFound = errors.New("provider not found")

func New(providers map[string]cache.Provider) cache.Provider {
	return func(ctx context.Context, operation string, item *cache.Item) error {
		if prov, ok := providers[item.Key.Prefix]; ok {
			item.Key.Prefix = ""

			return prov(ctx, operation, item)
		}

		return ErrProviderNotFound
	}
}
