package logic

import (
	"context"

	"github.com/Yooz-1999/agentforge/apps/core-rpc/internal/svc"
	"github.com/Yooz-1999/agentforge/apps/core-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteAgentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAgentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAgentLogic {
	return &DeleteAgentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAgentLogic) DeleteAgent(in *pb.DeleteAgentRequest) (*pb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "接口还没有实现")
}
