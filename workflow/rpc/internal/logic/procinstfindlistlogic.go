package logic

import (
	"context"
	"zero-workflow/workflow/rpc/internal/svc"
	"zero-workflow/workflow/rpc/workflow"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcinstFindListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcinstFindListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcinstFindListLogic {
	return &ProcinstFindListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcinstFindListLogic) ProcinstFindList(in *workflow.ProcinstFindListReq) (*workflow.ProcinstFindListResp, error) {

	all, err := l.svcCtx.ProcinstModel.FindList(in.Current, in.PageSize, in.ProcType, in.ProcdefName, in.Title, in.StartTime, in.EndTime, in.StartUserName, in.IsFinished, in.TenantId)
	if err != nil {
		return nil, err
	}

	var list []*workflow.ProcinstListData
	for _, item := range *all {
		list = append(list, &workflow.ProcinstListData{
			Id:            item.Id,
			ProcType:      item.ProcType,
			ProcdefName:   item.ProcdefName,
			Title:         item.Title,
			StartTime:     item.StartTime,
			EndTime:       item.EndTime.Int64,
			StartUserId:   item.StartUserId,
			StartUserName: item.StartUserName,
			IsFinished:    item.IsFinished,
			TaskId:        item.TaskId,
			NodeId:        item.NodeId,
			NodeName:      item.NodeName,
			NodeType:      item.NodeType,
			Step:          item.Step,
			AssigneeId:    item.AssigneeId,
			AssigneeName:  item.AssigneeName,
		})
	}
	total := l.svcCtx.ProcinstModel.Count(in.ProcType, in.ProcdefName, in.Title, in.StartTime, in.EndTime, in.StartUserName, in.IsFinished, in.TenantId)

	return &workflow.ProcinstFindListResp{
		Total: total,
		List:  list,
	}, nil
}
