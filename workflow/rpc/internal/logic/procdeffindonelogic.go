package logic

import (
	"context"
	"errors"
	"zero-workflow/workflow/model"
	"zero-workflow/workflow/rpc/internal/svc"
	"zero-workflow/workflow/rpc/workflow"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcdefFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcdefFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcdefFindOneLogic {
	return &ProcdefFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcdefFindOneLogic) ProcdefFindOne(in *workflow.ProcdefFindOneReq) (*workflow.ProcdefFindOneResp, error) {
	res, err := l.svcCtx.ProcdefModel.FindOne(in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New("流程不存在")
		}
		return nil, err
	}
	if res.TenantId != in.TenantId {
		return nil, errors.New("不是一个租户非法操作 ")
	}
	return &workflow.ProcdefFindOneResp{
		Id:          res.Id,
		CreatedAt:   res.CreatedAt.UnixMilli(),
		CreatedName: res.CreatedName,
		Data:        res.Data.String,
		Name:        res.Name,
		ProcType:    res.ProcType,
		Resource:    res.Resource,
	}, nil
}
