package redis

import (
	"context"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// NewPool creates redigo pool.
func NewPool(pool *redis.Pool) func(context.Context) (Conn, error) {
	return func(ctx context.Context) (Conn, error) {
		conn, err := pool.GetContext(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed get connect: %w", err)
		}

		return conn, nil
	}
}
