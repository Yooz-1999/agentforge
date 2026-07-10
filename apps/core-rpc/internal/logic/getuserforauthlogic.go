package logic

import (
	"context"
	"time"

	"github.com/Yooz-1999/agentforge/apps/core-rpc/internal/svc"
	"github.com/Yooz-1999/agentforge/apps/core-rpc/pb"
	"github.com/Yooz-1999/agentforge/pkg/constants"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetUserForAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserForAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserForAuthLogic {
	return &GetUserForAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserForAuthLogic) GetUserForAuth(in *pb.GetUserForAuthRequest) (*pb.GetUserForAuthResponse, error) {
	if in.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

	user, err := l.svcCtx.UserRepo.FindByID(l.ctx, in.Id)
	if err != nil {
		l.Errorf("find user by id failed: %v", err)
		return nil, status.Error(codes.Internal, "failed to query user")
	}
	if user == nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	if user.Status != constants.UserStatusActive {
		return nil, status.Error(codes.PermissionDenied, "user is disabled")
	}

	return &pb.GetUserForAuthResponse{
		User: &pb.User{
			Id:        user.ID,
			Email:     user.Email,
			Nickname:  user.Nickname,
			Status:    user.Status,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		},
	}, nil
}
