package logic

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/Yooz-1999/agentforge/apps/core-rpc/internal/svc"
	"github.com/Yooz-1999/agentforge/apps/core-rpc/pb"
	"github.com/Yooz-1999/agentforge/pkg/auth"
	"github.com/Yooz-1999/agentforge/pkg/constants"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VerifyLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLoginLogic {
	return &VerifyLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyLoginLogic) VerifyLogin(in *pb.VerifyLoginRequest) (*pb.VerifyLoginResponse, error) {
	email := strings.TrimSpace(strings.ToLower(in.Email))
	password := in.Password
	if email == "" || password == "" {
		return nil, status.Error(codes.InvalidArgument, "email and password are required")
	}

	user, err := l.svcCtx.UserRepo.FindByEmail(l.ctx, email)
	if err != nil {
		l.Errorf("find user by email failed: %v", err)
		return nil, status.Error(codes.Internal, "failed to query user")
	}
	if user == nil {
		return nil, status.Error(codes.Unauthenticated, "invalid email or password")
	}
	if user.Status != constants.UserStatusActive {
		return nil, status.Error(codes.Unauthenticated, "invalid email or password")
	}
	if err := auth.ComparePassword(user.PasswordHash, password); err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid email or password")
	}

	now := time.Now()
	updatedAt := user.UpdatedAt
	if err := l.svcCtx.UserRepo.UpdateLastLoginAt(l.ctx, user.ID, sql.NullTime{Time: now, Valid: true}); err != nil {
		l.Errorf("update last login time failed: %v", err)
	} else {
		updatedAt = now
	}

	return &pb.VerifyLoginResponse{
		User: &pb.User{
			Id:        user.ID,
			Email:     user.Email,
			Nickname:  user.Nickname,
			Status:    user.Status,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: updatedAt.Format(time.RFC3339),
		},
	}, nil
}
