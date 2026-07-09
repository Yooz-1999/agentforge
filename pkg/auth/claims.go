package auth

type UserClaims struct {
	UserID   int64
	Email    string
	Nickname string
}
