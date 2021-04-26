package cache

import "context"

// available operation.
const (
	OperationGet    = "get"
	OperationSet    = "set"
	OperationDelete = "delete"
)

// OperationProvider creating a provider based on available operations.
func OperationProvider(prov map[string]func(ctx context.Context, item *Item) error) Provider {
	return func(ctx context.Context, operation string, item *Item) error {
		if method, ok := prov[operation]; ok {
			return method(ctx, item)
		}

		return ErrOperationNotAllwed
	}
}
