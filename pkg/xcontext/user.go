package xcontext

import "context"

type userIDKey struct{}

func WithUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

func UserID(ctx context.Context) (int64, bool) {
	val, ok := ctx.Value(userIDKey{}).(int64)
	return val, ok
}
