package logic

import (
	"context"

	"github.com/Yooz-1999/agentforge/apps/core-rpc/internal/svc"
	"github.com/Yooz-1999/agentforge/apps/core-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ListConversationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListConversationsLogic {
	return &ListConversationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListConversationsLogic) ListConversations(in *pb.ListConversationsRequest) (*pb.ListConversationsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "接口还没有实现")
}
