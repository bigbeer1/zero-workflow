package Procdef

import (
	"net/http"
	"zero-workflow/common"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-workflow/workflow/api/internal/logic/Procdef"
	"zero-workflow/workflow/api/internal/svc"

	"zero-workflow/workflow/api/internal/types"
)

func ProcdefInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProcdefInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, common.NewDefaultError(err.Error()))
			return
		}

		l := Procdef.NewProcdefInfoLogic(r.Context(), svcCtx)
		resp, err := l.ProcdefInfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
