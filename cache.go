package cache

import (
	"context"
	"time"
)

// Configure configure cache.
type Configure func(*Cache)

// WithDataOption sets cache default data options.
func WithDataOption(do ...Option) Configure {
	return func(c *Cache) {
		factory := c.dataFactory
		c.dataFactory = func(key, value interface{}, opts ...Option) *Item {
			return factory(key, value, append(do, opts...)...)
		}
	}
}

// WithDefaultNamespace sets cache default namespace.
func WithDefaultNamespace(ns, separator string) Configure {
	return WithDataOption(WithNamespace(ns, separator))
}

// WithDefaultTTL sets cache default ttl.
func WithDefaultTTL(ttl time.Duration) Configure {
	return WithDataOption(WithTTL(ttl))
}

// WithHandleSet add a handler for the set operation.
func WithHandleSet(m ...Handle) Configure {
	return WithHandleOperation(OperationSet, m...)
}

// WithHandleGet add a handler for the get operation.
func WithHandleGet(m ...Handle) Configure {
	return WithHandleOperation(OperationGet, m...)
}

// WithHandleDelete add a handler for the delete operation.
func WithHandleDelete(m ...Handle) Configure {
	return WithHandleOperation(OperationDelete, m...)
}

// WithHandleOperation add a handler for the operation.
func WithHandleOperation(op string, m ...Handle) Configure {
	handle := ChainHandle(m...)

	return WithMiddleware(func(ctx context.Context, operation string, item *Item, next Provider) error {
		if operation == op {
			return handle(ctx, op, item, next)
		}

		return next(ctx, operation, item)
	})
}

// WithMiddleware sets middleware to provider.
func WithMiddleware(mw ...Handle) Configure {
	return func(c *Cache) {
		prov := c.provider
		c.provider = chain(prov, mw...)
	}
}

// New creates new cache by provider.
func New(prov Provider, opts ...Configure) *Cache {
	c := &Cache{
		provider:    prov,
		dataFactory: NewItem,
	}

	for _, o := range opts {
		o(c)
	}

	return c
}

// Cache base cache.
type Cache struct {
	dataFactory func(key, value interface{}, opts ...Option) *Item
	provider    Provider
}

func (c *Cache) With(opts ...Configure) *Cache {
	cache := &Cache{
		provider:    c.provider,
		dataFactory: c.dataFactory,
	}

	for _, o := range opts {
		o(cache)
	}

	return cache
}

func (c *Cache) Item(key, value interface{}, opts ...Option) *Item {
	return c.dataFactory(key, value, opts...)
}

func (c *Cache) Execute(ctx context.Context, operation string, key, value interface{}, opts ...Option) error {
	return c.provider(ctx, operation, c.Item(key, value, opts...))
}

// Set handles middlewares and sets value by key and options.
func (c *Cache) Set(ctx context.Context, key, value interface{}, opts ...Option) error {
	return c.Execute(ctx, OperationSet, key, value, opts...)
}

// Get handles middlewares and gets value by key and options.
func (c *Cache) Get(ctx context.Context, key, value interface{}, opts ...Option) error {
	return c.Execute(ctx, OperationGet, key, value, opts...)
}

// Delete handles middlewares and delete value by key and options.
func (c *Cache) Delete(ctx context.Context, key interface{}, opts ...Option) error {
	return c.Execute(ctx, OperationDelete, key, nil, opts...)
}
