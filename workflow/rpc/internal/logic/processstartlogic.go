package logic

import (
	"container/list"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"zero-workflow/common/jsonx"
	"zero-workflow/workflow/model"
	"zero-workflow/workflow/rpc/internal/svc"
	"zero-workflow/workflow/rpc/workflow"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcessStartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcessStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcessStartLogic {
	return &ProcessStartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  -----------------------流程实例-----------------------
func (l *ProcessStartLogic) ProcessStart(in *workflow.ProcessStartReq) (*workflow.CommonResp, error) {
	var err error
	var prodefRes *model.Procdef
	// 获取流程定义
	// 根据流程id 查询流程
	if in.ProcinstId == "" {
		return nil, errors.New("必须传入ProcinstId,uuid")
	}

	// 如果不是实例方法 则根据转换去创建工作流
	if in.Procdefdata != nil {
		err := json.Unmarshal(in.Procdefdata, &prodefRes)
		if err != nil {
			return nil, err
		}
	} else {
		prodefRes, err = l.svcCtx.ProcdefModel.FindOne(in.ProcdefId)
		if err != nil {
			if err == model.ErrNotFound {
				return nil, errors.New("流程不存在")
			}
			return nil, err
		}
		if prodefRes.TenantId != in.TenantId {
			return nil, errors.New("不是一个租户非法操作 ")
		}
	}

	// 将Resource 转换为node结构体
	node := &model.Node{}
	err = jsonx.Str2Struct(prodefRes.Resource, node)
	if err != nil {
		return nil, err
	}
	// 开启事务
	err = l.svcCtx.ProcinstModel.TransCtx(l.ctx, func(ctx context.Context, sqlx sqlx.Session) error {
		// 新建流程实例
		procinst := &model.Procinst{
			Id:            in.ProcinstId,
			ProcType:      prodefRes.ProcType,
			ProcdefName:   prodefRes.Name,
			Title:         in.Title,
			StartTime:     time.Now().UnixMilli(),
			StartUserId:   in.UserId,
			StartUserName: in.NickName,
			TenantId:      in.TenantId,
		}
		_, err = l.svcCtx.ProcinstModel.TransInsert(ctx, sqlx, procinst)
		if err != nil {
			return err
		}

		// 生成执行流，一串运行节点
		str, err := GenerateExec(node, in.UserId, in.NickName) //事务
		if err != nil {
			return err
		}

		// 创建执行流
		executionID := uuid.NewV4().String()
		_, err = l.svcCtx.ExecutionModel.TansInsert(l.ctx, sqlx, &model.Execution{
			Id:          executionID,
			ProcinstId:  procinst.Id,
			ProcdefName: prodefRes.Name,
			NodeInfos:   str,
			StartTime:   time.Now().UnixMilli(),
			TenantId:    in.TenantId,
		})
		if err != nil {
			return err
		}
		// 获取执行流信息
		var nodeinfos []*model.NodeInfo
		err = jsonx.Str2Struct(str, &nodeinfos)
		if err != nil {
			return err
		}

		// 创建任务
		taskID := uuid.NewV4().String()
		task := &model.Task{
			Id:        taskID,
			CreatedAt: time.Now().UnixMilli(),
			ClaimTime: sql.NullInt64{
				Int64: time.Now().UnixMilli(),
				Valid: true,
			},
			NodeId:        "开始",
			NodeName:      "开始",
			NodeType:      "开始",
			Step:          0,
			ProcinstId:    procinst.Id,
			AssigneeId:    in.UserId,
			AssigneeName:  in.NickName,
			UnCompleteNum: 0,
			AgreeNum:      1,
			IsFinished:    1,
			TenantId:      in.TenantId,
		}

		_, err = l.svcCtx.TaskModel.TransInsert(l.ctx, sqlx, task)
		if err != nil {
			return err
		}

		// 流程移动到下一环节
		err = MoveStage(ctx, l.svcCtx, sqlx, nodeinfos, procinst, in.UserId, in.NickName,
			"启动流程", taskID, in.TenantId, 0, true)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	// 新建执行
	return &workflow.CommonResp{}, nil
}

// 根据流程定义node生成执行流
func GenerateExec(node *model.Node, userId, userName string) (string, error) {
	list, err := ParseProcessConfig(node)
	if err != nil {
		return "", err
	}
	list.PushBack(model.NodeInfo{
		NodeID: "结束",
	})
	list.PushFront(model.NodeInfo{
		NodeID:       "开始",
		AssigneeID:   userId,
		AssigneeName: userName,
	})
	arr := jsonx.List2Array(list)
	str, err := jsonx.ToJSONStr(arr)
	if err != nil {
		return "", err
	}
	return str, err
}

// ParseProcessConfig 解析流程定义json数据
func ParseProcessConfig(node *model.Node) (*list.List, error) {
	// defer fmt.Println("----------解析结束--------")
	list := list.New()
	err := parseProcessConfig(node, list)
	return list, err
}
func parseProcessConfig(node *model.Node, list *list.List) (err error) {
	// fmt.Printf("nodeId=%s\n", node.NodeID)
	node.Add2ExecutionList(list)
	// 存在条件节点
	// 存在子节点
	if node.ChildNode != nil {
		err = parseProcessConfig(node.ChildNode, list)
		if err != nil {
			return err
		}
	}
	return nil
}

// MoveStage MoveStage
// 流程流转
func MoveStage(ctx context.Context, svcCtx *svc.ServiceContext, sqlx sqlx.Session, nodeInfos []*model.NodeInfo,
	procinst *model.Procinst, userId, userName, comment, taskId, tenantId string, step int, pass bool) (err error) {
	// 添加审批结果
	if comment == "启动流程" {
		_, err = svcCtx.IdentitylinkModel.TransInsert(ctx, sqlx, &model.Identitylink{
			Id:         uuid.NewV4().String(),
			UserId:     userId,
			UserName:   userName,
			TaskId:     taskId,
			Step:       step,
			IsAgree:    1, //同意
			ProcinstId: procinst.Id,
			Comment: sql.NullString{
				String: comment,
				Valid:  comment != "",
			},
			TenantId: tenantId,
		})
	} else {
		identitylinkRes, err := svcCtx.IdentitylinkModel.FindOneByUserIdAndTaskIdAndTenantId("", taskId, tenantId)
		if err != nil {
			if err == model.ErrNotFound {
				return errors.New("没有获取到该流程身份")
			}
			return err
		}
		var passInt int64 = 0
		if pass {
			passInt = 1
		}
		identitylinkRes.IsAgree = passInt
		identitylinkRes.Comment.String = comment
		identitylinkRes.Comment.Valid = true
		err = svcCtx.IdentitylinkModel.TransUpdate(ctx, sqlx, identitylinkRes)
		if err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}
	if pass {
		step++
		if step-1 > len(nodeInfos) {
			return errors.New("已经结束无法流转到下一个节点")
		}
	} else {
		step--
		if step < 0 {
			return errors.New("处于开始位置，无法回退到上一个节点")
		}
	}

	// 判断下一流程： 如果是审批人是：抄送人  预留
	// 添加任务
	task := &model.Task{
		Id:            uuid.NewV4().String(),
		CreatedAt:     time.Now().UnixMilli(),
		ClaimTime:     sql.NullInt64{},
		NodeId:        nodeInfos[step].NodeID,
		NodeName:      nodeInfos[step].Name,
		NodeType:      nodeInfos[step].Type,
		Step:          step,
		ProcinstId:    procinst.Id,
		AssigneeId:    nodeInfos[step].AssigneeID,
		AssigneeName:  nodeInfos[step].AssigneeName,
		UnCompleteNum: 1,
		AgreeNum:      0,
		IsFinished:    0,
		TenantId:      tenantId,
	}

	// 修改流程实例
	procinst.NodeId = nodeInfos[step].NodeID
	procinst.TaskId = task.Id
	if pass {
		// 通过
		err := MoveToNextStage(ctx, svcCtx, sqlx, nodeInfos, task, procinst, tenantId, step)
		if err != nil {
			return err
		}
	} else {
		// 驳回
		err = MoveToPrevStage(ctx, svcCtx, sqlx, nodeInfos, task, procinst, tenantId, step)
		if err != nil {
			return err
		}
	}

	return nil

}

// MoveToNextStage MoveToNextStage
// 通过
func MoveToNextStage(ctx context.Context, svcCtx *svc.ServiceContext, sqlx sqlx.Session,
	nodeInfos []*model.NodeInfo, task *model.Task, procInst *model.Procinst, tenantId string, step int) error {
	var currentTime = time.Now().UnixMilli() // 当前时间
	if (step + 1) < len(nodeInfos) {         // 下一步不是【结束】
		// 生成新的任务
		_, err := svcCtx.TaskModel.TransInsert(ctx, sqlx, task)
		if err != nil {
			return err
		}
		// 添加下一个审批人
		_, err = svcCtx.IdentitylinkModel.TransInsert(ctx, sqlx, &model.Identitylink{
			Id:         uuid.NewV4().String(),
			UserId:     nodeInfos[step].AssigneeID,
			UserName:   nodeInfos[step].AssigneeName,
			TaskId:     task.Id,
			Step:       step,
			IsAgree:    2, // 未审批
			ProcinstId: procInst.Id,
			TenantId:   tenantId,
		})
		if err != nil {
			return err
		}
		// 更新流程实例
		err = svcCtx.ProcinstModel.TransUpdate(ctx, sqlx, procInst)
		if err != nil {
			return err
		}
	} else { // 最后一步直接结束
		// 生成新的任务
		task.IsFinished = 1
		task.NodeName = "结束"
		task.NodeType = "结束"
		task.UnCompleteNum = 0
		task.AgreeNum = 1
		task.ClaimTime = sql.NullInt64{
			Int64: currentTime,
			Valid: true,
		}
		// 添加结束take
		_, err := svcCtx.TaskModel.TransInsert(ctx, sqlx, task)
		if err != nil {
			return err
		}

		// 更新流程实例
		procInst.EndTime = sql.NullInt64{
			Int64: currentTime,
			Valid: true,
		}
		procInst.IsFinished = 1
		err = svcCtx.ProcinstModel.TransUpdate(ctx, sqlx, procInst)
		if err != nil {
			return err
		}
	}
	return nil
}

// MoveToPrevStage MoveToPrevStage
// 驳回
func MoveToPrevStage(ctx context.Context, svcCtx *svc.ServiceContext, sqlx sqlx.Session,
	nodeInfos []*model.NodeInfo, task *model.Task, procInst *model.Procinst, tenantId string, step int) error {
	var currentTime = time.Now().UnixMilli() // 当前时间
	if step == 0 {
		//结束流程
		task.IsFinished = 1
		task.UnCompleteNum = 0
		task.AgreeNum = 1
		task.NodeId = "结束"
		task.ClaimTime = sql.NullInt64{
			Int64: currentTime,
			Valid: true,
		}
		_, err := svcCtx.TaskModel.TransInsert(ctx, sqlx, task)
		if err != nil {
			return err
		}
		// 更新流程实例
		procInst.EndTime = sql.NullInt64{
			Int64: currentTime,
			Valid: true,
		}
		procInst.IsFinished = 1
		err = svcCtx.ProcinstModel.TransUpdate(ctx, sqlx, procInst)
		if err != nil {
			return err
		}
	} else {
		// 生成新的上一个任务
		_, err := svcCtx.TaskModel.TransInsert(ctx, sqlx, task)
		if err != nil {
			return err
		}
	}

	// 更新流程实例
	err := svcCtx.ProcinstModel.TransUpdate(ctx, sqlx, procInst)
	if err != nil {
		return err
	}
	identitylink := &model.Identitylink{
		Id:         uuid.NewV4().String(),
		UserId:     nodeInfos[step].AssigneeID,
		UserName:   nodeInfos[step].AssigneeName,
		TaskId:     task.Id,
		Step:       step,
		ProcinstId: procInst.Id,
		TenantId:   tenantId,
	}
	// 添加
	_, err = svcCtx.IdentitylinkModel.TransInsert(ctx, sqlx, identitylink)
	if err != nil {
		return err
	}

	return nil
}
