package test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitoa.ru/go-4devs/cache"
)

const (
	expire     = time.Second
	waitExpire = expire * 2
)

// Option configure sute.
type Option func(*ProviderSuite)

// WithExpire sets expired errors.
func WithExpire(err error) Option {
	return func(ps *ProviderSuite) { ps.expire = err }
}

func WithWaitGet(f func()) Option {
	return func(ps *ProviderSuite) { ps.waitGet = f }
}

// RunSute run test by provider.
func RunSute(t *testing.T, provider cache.Provider, opts ...Option) {
	t.Helper()

	cs := &ProviderSuite{
		provider: provider,
		expire:   cache.ErrCacheExpired,
		waitGet:  func() {},
	}

	for _, o := range opts {
		o(cs)
	}

	suite.Run(t, cs)
}

// ProviderSuite for testing providers.
type ProviderSuite struct {
	provider cache.Provider
	expire   error
	waitGet  func()
	suite.Suite
}

// TestGet tested get.
func (s *ProviderSuite) TestGet() {
	s.T().Parallel()

	ctx := context.Background()

	var val string

	require.Nil(s.T(), s.provider(ctx, cache.OperationSet, cache.NewItem("get", "some value")))
	s.waitGet()
	require.Nil(s.T(), s.provider(ctx, cache.OperationGet, cache.NewItem("get", &val)))
	require.Equal(s.T(), "some value", val)

	var user User

	cachedUser := NewUser(1)

	require.Nil(s.T(), s.provider(ctx, cache.OperationSet, cache.NewItem("get_user", cachedUser)))
	s.waitGet()
	require.Nil(s.T(), s.provider(ctx, cache.OperationGet, cache.NewItem("get_user", &user)))
	require.Equal(s.T(), cachedUser, user)
}

// TestCacheMiss tested cache miss error.
func (s *ProviderSuite) TestCacheMiss() {
	s.T().Parallel()

	ctx := context.Background()

	require.True(s.T(),
		errors.Is(s.provider(ctx, cache.OperationGet, cache.NewItem("cache_miss", nil)), cache.ErrCacheMiss),
		"failed expect errors",
	)
}

// TestExpired tested error expired.
func (s *ProviderSuite) TestExpired() {
	s.T().Parallel()

	ctx := context.Background()

	var val string

	require.Nil(s.T(), s.provider(ctx, cache.OperationSet, cache.NewItem("expired", "some value", cache.WithTTL(expire))))
	time.Sleep(waitExpire)

	err := s.provider(ctx, cache.OperationGet, cache.NewItem("expired", nil))
	require.Truef(s.T(), errors.Is(err, s.expire), "failed expired error got:%s", err)
	require.Equal(s.T(), "", val)
}

// TestTTL tested set ttl.
func (s *ProviderSuite) TestTTL() {
	s.T().Parallel()

	ctx := context.Background()

	var val string

	require.Nil(s.T(), s.provider(ctx, cache.OperationSet, cache.NewItem("ttl", "some ttl value", cache.WithTTL(time.Hour))))
	s.waitGet()
	require.Nil(s.T(), s.provider(ctx, cache.OperationGet, cache.NewItem("ttl", &val)))
	require.Equal(s.T(), "some ttl value", val)
}

// TestDelete tested delete method.
func (s *ProviderSuite) TestDelete() {
	s.T().Parallel()

	ctx := context.Background()

	require.Nil(s.T(), s.provider(ctx, cache.OperationSet, cache.NewItem("delete:key", "some delete value")))
	require.Nil(s.T(), s.provider(ctx, cache.OperationDelete, cache.NewItem("delete:key", nil)))
	require.True(s.T(),
		errors.Is(s.provider(ctx, cache.OperationGet, cache.NewItem("cache_miss", nil)), cache.ErrCacheMiss),
		"failed delete errors",
	)
}
