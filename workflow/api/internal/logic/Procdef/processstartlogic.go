package Procdef

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"zero-workflow/common"
	"zero-workflow/common/jwtx"
	"zero-workflow/common/msg"
	"zero-workflow/workflow/api/internal/svc"
	"zero-workflow/workflow/api/internal/types"
	"zero-workflow/workflow/rpc/workflowclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcessStartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcessStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProcessStartLogic {
	return ProcessStartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcessStartLogic) ProcessStart(req types.ProcessStartRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.WorkflowRpc.ProcessStart(l.ctx, &workflowclient.ProcessStartReq{
		ProcdefId:  req.ProcdefId,
		ProcinstId: uuid.NewV4().String(),
		Title:      req.Title,
		UserId:     tokenData.Uid,
		NickName:   tokenData.NickName,
		TenantId:   tokenData.TenantId,
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
