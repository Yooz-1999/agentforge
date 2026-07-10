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

type GetAgentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAgentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAgentLogic {
	return &GetAgentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAgentLogic) GetAgent(req *types.GetAgentRequest) (resp *types.AgentResponse, err error) {
	return nil, apperrors.ErrUnimplemented
}
