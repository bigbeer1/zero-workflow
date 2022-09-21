package Take

import (
	"net/http"
	"zero-workflow/common"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-workflow/workflow/api/internal/logic/Take"
	"zero-workflow/workflow/api/internal/svc"

	"zero-workflow/workflow/api/internal/types"
)

func TaskCompleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskCompleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, common.NewDefaultError(err.Error()))
			return
		}

		l := Take.NewTaskCompleteLogic(r.Context(), svcCtx)
		resp, err := l.TaskComplete(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
