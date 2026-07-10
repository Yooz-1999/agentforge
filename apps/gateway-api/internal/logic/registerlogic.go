// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/Yooz-1999/agentforge/apps/core-rpc/pb"
	"github.com/Yooz-1999/agentforge/apps/gateway-api/internal/svc"
	"github.com/Yooz-1999/agentforge/apps/gateway-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	result, err := l.svcCtx.Core.RegisterUser(l.ctx, &pb.RegisterUserRequest{
		Email:    req.Email,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResponse{
		Id:       result.User.Id,
		Email:    result.User.Email,
		Nickname: result.User.Nickname,
	}, nil
}
