package logic

import (
	"context"

	"github.com/Yooz-1999/agentforge/apps/core-rpc/internal/svc"
	"github.com/Yooz-1999/agentforge/apps/core-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return &pb.MessageResponse{}, nil
}
