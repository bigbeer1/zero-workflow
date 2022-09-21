package logic

import (
	"context"
	"errors"
	"fmt"
	"zero-workflow/workflow/model"
	"zero-workflow/workflow/rpc/internal/svc"
	"zero-workflow/workflow/rpc/workflow"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExecutionFindOneByProcinstIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExecutionFindOneByProcinstIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExecutionFindOneByProcinstIdLogic {
	return &ExecutionFindOneByProcinstIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  -----------------------实例内容-----------------------------
func (l *ExecutionFindOneByProcinstIdLogic) ExecutionFindOneByProcinstId(in *workflow.ExecutionFindOneByProcinstIdReq) (*workflow.ExecutionFindOneByProcinstIdResp, error) {

	res, err := l.svcCtx.ExecutionModel.FindOneByProcinstIdAndTenantId(in.ProcinstId, in.TenantId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New(fmt.Sprintf("没有找到Execution流程实例内容Id:%s", in.ProcinstId))
		}
		return nil, err
	}
	if res.TenantId != in.TenantId {
		return nil, errors.New("不是一个租户非法操作 ")
	}

	return &workflow.ExecutionFindOneByProcinstIdResp{
		Id:          res.Id,
		ProcdefName: res.ProcdefName,
		NodeInfos:   res.NodeInfos,
		StartTime:   res.StartTime,
	}, nil
}
