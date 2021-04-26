package cache

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"gitoa.ru/go-4devs/encoding"
)

var _ fmt.Stringer = (*Key)(nil)

// Option ffor the configuration item.
type Option func(*Item)

// WithNamespace sets prefix and separator.
func WithNamespace(prefix, sep string) Option {
	return func(d *Item) {
		d.Key.Prefix = prefix
		d.Key.Separator = sep
	}
}

// WithTTL sets ttl.
func WithTTL(ttl time.Duration) Option {
	return func(d *Item) {
		d.TTL = ttl
	}
}

// WithMarshal sets marshal and unmarshal.
func WithMarshal(unmarshal encoding.Unmarshal, marshal encoding.Marshal) Option {
	return func(d *Item) {
		d.unmarshal = unmarshal
		d.marshal = marshal
	}
}

// NewItem creates and configure new item.
func NewItem(key, value interface{}, opts ...Option) *Item {
	item := &Item{
		Key: Key{
			Key:       key,
			Prefix:    "",
			Separator: "",
		},
		Value:     value,
		TTL:       0,
		unmarshal: json.Unmarshal,
		marshal:   json.Marshal,
	}

	for _, opt := range opts {
		opt(item)
	}

	return item
}

// Item to pass to the provider.
type Item struct {
	Key       Key
	Value     interface{}
	TTL       time.Duration
	unmarshal encoding.Unmarshal
	marshal   encoding.Marshal
}

func (i *Item) With(key, val interface{}, opts ...Option) *Item {
	return NewItem(key, val, append(i.Options(), opts...)...)
}

// IsExpired checks expired item.
func (i *Item) IsExpired() bool {
	return i.TTL < 0
}

func (i *Item) Marshal() ([]byte, error) {
	return i.marshal(i.Value)
}

func (i *Item) Unmarshal(data []byte) error {
	return i.unmarshal(data, i.Value)
}

// Options gets item options.
func (i *Item) Options() []Option {
	opts := []Option{WithTTL(i.TTL), WithMarshal(i.unmarshal, i.marshal)}

	if i.Key.Prefix != "" {
		opts = append(opts, WithNamespace(i.Key.Prefix, i.Key.Separator))
	}

	return opts
}

// TTLInSecond to set the ttl in seconds.
func (i *Item) TTLInSecond(in int64) {
	i.TTL = time.Second * time.Duration(in)
}

// Expired get the time when the ttl is outdated.
func (i *Item) Expired() time.Time {
	return time.Now().Add(i.TTL)
}

// Key with prefix and separator.
type Key struct {
	Key       interface{}
	Prefix    string
	Separator string
}

func (k Key) Value() interface{} {
	if v, ok := k.Key.(interface{ Value() interface{} }); ok {
		return v.Value()
	}

	return k.Key
}

// String returns a formatted key.
func (k Key) String() string {
	if k.Prefix != "" {
		return fmt.Sprint(k.Prefix, k.Separator, k.Key)
	}

	switch v := k.Key.(type) {
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case string:
		return v
	default:
		return fmt.Sprint(v)
	}
}
