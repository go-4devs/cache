package pebble_test

import (
	"testing"

	"gitoa.ru/go-4devs/cache/provider/pebble"
	"gitoa.ru/go-4devs/cache/test"
)

func TestPebble(t *testing.T) {
	t.Parallel()

	db, cl := test.PebbleDB()
	defer cl()

	test.RunSute(t, pebble.New(db))
}
