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

type ProcinstFindListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcinstFindListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProcinstFindListLogic {
	return ProcinstFindListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcinstFindListLogic) ProcinstFindList(req types.ProcinstFindListRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	data, err := l.svcCtx.WorkflowRpc.ProcinstFindList(l.ctx, &workflowclient.ProcinstFindListReq{
		Current:       req.Current,
		PageSize:      req.PageSize,
		ProcType:      req.ProcType,
		ProcdefName:   req.ProcdefName,
		Title:         req.Title,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		StartUserName: req.StartUserName,
		IsFinished:    req.IsFinished,
		TenantId:      tokenData.TenantId,
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
