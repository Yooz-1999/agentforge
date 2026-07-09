// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/Yooz-1999/agentforge/apps/chat-api/internal/svc"
	"github.com/Yooz-1999/agentforge/apps/chat-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatStreamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatStreamLogic {
	return &ChatStreamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatStreamLogic) ChatStream(req *types.ChatStreamRequest) (resp *types.ChatStreamBootstrapResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
