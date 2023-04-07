package handler

import (
	"net/http"

	"evm-container/api/internal/logic"
	"evm-container/api/internal/svc"
	"evm-container/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DelegateCallHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelegateCallRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDelegateCallLogic(r.Context(), svcCtx)
		resp, err := l.DelegateCall(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
