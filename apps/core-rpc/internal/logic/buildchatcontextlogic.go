package logic

import (
	"context"

	"github.com/Yooz-1999/agentforge/apps/core-rpc/internal/svc"
	"github.com/Yooz-1999/agentforge/apps/core-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BuildChatContextLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBuildChatContextLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BuildChatContextLogic {
	return &BuildChatContextLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BuildChatContextLogic) BuildChatContext(in *pb.BuildChatContextRequest) (*pb.BuildChatContextResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.BuildChatContextResponse{}, nil
}
