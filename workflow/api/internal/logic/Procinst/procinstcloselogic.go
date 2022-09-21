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

type ProcinstCloseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcinstCloseLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProcinstCloseLogic {
	return ProcinstCloseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcinstCloseLogic) ProcinstClose(req types.ProcinstCloseRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.WorkflowRpc.ProcinstClose(l.ctx, &workflowclient.ProcinstCloseReq{
		Id:       req.Id,
		UserId:   tokenData.Uid,
		UserName: tokenData.NickName,
		TenantId: tokenData.TenantId,
	})

	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}
	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: nil,
	}, nil
}
