package constants

const (
	UserStatusActive   int64 = 1
	UserStatusDisabled int64 = 2
)

const (
	EmailMaxCharacters    = 128
	NicknameMaxCharacters = 64
	PasswordMinCharacters = 6
	PasswordMaxBytes      = 72
)

const (
	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
)
