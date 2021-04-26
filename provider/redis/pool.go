package redis

import (
	"context"
	"fmt"

	"gitoa.ru/go-4devs/cache"
)

type Conn interface {
	Do(commandName string, args ...interface{}) (reply interface{}, err error)
	Send(commandName string, args ...interface{}) error
	Flush() error
	Close() error
}

// New creates new provider.
func New(pool func(context.Context) (Conn, error)) cache.Provider {
	return func(ctx context.Context, operation string, item *cache.Item) error {
		conn, err := pool(ctx)
		if err != nil {
			return wrapErr(err)
		}
		defer conn.Close()

		key := item.Key.String()

		switch operation {
		case cache.OperationGet:
			data, ttl, err := get(conn, key)
			if err != nil {
				return err
			}

			item.TTLInSecond(ttl)

			return wrapErr(item.Unmarshal(data))
		case cache.OperationSet:
			data, err := item.Marshal()
			if err != nil {
				return wrapErr(err)
			}

			return set(conn, key, data, int(item.TTL.Seconds()))
		case cache.OperationDelete:
			return del(conn, key)
		}

		return wrapErr(cache.ErrOperationNotAllwed)
	}
}

func get(conn Conn, key string) ([]byte, int64, error) {
	data, err := conn.Do("GET", key)
	if err != nil {
		return nil, 0, wrapErr(err)
	}

	if data == nil {
		return nil, 0, wrapErr(cache.ErrCacheMiss)
	}

	v, ok := data.([]byte)
	if !ok {
		return nil, 0, wrapErr(cache.ErrSourceNotValid)
	}

	expire, err := conn.Do("TTL", key)
	if err != nil {
		return v, 0, wrapErr(err)
	}

	ex, _ := expire.(int64)

	return v, ex, nil
}

func set(conn Conn, key string, data []byte, ttl int) error {
	if err := conn.Send("SET", key, data); err != nil {
		return wrapErr(err)
	}

	if ttl > 0 {
		if err := conn.Send("EXPIRE", key, ttl); err != nil {
			return wrapErr(err)
		}
	}

	if err := conn.Flush(); err != nil {
		return fmt.Errorf("failed flush then set %s by %w", key, conn.Flush())
	}

	return nil
}

func del(conn Conn, key string) error {
	if _, err := conn.Do("DEL", key); err != nil {
		return wrapErr(err)
	}

	return nil
}

func wrapErr(err error) error {
	if err != nil {
		return fmt.Errorf("%w: redis pool", err)
	}

	return nil
}
