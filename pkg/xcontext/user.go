package xcontext

import (
	"context"
	"math"
)

type userIDKey struct{}

func WithUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

func UserID(ctx context.Context) (int64, bool) {
	if val, ok := ctx.Value(userIDKey{}).(int64); ok {
		return val, true
	}

	val, ok := ctx.Value("user_id").(float64)
	if !ok || val <= 0 || val > math.MaxInt64 || math.Trunc(val) != val {
		return 0, false
	}

	return int64(val), true
}
