package xcontext

import (
	"context"
	"testing"
)

func TestUserID(t *testing.T) {
	tests := []struct {
		name   string
		ctx    context.Context
		want   int64
		wantOK bool
	}{
		{
			name:   "project context value",
			ctx:    WithUserID(context.Background(), 123),
			want:   123,
			wantOK: true,
		},
		{
			name:   "go-zero jwt context value",
			ctx:    context.WithValue(context.Background(), "user_id", float64(456)),
			want:   456,
			wantOK: true,
		},
		{
			name: "fractional jwt value",
			ctx:  context.WithValue(context.Background(), "user_id", 1.5),
		},
		{
			name: "missing value",
			ctx:  context.Background(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := UserID(tt.ctx)
			if got != tt.want || ok != tt.wantOK {
				t.Fatalf("UserID() = (%d, %v), want (%d, %v)", got, ok, tt.want, tt.wantOK)
			}
		})
	}
}
