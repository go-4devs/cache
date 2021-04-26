package cache

import (
	"errors"
	"fmt"
)

// Cached errors.
var (
	ErrCacheMiss          = errors.New("cache miss")
	ErrCacheExpired       = errors.New("cache expired")
	ErrSourceNotValid     = errors.New("source is not valid")
	ErrKeyNotValid        = errors.New("key is not valid")
	ErrTargetNil          = errors.New("target is nil")
	ErrOperationNotAllwed = errors.New("operation not allowed")
)

var _ error = NewErrorTarget(nil)

// NewErrorTarget creates new target error.
func NewErrorTarget(target interface{}) ErrorTarget {
	return ErrorTarget{target: target}
}

// ErrorTarget errs target is not a settable.
type ErrorTarget struct {
	target interface{}
}

// ErrorTarget errors.
func (e ErrorTarget) Error() string {
	return fmt.Sprintf("target is not a settable %T", e.target)
}
