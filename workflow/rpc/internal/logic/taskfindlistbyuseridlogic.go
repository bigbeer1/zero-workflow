package logic

import (
	"context"
	"zero-workflow/workflow/rpc/internal/svc"
	"zero-workflow/workflow/rpc/workflow"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskFindListByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskFindListByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskFindListByUserIdLogic {
	return &TaskFindListByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskFindListByUserIdLogic) TaskFindListByUserId(in *workflow.TaskFindListByUserIdReq) (*workflow.TaskFindListByUserIdResp, error) {

	data, err := l.svcCtx.TaskModel.FindListByUserId(in.Current, in.PageSize, in.UserId)
	if err != nil {
		return nil, err
	}

	var list []*workflow.TaskListData
	for _, item := range *data {
		list = append(list, &workflow.TaskListData{
			TaskId:        item.TaskId,
			ProcinstId:    item.Procinstid,
			CreatedAt:     item.CreatedAt,
			NodeId:        item.NodeId,
			Step:          int64(item.Step),
			AgreeNum:      item.AgreeNum,
			ProcType:      item.ProcType,
			ProcinstName:  item.ProcinstName,
			ProcinstTitle: item.ProcinstTitle,
			StartTime:     item.StartTime,
			StartUserName: item.StartUserName,
			StartUserId:   item.StartUserId,
		})
	}

	total := l.svcCtx.TaskModel.CountByUserId(in.UserId)

	return &workflow.TaskFindListByUserIdResp{
		Total: total,
		List:  list,
	}, nil
}
