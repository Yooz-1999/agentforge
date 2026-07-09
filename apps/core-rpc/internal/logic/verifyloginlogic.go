package logic

import (
	"context"

	"github.com/Yooz-1999/agentforge/apps/core-rpc/internal/svc"
	"github.com/Yooz-1999/agentforge/apps/core-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return &pb.VerifyLoginResponse{}, nil
}
