package Procdef

import (
	"context"
	"encoding/json"
	"zero-workflow/common"
	"zero-workflow/common/jsonx"
	"zero-workflow/common/jwtx"
	"zero-workflow/common/msg"
	"zero-workflow/workflow/api/internal/svc"
	"zero-workflow/workflow/api/internal/types"
	"zero-workflow/workflow/model"
	"zero-workflow/workflow/rpc/workflowclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcdefAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcdefAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProcdefAddLogic {
	return ProcdefAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcdefAddLogic) ProcdefAdd(req types.ProcdefAddRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	// 转换Resource 为node来判断流程是否正常
	data, err := json.Marshal(req.Resource)
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}
	var resourceData model.Node
	err = json.Unmarshal(data, &resourceData)
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	//判断流程是否正常
	err = model.IfProcessConifgIsValid(&resourceData)
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	str, err := jsonx.ToJSONStr(req.Resource)
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	dataStr, err := jsonx.ToJSONStr(req.Data)
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	//添加流程
	_, err = l.svcCtx.WorkflowRpc.ProcdefAdd(l.ctx, &workflowclient.ProcdefAddReq{
		CreatedName: tokenData.NickName,
		Name:        req.Name,
		Data:        dataStr,
		ProcType:    req.ProcType,
		Resource:    str,
		TenantId:    tokenData.TenantId,
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
