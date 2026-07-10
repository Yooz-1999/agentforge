// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package middleware

import (
	"net/http"

	"github.com/Yooz-1999/agentforge/pkg/auth"
)

type AccessAuthMiddleware struct {
	tokenManager *auth.TokenManager
}

func NewAccessAuthMiddleware(tokenManager *auth.TokenManager) *AccessAuthMiddleware {
	return &AccessAuthMiddleware{tokenManager: tokenManager}
}

func (m *AccessAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return auth.AccessTokenMiddleware(m.tokenManager, next)
}
