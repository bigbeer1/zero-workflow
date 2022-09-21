package logic

import (
	"context"
	"zero-workflow/workflow/rpc/internal/svc"
	"zero-workflow/workflow/rpc/workflow"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcdefFindListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcdefFindListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcdefFindListLogic {
	return &ProcdefFindListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcdefFindListLogic) ProcdefFindList(in *workflow.ProcdefFindListReq) (*workflow.ProcdefFindListResp, error) {

	data, err := l.svcCtx.ProcdefModel.FindList(in.Current, in.PageSize, in.Name, in.ProcType, in.TenantId)
	if err != nil {
		return nil, err
	}
	var list []*workflow.ProcdefListData
	for _, item := range *data {
		list = append(list, &workflow.ProcdefListData{
			Id:          item.Id,
			CreatedAt:   item.CreatedAt.UnixMilli(),
			CreatedName: item.CreatedName,
			Name:        item.Name,
			ProcType:    item.ProcType,
		})
	}

	total := l.svcCtx.ProcdefModel.Count(in.Name, in.ProcType, in.TenantId)
	if err != nil {
		return nil, err
	}
	return &workflow.ProcdefFindListResp{
		Total: total,
		List:  list,
	}, nil
}
