package Take

import (
	"net/http"
	"zero-workflow/common"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-workflow/workflow/api/internal/logic/Take"
	"zero-workflow/workflow/api/internal/svc"

	"zero-workflow/workflow/api/internal/types"
)

func TaskListByUserIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskListByUserIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, common.NewDefaultError(err.Error()))
			return
		}

		l := Take.NewTaskListByUserIdLogic(r.Context(), svcCtx)
		resp, err := l.TaskListByUserId(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
