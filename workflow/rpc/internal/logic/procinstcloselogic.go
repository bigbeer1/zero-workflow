package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"zero-workflow/workflow/model"
	"zero-workflow/workflow/rpc/internal/svc"
	"zero-workflow/workflow/rpc/workflow"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcinstCloseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcinstCloseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcinstCloseLogic {
	return &ProcinstCloseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcinstCloseLogic) ProcinstClose(in *workflow.ProcinstCloseReq) (*workflow.CommonResp, error) {
	completeLock.Lock()         // 关锁
	defer completeLock.Unlock() //解锁
	// 查询任务
	ProcinstRes, err := l.svcCtx.ProcinstModel.FindOne(in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New("流程不存在")
		}
		return nil, err
	}
	if ProcinstRes.TenantId != in.TenantId {
		return nil, errors.New("不是一个租户非法操作 ")
	}
	if ProcinstRes.IsFinished == 1 {
		return nil, errors.New("流程已经结束")
	}

	takeRes, err := l.svcCtx.TaskModel.FindOne(ProcinstRes.TaskId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New(fmt.Sprintf("%s,没有这个任务", ProcinstRes.TaskId))
		}
		return nil, err
	}

	//开启事务
	err = l.svcCtx.ProcinstModel.TransCtx(l.ctx, func(ctx context.Context, sqlx sqlx.Session) error {
		// 更新上一个take为结束
		takeRes.IsFinished = 1
		err := l.svcCtx.TaskModel.TransUpdate(ctx, sqlx, takeRes)
		if err != nil {
			return err
		}

		// 添加终止任务
		task := &model.Task{
			Id:        uuid.NewV4().String(),
			CreatedAt: time.Now().UnixMilli(),
			ClaimTime: sql.NullInt64{
				Int64: time.Now().UnixMilli(),
				Valid: true,
			},
			NodeId:        "终止流程",
			NodeName:      "终止流程",
			NodeType:      "终止流程",
			Step:          takeRes.Step + 1,
			ProcinstId:    takeRes.ProcinstId,
			AssigneeId:    in.UserId,
			AssigneeName:  in.TenantId,
			UnCompleteNum: 0,
			AgreeNum:      1,
			IsFinished:    1,
			TenantId:      in.TenantId,
		}
		_, err = l.svcCtx.TaskModel.TransInsert(ctx, sqlx, task)
		if err != nil {
			return err
		}

		// 更新流程实例
		ProcinstRes.EndTime = sql.NullInt64{
			Int64: time.Now().UnixMilli(),
			Valid: true,
		}
		ProcinstRes.IsFinished = 1
		err = l.svcCtx.ProcinstModel.TransUpdate(ctx, sqlx, ProcinstRes)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &workflow.CommonResp{}, nil
}
