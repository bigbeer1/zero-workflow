package Take

import (
	"context"
	"zero-workflow/common"
	"zero-workflow/common/jwtx"
	"zero-workflow/workflow/api/internal/svc"
	"zero-workflow/workflow/api/internal/types"
	"zero-workflow/workflow/rpc/workflowclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskCompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) TaskCompleteLogic {
	return TaskCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskCompleteLogic) TaskComplete(req types.TaskCompleteReq) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.WorkflowRpc.TaskComplete(l.ctx, &workflowclient.TaskCompleteReq{
		TaskId:   req.TaskId,
		UserId:   tokenData.Uid,
		UserName: tokenData.NickName,
		Pass:     req.Pass,
		Comment:  req.Comment,
		TenantId: tokenData.TenantId,
	})

	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	return
}
