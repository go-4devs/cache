package test

import (
	"context"
	"fmt"
	"testing"

	"gitoa.ru/go-4devs/cache"
)

var _ cache.Provider = NewProviderMock(&testing.T{})

// OptionMock configure mock.
type OptionMock func(*ProviderMock)

// WithDelete sets delete method.
func WithDelete(f func(t *testing.T) func(ctx context.Context, item *cache.Item) error) OptionMock {
	return func(pm *ProviderMock) { pm.operations[cache.OperationDelete] = f }
}

// WithGet sets get method.
func WithGet(f func(t *testing.T) func(ctx context.Context, item *cache.Item) error) OptionMock {
	return func(pm *ProviderMock) { pm.operations[cache.OperationGet] = f }
}

// WithSet sets set method.
func WithSet(f func(t *testing.T) func(ctx context.Context, item *cache.Item) error) OptionMock {
	return func(pm *ProviderMock) { pm.operations[cache.OperationSet] = f }
}

// NewProviderMock create new mock provider.
func NewProviderMock(t *testing.T, opts ...OptionMock) cache.Provider {
	t.Helper()

	pm := &ProviderMock{
		t:          t,
		operations: make(map[string]func(t *testing.T) func(ctx context.Context, item *cache.Item) error, 3),
	}

	for _, o := range opts {
		o(pm)
	}

	return func(ctx context.Context, operation string, item *cache.Item) error {
		if m, ok := pm.operations[operation]; ok {
			return m(pm.t)(ctx, item)
		}

		return fmt.Errorf("%w: %s", cache.ErrOperationNotAllwed, operation)
	}
}

// ProviderMock mock.
type ProviderMock struct {
	t          *testing.T
	operations map[string]func(t *testing.T) func(ctx context.Context, item *cache.Item) error
}
