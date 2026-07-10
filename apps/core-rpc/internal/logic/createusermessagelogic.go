package logic

import (
	"context"

	"github.com/Yooz-1999/agentforge/apps/core-rpc/internal/svc"
	"github.com/Yooz-1999/agentforge/apps/core-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateUserMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserMessageLogic {
	return &CreateUserMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserMessageLogic) CreateUserMessage(in *pb.CreateUserMessageRequest) (*pb.MessageResponse, error) {
	return nil, status.Error(codes.Unimplemented, "接口还没有实现")
}
