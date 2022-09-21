package Take

import (
	"context"
	"zero-workflow/common"
	"zero-workflow/common/jwtx"
	"zero-workflow/common/msg"
	"zero-workflow/workflow/api/internal/svc"
	"zero-workflow/workflow/api/internal/types"
	"zero-workflow/workflow/rpc/workflowclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskListByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskListByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) TaskListByUserIdLogic {
	return TaskListByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskListByUserIdLogic) TaskListByUserId(req types.TaskListByUserIdReq) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	data, err := l.svcCtx.WorkflowRpc.TaskFindListByUserId(l.ctx, &workflowclient.TaskFindListByUserIdReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		UserId:   tokenData.Uid,
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: data,
	}, nil

}
