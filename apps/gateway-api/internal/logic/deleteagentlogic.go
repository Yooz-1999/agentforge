// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/Yooz-1999/agentforge/apps/gateway-api/internal/svc"
	"github.com/Yooz-1999/agentforge/apps/gateway-api/internal/types"
	apperrors "github.com/Yooz-1999/agentforge/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAgentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAgentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAgentLogic {
	return &DeleteAgentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAgentLogic) DeleteAgent(req *types.DeleteAgentRequest) (resp *types.EmptyResponse, err error) {
	return nil, apperrors.ErrUnimplemented
}
