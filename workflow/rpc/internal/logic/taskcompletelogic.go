package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"sync"
	"time"
	"zero-workflow/common/jsonx"
	"zero-workflow/workflow/model"
	"zero-workflow/workflow/rpc/internal/svc"
	"zero-workflow/workflow/rpc/workflow"

	"github.com/zeromicro/go-zero/core/logx"
)

var completeLock sync.Mutex

type TaskCompleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskCompleteLogic {
	return &TaskCompleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  -----------------------任务--------------------------
func (l *TaskCompleteLogic) TaskComplete(in *workflow.TaskCompleteReq) (*workflow.CommonResp, error) {
	// 查询任务
	completeLock.Lock()         // 关锁
	defer completeLock.Unlock() //解锁
	taskRes, err := l.svcCtx.TaskModel.FindOne(in.TaskId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New(fmt.Sprintf("%s,没有这个任务", in.TaskId))
		}
		return nil, err
	}
	if taskRes.TenantId != in.TenantId {
		return nil, errors.New("不是一个租户非法操作 ")
	}
	// 判断是否已经结束
	if taskRes.IsFinished == 1 {
		if taskRes.NodeId == "结束" {
			return nil, errors.New("流程已经结束")
		}
		return nil, errors.New("任务【" + fmt.Sprintf("%d", in.TaskId) + "】已经被审批过了！！")
	}

	if taskRes.AssigneeId != in.UserId {
		return nil, errors.New("你不是当前节点审批人，你无权审批")
	}

	taskRes.AssigneeId = in.UserId
	taskRes.ClaimTime.Int64 = time.Now().UnixMilli()
	taskRes.ClaimTime.Valid = true

	//同意
	if in.Pass {
		taskRes.AgreeNum++
	} else {
		taskRes.IsFinished = 1
	}
	// 未审批人数减一
	taskRes.UnCompleteNum--
	// 判断是否结束
	if taskRes.UnCompleteNum <= 0 {
		taskRes.IsFinished = 1
	}

	// 开启事务
	err = l.svcCtx.TaskModel.TransCtx(l.ctx, func(ctx context.Context, sqlx sqlx.Session) error {
		err = l.svcCtx.TaskModel.TransUpdate(ctx, sqlx, taskRes)
		if err != nil {
			return err
		}

		// 流转到下一流程
		// 获取执行流经过的节点信息
		ExecutionRes, err := l.svcCtx.ExecutionModel.FindOneByProcinstIdAndTenantId(taskRes.ProcinstId, in.TenantId)
		if err != nil {
			if err == model.ErrNotFound {
				return errors.New("没有获取到流程执行流")
			}
			return err
		}
		var nodeInfos []*model.NodeInfo
		err = jsonx.Str2Struct(ExecutionRes.NodeInfos, &nodeInfos)
		if err != nil {
			return err
		}
		// 根据takeId 查询流程实例
		procinstRes, err := l.svcCtx.ProcinstModel.FindOne(taskRes.ProcinstId)
		if err != nil {
			if err == model.ErrNotFound {
				return errors.New(fmt.Sprintf("%s,没有该procInst", procinstRes.Id))
			}
			return err
		}
		// 流程运转
		err = MoveStage(ctx, l.svcCtx, sqlx, nodeInfos, procinstRes, in.UserId, in.UserName, in.Comment, taskRes.Id, in.TenantId, taskRes.Step, in.Pass)
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
