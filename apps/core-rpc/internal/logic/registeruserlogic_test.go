package logic

import (
	"strings"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNormalizeRegistrationInput(t *testing.T) {
	tests := []struct {
		name       string
		email      string
		password   string
		nickname   string
		wantEmail  string
		wantStatus codes.Code
	}{
		{
			name:       "normalizes email and nickname",
			email:      "  USER@Example.com ",
			password:   "123456",
			nickname:   " tester ",
			wantEmail:  "user@example.com",
			wantStatus: codes.OK,
		},
		{
			name:       "rejects display-name address",
			email:      "Tester <user@example.com>",
			password:   "123456",
			nickname:   "tester",
			wantStatus: codes.InvalidArgument,
		},
		{
			name:       "accepts six-character unicode password",
			email:      "user@example.com",
			password:   "密码测试一二",
			nickname:   "tester",
			wantEmail:  "user@example.com",
			wantStatus: codes.OK,
		},
		{
			name:       "rejects password over bcrypt limit",
			email:      "user@example.com",
			password:   strings.Repeat("a", 73),
			nickname:   "tester",
			wantStatus: codes.InvalidArgument,
		},
		{
			name:       "accepts 64-character unicode nickname",
			email:      "user@example.com",
			password:   "123456",
			nickname:   strings.Repeat("名", 64),
			wantEmail:  "user@example.com",
			wantStatus: codes.OK,
		},
		{
			name:       "rejects 65-character unicode nickname",
			email:      "user@example.com",
			password:   "123456",
			nickname:   strings.Repeat("名", 65),
			wantStatus: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email, _, _, err := normalizeRegistrationInput(tt.email, tt.password, tt.nickname)
			if got := status.Code(err); got != tt.wantStatus {
				t.Fatalf("status.Code(err) = %v, want %v; err = %v", got, tt.wantStatus, err)
			}
			if email != tt.wantEmail {
				t.Fatalf("email = %q, want %q", email, tt.wantEmail)
			}
		})
	}
}
