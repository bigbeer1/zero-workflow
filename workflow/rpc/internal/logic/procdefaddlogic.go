package logic

import (
	"context"
	"database/sql"
	uuid "github.com/satori/go.uuid"
	"time"
	"zero-workflow/workflow/model"
	"zero-workflow/workflow/rpc/internal/svc"
	"zero-workflow/workflow/rpc/workflow"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcdefAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcdefAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcdefAddLogic {
	return &ProcdefAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -------------------------添加流程定义----------------------
func (l *ProcdefAddLogic) ProcdefAdd(in *workflow.ProcdefAddReq) (*workflow.CommonResp, error) {

	_, err := l.svcCtx.ProcdefModel.Insert(&model.Procdef{
		Id:          uuid.NewV4().String(),
		CreatedAt:   time.Now(),
		CreatedName: in.CreatedName,
		Name:        in.Name,
		Data: sql.NullString{
			String: in.Data,
			Valid:  in.Data != "",
		},
		ProcType: in.ProcType,
		Resource: in.Resource,
		TenantId: in.TenantId,
	})
	if err != nil {
		return nil, err
	}

	return &workflow.CommonResp{}, nil
}
