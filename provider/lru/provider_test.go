package lru_test

import (
	"testing"

	glru "github.com/hashicorp/golang-lru"
	"github.com/stretchr/testify/require"
	"gitoa.ru/go-4devs/cache/provider/lru"
	"gitoa.ru/go-4devs/cache/test"
)

func TestEncoding(t *testing.T) {
	t.Parallel()

	client, err := glru.New(10)
	require.Nil(t, err)
	test.RunSute(t, lru.New(client))
}
