package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Yooz-1999/agentforge/pkg/xcontext"
)

func TestAccessTokenMiddleware(t *testing.T) {
	manager := NewTokenManager("access-test-secret", "refresh-test-secret", 60, 120)
	accessToken, refreshToken, _, err := manager.GenerateTokenPair(42, "user@example.com", "tester")
	if err != nil {
		t.Fatalf("GenerateTokenPair() error = %v", err)
	}

	tests := []struct {
		name       string
		header     string
		wantStatus int
		wantUserID int64
	}{
		{
			name:       "accepts access token",
			header:     "Bearer " + accessToken,
			wantStatus: http.StatusNoContent,
			wantUserID: 42,
		},
		{
			name:       "rejects refresh token",
			header:     "Bearer " + refreshToken,
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "rejects malformed header",
			header:     "bad",
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "rejects missing header",
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/protected", nil)
			request.Header.Set("Authorization", tt.header)
			response := httptest.NewRecorder()

			handler := AccessTokenMiddleware(manager, func(w http.ResponseWriter, r *http.Request) {
				userID, ok := xcontext.UserID(r.Context())
				if !ok || userID != tt.wantUserID {
					t.Fatalf("UserID() = (%d, %v), want (%d, true)", userID, ok, tt.wantUserID)
				}
				w.WriteHeader(http.StatusNoContent)
			})
			handler(response, request)

			if response.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", response.Code, tt.wantStatus)
			}
			if got := request.Header.Get("Authorization"); got != "" {
				t.Fatalf("Authorization header was not removed: %q", got)
			}
		})
	}
}
