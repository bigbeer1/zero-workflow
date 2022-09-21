package logic

import (
	"context"
	"errors"
	"zero-workflow/workflow/model"
	"zero-workflow/workflow/rpc/internal/svc"
	"zero-workflow/workflow/rpc/workflow"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcdefDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcdefDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcdefDeleteLogic {
	return &ProcdefDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcdefDeleteLogic) ProcdefDelete(in *workflow.ProcdefDeleteReq) (*workflow.CommonResp, error) {

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
	err = l.svcCtx.ProcdefModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &workflow.CommonResp{}, nil
}
