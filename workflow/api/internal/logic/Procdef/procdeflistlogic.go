package Procdef

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

type ProcdefListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcdefListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProcdefListLogic {
	return ProcdefListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcdefListLogic) ProcdefList(req types.ProcdefListRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	data, err := l.svcCtx.WorkflowRpc.ProcdefFindList(l.ctx, &workflowclient.ProcdefFindListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
		ProcType: req.ProcType,
		TenantId: tokenData.TenantId,
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
