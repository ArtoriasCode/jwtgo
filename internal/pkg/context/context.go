package ctxutil

import (
	"context"
	"fmt"
)

type contextKey string

func WithPairs(base context.Context, pairs map[string]any) context.Context {
	for key, value := range pairs {
		if value == nil {
			continue
		}
		base = context.WithValue(base, contextKey(key), value)
	}

	return base
}

func GetValue(ctx context.Context, key string) (any, error) {
	value := ctx.Value(contextKey(key))
	if value == nil {
		return nil, fmt.Errorf("missing %s in context", key)
	}

	return value, nil
}

func GetTyped[T any](ctx context.Context, key string) (T, error) {
	var zero T
	value := ctx.Value(contextKey(key))
	if value == nil {
		return zero, fmt.Errorf("missing %s in context", key)
	}

	casted, ok := value.(T)
	if !ok {
		return zero, fmt.Errorf("invalid type for %s in context: expected %T", key, zero)
	}

	return casted, nil
}
