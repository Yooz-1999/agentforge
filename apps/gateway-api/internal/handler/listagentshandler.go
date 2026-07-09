// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"

	"github.com/Yooz-1999/agentforge/apps/gateway-api/internal/logic"
	"github.com/Yooz-1999/agentforge/apps/gateway-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListAgentsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewListAgentsLogic(r.Context(), svcCtx)
		resp, err := l.ListAgents()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
