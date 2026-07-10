package model

import "time"

type User struct {
	ID           int64
	Email        string
	PasswordHash string
	Nickname     string
	Status       int64
	LastLoginAt  *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
