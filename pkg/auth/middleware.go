package auth

import (
	"net/http"
	"strings"

	"github.com/Yooz-1999/agentforge/pkg/xcontext"
)

func AccessTokenMiddleware(manager *TokenManager, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		r.Header.Del("Authorization")

		parts := strings.Fields(authorization)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		claims, err := manager.ParseAccessToken(parts[1])
		if err != nil || claims.UserID <= 0 {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		ctx := xcontext.WithUserID(r.Context(), claims.UserID)
		next(w, r.WithContext(ctx))
	}
}
