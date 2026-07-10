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

type UpdateAgentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAgentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAgentLogic {
	return &UpdateAgentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAgentLogic) UpdateAgent(req *types.UpdateAgentRequest) (resp *types.AgentResponse, err error) {
	return nil, apperrors.ErrUnimplemented
}
