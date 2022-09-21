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

type ProcdefInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcdefInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProcdefInfoLogic {
	return ProcdefInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcdefInfoLogic) ProcdefInfo(req types.ProcdefInfoRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	data, err := l.svcCtx.WorkflowRpc.ProcdefFindOne(l.ctx, &workflowclient.ProcdefFindOneReq{
		Id:       req.Id,
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
