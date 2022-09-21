package logic

import (
	"context"
	"zero-workflow/workflow/rpc/internal/svc"
	"zero-workflow/workflow/rpc/workflow"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskFindListByProcinstIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskFindListByProcinstIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskFindListByProcinstIdLogic {
	return &TaskFindListByProcinstIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskFindListByProcinstIdLogic) TaskFindListByProcinstId(in *workflow.TaskFindListByProcinstIdReq) (*workflow.TaskFindListByProcinstIdResp, error) {

	all, err := l.svcCtx.TaskModel.FindListByProcinstId(in.ProcinstId)
	if err != nil {
		return nil, err
	}

	var list []*workflow.TaskProcinstData
	for _, item := range *all {
		list = append(list, &workflow.TaskProcinstData{
			TaskId:       item.TaskId,
			CreatedAt:    item.CreatedAt,
			ClaimTime:    item.ClaimTime.Int64,
			NodeId:       item.NodeId,
			NodeName:     item.NodeName,
			NodeType:     item.NodeType,
			Step:         item.Step,
			AssigneeName: item.AssigneeName,
			IsFinished:   item.IsFinished,
			IsAgree:      item.IsAgree.Int64,
			Comment:      item.Comment.String,
		})

	}

	return &workflow.TaskFindListByProcinstIdResp{
		List: list,
	}, nil
}
