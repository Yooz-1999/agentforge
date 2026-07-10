package auth

import (
	"testing"
	"time"

	"github.com/Yooz-1999/agentforge/pkg/constants"
	"github.com/golang-jwt/jwt/v4"
)

func TestTokenManagerUsesSeparateSecrets(t *testing.T) {
	manager := NewTokenManager("access-test-secret", "refresh-test-secret", 60, 120)

	accessToken, refreshToken, expiresIn, err := manager.GenerateTokenPair(12, "user@example.com", "tester")
	if err != nil {
		t.Fatalf("GenerateTokenPair() error = %v", err)
	}
	if expiresIn != 60 {
		t.Fatalf("GenerateTokenPair() expiresIn = %d, want 60", expiresIn)
	}

	refreshClaims, err := manager.ParseRefreshToken(refreshToken)
	if err != nil {
		t.Fatalf("ParseRefreshToken(refreshToken) error = %v", err)
	}
	if refreshClaims.UserID != 12 || refreshClaims.TokenType != constants.TokenTypeRefresh {
		t.Fatalf("ParseRefreshToken(refreshToken) claims = %+v", refreshClaims)
	}

	if _, err := manager.ParseRefreshToken(accessToken); err == nil {
		t.Fatal("ParseRefreshToken(accessToken) expected an error")
	}

	accessClaims, err := manager.ParseAccessToken(accessToken)
	if err != nil {
		t.Fatalf("ParseAccessToken(accessToken) error = %v", err)
	}
	if accessClaims.TokenType != constants.TokenTypeAccess {
		t.Fatalf("access token type = %q, want %q", accessClaims.TokenType, constants.TokenTypeAccess)
	}
}

func TestParseRefreshTokenRejectsUnexpectedAlgorithm(t *testing.T) {
	manager := NewTokenManager("access-test-secret", "refresh-test-secret", 60, 120)
	claims := UserClaims{
		UserID:    12,
		TokenType: constants.TokenTypeRefresh,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString([]byte("refresh-test-secret"))
	if err != nil {
		t.Fatalf("sign token error = %v", err)
	}

	if _, err := manager.ParseRefreshToken(token); err == nil {
		t.Fatal("ParseRefreshToken(HS512 token) expected an error")
	}
}
