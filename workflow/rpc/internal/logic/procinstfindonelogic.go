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

type ProcinstFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcinstFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcinstFindOneLogic {
	return &ProcinstFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  -----------------------流程实例-----------------------
func (l *ProcinstFindOneLogic) ProcinstFindOne(in *workflow.ProcinstFindOneReq) (*workflow.ProcinstFindOneResp, error) {
	res, err := l.svcCtx.ProcinstModel.FindOneData(in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New(fmt.Sprintf("%s,流程不存在", in.Id))
		}
		return nil, err
	}

	return &workflow.ProcinstFindOneResp{
		Id:            res.Id,
		ProcType:      res.ProcType,
		ProcdefName:   res.ProcdefName,
		Title:         res.Title,
		StartTime:     res.StartTime,
		EndTime:       res.EndTime.Int64,
		StartUserId:   res.StartUserId,
		StartUserName: res.StartUserName,
		IsFinished:    res.IsFinished,
		TaskId:        res.TaskId,
		NodeId:        res.NodeId,
		NodeName:      res.NodeName,
		NodeType:      res.NodeType,
		Step:          res.Step,
		AssigneeId:    res.AssigneeId,
		AssigneeName:  res.AssigneeName,
	}, nil
}
