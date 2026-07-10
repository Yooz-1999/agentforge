package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Yooz-1999/agentforge/apps/core-rpc/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

// NewUserRepository 创建用户仓储对象。
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByEmail 根据邮箱查找用户。
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	const query = `
SELECT id, email, password_hash, nickname, status, last_login_at, created_at, updated_at, deleted_at
FROM users
WHERE email = ? AND deleted_at IS NULL
LIMIT 1`

	return r.findOne(ctx, query, email)
}

// FindByID 根据用户 ID 查找未删除的用户。
func (r *UserRepository) FindByID(ctx context.Context, userID int64) (*model.User, error) {
	const query = `
SELECT id, email, password_hash, nickname, status, last_login_at, created_at, updated_at, deleted_at
FROM users
WHERE id = ? AND deleted_at IS NULL
LIMIT 1`

	return r.findOne(ctx, query, userID)
}

func (r *UserRepository) findOne(ctx context.Context, query string, arg any) (*model.User, error) {
	var user model.User
	var lastLoginAt sql.NullTime
	var deletedAt sql.NullTime

	err := r.db.QueryRowContext(ctx, query, arg).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Nickname,
		&user.Status,
		&lastLoginAt,
		&user.CreatedAt,
		&user.UpdatedAt,
		&deletedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	if lastLoginAt.Valid {
		user.LastLoginAt = &lastLoginAt.Time
	}
	if deletedAt.Valid {
		user.DeletedAt = &deletedAt.Time
	}

	return &user, nil
}

// Create 创建用户记录，并回填新生成的用户 ID。
func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	const query = `
INSERT INTO users (email, password_hash, nickname, status, last_login_at, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(
		ctx,
		query,
		user.Email,
		user.PasswordHash,
		user.Nickname,
		user.Status,
		user.LastLoginAt,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id
	return nil
}

// UpdateLastLoginAt 更新用户最后登录时间。
func (r *UserRepository) UpdateLastLoginAt(ctx context.Context, userID int64, atTime sql.NullTime) error {
	const query = `
UPDATE users
SET last_login_at = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL`

	_, err := r.db.ExecContext(ctx, query, atTime, atTime.Time, userID)
	return err
}
