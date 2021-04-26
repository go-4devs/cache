package item

import (
	"fmt"
	"time"

	"gitoa.ru/go-4devs/cache"
)

//go:generate easyjson

//easyjson:json
type expiredByte struct {
	Data    []byte    `json:"d"`
	Expired time.Time `json:"e"`
}

func MarshalExpired(item *cache.Item) ([]byte, error) {
	var (
		e   expiredByte
		err error
	)

	e.Data, err = item.Marshal()
	if err != nil {
		return nil, fmt.Errorf("failed marshal expired: %w", err)
	}

	if item.TTL > 0 {
		e.Expired = item.Expired()
	}

	return e.MarshalJSON()
}

func UnmarshalExpired(item *cache.Item, d []byte) error {
	var e expiredByte

	if err := e.UnmarshalJSON(d); err != nil {
		return err
	}

	if !e.Expired.IsZero() {
		item.TTL = time.Until(e.Expired)
	}

	if item.IsExpired() {
		return cache.ErrCacheExpired
	}

	if err := item.Unmarshal(e.Data); err != nil {
		return fmt.Errorf("failed unmarshal expired: %w", err)
	}

	return nil
}
