// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/Yooz-1999/agentforge/apps/core-rpc/pb"
	"github.com/Yooz-1999/agentforge/apps/gateway-api/internal/svc"
	"github.com/Yooz-1999/agentforge/apps/gateway-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	result, err := l.svcCtx.Core.VerifyLogin(l.ctx, &pb.VerifyLoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, expiresIn, err := l.svcCtx.TokenManager.GenerateTokenPair(
		result.User.Id,
		result.User.Email,
		result.User.Nickname,
	)
	if err != nil {
		l.Errorf("generate token pair failed: %v", err)
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	key := fmt.Sprintf("auth:refresh:%d", result.User.Id)
	if err := l.svcCtx.Redis.Set(
		l.ctx,
		key,
		refreshToken,
		time.Duration(l.svcCtx.Config.JWT.RefreshExpireSeconds)*time.Second,
	).Err(); err != nil {
		l.Errorf("store refresh token failed: %v", err)
		return nil, status.Error(codes.Unavailable, "login service is temporarily unavailable")
	}

	return &types.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
		User: types.UserInfo{
			Id:       result.User.Id,
			Email:    result.User.Email,
			Nickname: result.User.Nickname,
		},
	}, nil
}
