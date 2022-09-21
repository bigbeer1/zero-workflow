package Procinst

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

type ProcinstInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcinstInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProcinstInfoLogic {
	return ProcinstInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcinstInfoLogic) ProcinstInfo(req types.ProcinstInfoRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	// 查询实例
	ProcinstRes, err := l.svcCtx.WorkflowRpc.ProcinstFindOne(l.ctx, &workflowclient.ProcinstFindOneReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	// 获取所有审批内容
	taskList, err := l.svcCtx.WorkflowRpc.TaskFindListByProcinstId(l.ctx, &workflowclient.TaskFindListByProcinstIdReq{
		ProcinstId: ProcinstRes.Id,
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}
	ExecutionRes, err := l.svcCtx.WorkflowRpc.ExecutionFindOneByProcinstId(l.ctx, &workflowclient.ExecutionFindOneByProcinstIdReq{
		ProcinstId: ProcinstRes.Id,
		TenantId:   tokenData.TenantId,
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	type resultData struct {
		Procinst  *workflowclient.ProcinstFindOneResp              `json:"procinst"`
		TaskList  *workflowclient.TaskFindListByProcinstIdResp     `json:"taskList"`
		Execution *workflowclient.ExecutionFindOneByProcinstIdResp `json:"execution"`
	}

	// 获取实例流程
	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: resultData{
			Procinst:  ProcinstRes,
			TaskList:  taskList,
			Execution: ExecutionRes,
		},
	}, nil
}
