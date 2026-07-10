package auth

import (
	"errors"
	"time"

	"github.com/Yooz-1999/agentforge/pkg/constants"
	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	UserID    int64  `json:"user_id"`
	Email     string `json:"email"`
	Nickname  string `json:"nickname"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

type TokenManager struct {
	accessSecret         []byte
	refreshSecret        []byte
	accessExpireSeconds  int64
	refreshExpireSeconds int64
}

func NewTokenManager(accessSecret, refreshSecret string, accessExpireSeconds, refreshExpireSeconds int64) *TokenManager {
	return &TokenManager{
		accessSecret:         []byte(accessSecret),
		refreshSecret:        []byte(refreshSecret),
		accessExpireSeconds:  accessExpireSeconds,
		refreshExpireSeconds: refreshExpireSeconds,
	}
}

func (m *TokenManager) GenerateTokenPair(userID int64, email, nickname string) (string, string, int64, error) {
	accessToken, err := m.generateToken(
		userID,
		email,
		nickname,
		constants.TokenTypeAccess,
		m.accessExpireSeconds,
		m.accessSecret,
	)
	if err != nil {
		return "", "", 0, err
	}

	refreshToken, err := m.generateToken(
		userID,
		email,
		nickname,
		constants.TokenTypeRefresh,
		m.refreshExpireSeconds,
		m.refreshSecret,
	)
	if err != nil {
		return "", "", 0, err
	}

	return accessToken, refreshToken, m.accessExpireSeconds, nil
}

func (m *TokenManager) ParseAccessToken(tokenString string) (*UserClaims, error) {
	return m.parseToken(tokenString, m.accessSecret, constants.TokenTypeAccess)
}

func (m *TokenManager) ParseRefreshToken(tokenString string) (*UserClaims, error) {
	return m.parseToken(tokenString, m.refreshSecret, constants.TokenTypeRefresh)
}

func (m *TokenManager) parseToken(tokenString string, secret []byte, expectedType string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}

		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims.TokenType != expectedType {
		return nil, errors.New("invalid token type")
	}

	return claims, nil
}

func (m *TokenManager) generateToken(
	userID int64,
	email, nickname, tokenType string,
	expireSeconds int64,
	secret []byte,
) (string, error) {
	now := time.Now()
	claims := UserClaims{
		UserID:    userID,
		Email:     email,
		Nickname:  nickname,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(expireSeconds) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
