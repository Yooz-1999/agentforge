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

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenRequest) (resp *types.RefreshTokenResponse, err error) {
	claims, err := l.svcCtx.TokenManager.ParseRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, invalidRefreshTokenError()
	}
	if claims.UserID <= 0 {
		return nil, invalidRefreshTokenError()
	}

	key := fmt.Sprintf("auth:refresh:%d", claims.UserID)
	storedToken, err := l.svcCtx.Redis.Get(l.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, invalidRefreshTokenError()
		}

		l.Errorf("get refresh token failed: %v", err)
		return nil, status.Error(codes.Unavailable, "refresh service is temporarily unavailable")
	}
	if storedToken != req.RefreshToken {
		return nil, invalidRefreshTokenError()
	}

	userResult, err := l.svcCtx.Core.GetUserForAuth(l.ctx, &pb.GetUserForAuthRequest{Id: claims.UserID})
	if err != nil {
		switch status.Code(err) {
		case codes.InvalidArgument, codes.NotFound, codes.PermissionDenied:
			if deleteErr := l.svcCtx.Redis.Del(l.ctx, key).Err(); deleteErr != nil {
				l.Errorf("delete invalid refresh token failed: %v", deleteErr)
			}
			return nil, invalidRefreshTokenError()
		default:
			l.Errorf("get user for refresh failed: %v", err)
			return nil, status.Error(codes.Unavailable, "refresh service is temporarily unavailable")
		}
	}
	if userResult.User == nil {
		return nil, invalidRefreshTokenError()
	}

	accessToken, refreshToken, expiresIn, err := l.svcCtx.TokenManager.GenerateTokenPair(
		userResult.User.Id,
		userResult.User.Email,
		userResult.User.Nickname,
	)
	if err != nil {
		l.Errorf("generate token pair failed: %v", err)
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	if err := l.svcCtx.Redis.Set(
		l.ctx,
		key,
		refreshToken,
		time.Duration(l.svcCtx.Config.JWT.RefreshExpireSeconds)*time.Second,
	).Err(); err != nil {
		l.Errorf("store refresh token failed: %v", err)
		return nil, status.Error(codes.Unavailable, "refresh service is temporarily unavailable")
	}

	return &types.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}, nil
}

func invalidRefreshTokenError() error {
	return status.Error(codes.Unauthenticated, "invalid refresh token")
}
