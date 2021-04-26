package memory_test

import (
	"testing"

	"gitoa.ru/go-4devs/cache/provider/memory"
	"gitoa.ru/go-4devs/cache/test"
)

func TestMap(t *testing.T) {
	t.Parallel()
	test.RunSute(t, memory.NewMap())
}

func TestMapShard(t *testing.T) {
	t.Parallel()
	test.RunSute(t, memory.NewMapShard())
}
