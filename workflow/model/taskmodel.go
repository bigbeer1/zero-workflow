package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	taskFieldNames          = builder.RawFieldNames(&Task{})
	taskRows                = strings.Join(taskFieldNames, ",")
	taskRowsExpectAutoSet   = strings.Join(stringx.Remove(taskFieldNames, "`create_time`", "`update_time`"), ",")
	taskRowsWithPlaceHolder = strings.Join(stringx.Remove(taskFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheTaskIdPrefix = "cache:task:id:"
)

type (
	TaskModel interface {
		Insert(data *Task) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *Task) (sql.Result, error)
		FindOne(id string) (*Task, error)
		Update(data *Task) error
		TransUpdate(ctx context.Context, session sqlx.Session, data *Task) error
		Delete(id string) error
		TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error
		FindListByUserId(current, pageSize int64, assigneeId string) (*[]TaskData, error)
		CountByUserId(assigneeId string) int64
		FindListByProcinstId(procinstId string) (*[]TaskProcinstData, error)
	}

	defaultTaskModel struct {
		sqlc.CachedConn
		table string
	}

	Task struct {
		Id            string        `db:"id"`              // 任务ID
		CreatedAt     int64         `db:"created_at"`      // 创建时间
		ClaimTime     sql.NullInt64 `db:"claim_time"`      // 最近通过时间
		NodeId        string        `db:"node_id"`         // 当前执行流所在的节点ID
		NodeName      string        `db:"node_name"`       // 步骤名
		NodeType      string        `db:"node_type"`       // 步骤类型
		Step          int           `db:"step"`            // 第几步
		ProcinstId    string        `db:"procinst_id"`     // 流程实例id
		AssigneeId    string        `db:"assignee_id"`     // 受理人ID
		AssigneeName  string        `db:"assignee_name"`   // 受理人姓名
		UnCompleteNum int64         `db:"un_complete_num"` // 需要审批的数
		AgreeNum      int64         `db:"agree_num"`       // 审批通过数
		IsFinished    int64         `db:"is_finished"`     // 是否完成
		TenantId      string        `db:"tenant_id"`       // 租户ID
	}

	TaskData struct {
		TaskId        string `json:"task_id"`         // 任务ID
		Procinstid    string `json:"procinst_id"`     // 流程实例ID
		CreatedAt     int64  `json:"created_at"`      // take创建时间
		NodeId        string `json:"node_id"`         // 当前执行流所在的节点ID
		Step          int    `json:"step"`            // 第几步
		AgreeNum      string `json:"agree_num"`       // 同意数
		ProcType      string `json:"proc_type"`       // 流程类型
		ProcinstName  string `json:"procinst_name"`   // 流程名
		ProcinstTitle string `json:"procinst_title"`  // 流程标题
		StartTime     int64  `json:"start_time"`      // 开始时间
		StartUserName string `json:"start_user_name"` // 创建人名
		StartUserId   string `json:"start_user_id"`   // 创建人Id
	}

	TaskProcinstData struct {
		TaskId       string         `json:"task_id"`       // 任务ID
		CreatedAt    int64          `json:"created_at"`    // take创建时间
		ClaimTime    sql.NullInt64  `json:"claim_time"`    // take结束
		NodeId       string         `json:"node_id"`       // 节点名
		NodeName     string         `json:"node_name"`     // 节点名
		NodeType     string         `json:"node_type"`     // 字节类型
		Step         int64          `json:"step"`          // 第几步
		AssigneeName string         `json:"assignee_name"` // 审批人
		IsFinished   int64          `json:"is_finished"`   // 是否完成
		IsAgree      sql.NullInt64  `json:"is_agree"`      // 是否同意
		Comment      sql.NullString `json:"comment"`       // 提交内容
	}
)

func NewTaskModel(conn sqlx.SqlConn, c cache.CacheConf) TaskModel {
	return &defaultTaskModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`task`",
	}
}

func (m *defaultTaskModel) Insert(data *Task) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, taskRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Id, data.CreatedAt, data.ClaimTime, data.NodeId, data.NodeName, data.NodeType, data.Step, data.ProcinstId, data.AssigneeId, data.AssigneeName, data.UnCompleteNum, data.AgreeNum, data.IsFinished, data.TenantId)

	return ret, err
}

func (m *defaultTaskModel) TransInsert(ctx context.Context, session sqlx.Session, data *Task) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, taskRowsExpectAutoSet)
	ret, err := session.ExecCtx(ctx, query, data.Id, data.CreatedAt, data.ClaimTime, data.NodeId, data.NodeName, data.NodeType, data.Step, data.ProcinstId, data.AssigneeId, data.AssigneeName, data.UnCompleteNum, data.AgreeNum, data.IsFinished, data.TenantId)

	return ret, err
}

func (m *defaultTaskModel) FindOne(id string) (*Task, error) {
	taskIdKey := fmt.Sprintf("%s%v", cacheTaskIdPrefix, id)
	var resp Task
	err := m.QueryRow(&resp, taskIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", taskRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTaskModel) Update(data *Task) error {
	taskIdKey := fmt.Sprintf("%s%v", cacheTaskIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, taskRowsWithPlaceHolder)
		return conn.Exec(query, data.CreatedAt, data.ClaimTime, data.NodeId, data.NodeName, data.NodeType, data.Step, data.ProcinstId, data.AssigneeId, data.AssigneeName, data.UnCompleteNum, data.AgreeNum, data.IsFinished, data.TenantId, data.Id)
	}, taskIdKey)
	return err
}

func (m *defaultTaskModel) TransUpdate(ctx context.Context, session sqlx.Session, data *Task) error {
	taskIdKey := fmt.Sprintf("%s%v", cacheTaskIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, taskRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.CreatedAt, data.ClaimTime, data.NodeId, data.NodeName, data.NodeType, data.Step, data.ProcinstId, data.AssigneeId, data.AssigneeName, data.UnCompleteNum, data.AgreeNum, data.IsFinished, data.TenantId, data.Id)
	}, taskIdKey)
	return err
}

func (m *defaultTaskModel) Delete(id string) error {

	taskIdKey := fmt.Sprintf("%s%v", cacheTaskIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, taskIdKey)
	return err
}

func (m *defaultTaskModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTaskIdPrefix, primary)
}

func (m *defaultTaskModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", taskRows, m.table)
	return conn.QueryRow(v, query, primary)
}

func (m *defaultTaskModel) TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error {
	return m.Transact(func(s sqlx.Session) error {
		return fn(ctx, s)
	})
}

func (m *defaultTaskModel) FindListByUserId(current, pageSize int64, assigneeId string) (*[]TaskData, error) {
	var resp []TaskData
	var where string
	if len(assigneeId) > 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "assignee_id", assigneeId)
	}
	query := fmt.Sprintf("SELECT "+
		"task.id as take_id, "+
		"task.procinst_id as procinst_id, "+
		"task.created_at, "+
		"task.node_id, "+
		"task.step, "+
		"task.agree_num, "+
		"procinst.proc_type, "+
		"procinst.procdef_name as procinst_name, "+
		"procinst.title as procinst_title, "+
		"procinst.start_time, "+
		"procinst.start_user_name, "+
		"procinst.start_user_id "+
		"FROM "+
		"task "+
		"LEFT JOIN "+
		"procinst "+
		"ON "+
		"task.procinst_id = procinst.id "+
		"WHERE "+
		"procinst.is_finished=0 "+
		"And "+
		"task.is_finished=0 %s limit ?, ?", where)
	err := m.CachedConn.QueryRowsNoCache(&resp, query, (current-1)*pageSize, pageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTaskModel) CountByUserId(assigneeId string) int64 {
	var count int64
	var where string
	if len(assigneeId) > 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "assignee_id", assigneeId)
	}
	where += fmt.Sprintf(" AND %s = '%s'", "is_finished", "0")
	query := fmt.Sprintf("SELECT count(*) as count from %s where 1=1 %s ", m.table, where)
	err := m.CachedConn.QueryRowNoCache(&count, query)
	switch err {
	case nil:
		return count
	case sqlc.ErrNotFound:
		return 0
	default:
		return 0
	}
}

func (m *defaultTaskModel) FindListByProcinstId(procinstId string) (*[]TaskProcinstData, error) {
	var resp []TaskProcinstData
	var where string
	if len(procinstId) > 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "task.procinst_id", procinstId)
	}
	query := fmt.Sprintf("SELECT "+
		"task.id, "+
		"task.created_at, "+
		"task.claim_time, "+
		"task.node_id, "+
		"task.node_name, "+
		"task.node_type, "+
		"task.step, "+
		"task.assignee_name, "+
		"task.is_finished, "+
		"identitylink.`is_agree`, "+
		"identitylink.`comment` "+
		"FROM "+
		"task "+
		"LEFT JOIN "+
		"identitylink "+
		"ON "+
		"task.id = identitylink.task_id "+
		"WHERE "+
		"1=1 %s", where)
	err := m.CachedConn.QueryRowsNoCache(&resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
