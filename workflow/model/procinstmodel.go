package model

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	procinstFieldNames          = builder.RawFieldNames(&Procinst{})
	procinstRows                = strings.Join(procinstFieldNames, ",")
	procinstRowsExpectAutoSet   = strings.Join(stringx.Remove(procinstFieldNames, "`create_time`", "`update_time`"), ",")
	procinstRowsWithPlaceHolder = strings.Join(stringx.Remove(procinstFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheProcinstIdPrefix = "cache:procinst:id:"
)

type (
	ProcinstModel interface {
		Insert(data *Procinst) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *Procinst) (sql.Result, error)
		FindOne(id string) (*Procinst, error)
		FindOneData(id string) (*ProcinstData, error)
		Update(data *Procinst) error
		TransUpdate(ctx context.Context, session sqlx.Session, data *Procinst) error
		Delete(id string) error
		TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error
		FindList(current, pageSize int64, procType, procdefName, title string, startTime,
			endTime int64, startUserName string, isFinished int64, tenantId string) (*[]ProcinstData, error)
		Count(procType, procdefName, title string, startTime, endTime int64, startUserName string, isFinished int64, tenantId string) int64
	}

	defaultProcinstModel struct {
		sqlc.CachedConn
		table string
	}

	Procinst struct {
		Id            string        `db:"id"`              // 流程实例ID
		ProcType      string        `db:"proc_type"`       // 流程类型
		ProcdefName   string        `db:"procdef_name"`    // 流程名
		Title         string        `db:"title"`           // 标题
		NodeId        string        `db:"node_id"`         // 当前步骤node_id
		TaskId        string        `db:"task_id"`         // 当前任务
		StartTime     int64         `db:"start_time"`      // 开始时间
		EndTime       sql.NullInt64 `db:"end_time"`        // 结束时间
		StartUserId   string        `db:"start_user_id"`   // 开始用户id
		StartUserName string        `db:"start_user_name"` // 开始用户名
		IsFinished    int64         `db:"is_finished"`     // 是否完成
		TenantId      string        `db:"tenant_id"`       // 租户ID
	}

	ProcinstJson struct {
		Id            string `json:"id"`              // 流程实例ID
		ProcType      string `json:"proc_type"`       // 流程类型
		ProcdefName   string `json:"procdef_name"`    // 流程名
		Title         string `json:"title"`           // 标题
		NodeId        string `json:"node_id"`         // 当前步骤node_id
		TaskId        string `json:"task_id"`         // 当前任务
		StartTime     int64  `json:"start_time"`      // 开始时间
		EndTime       int64  `json:"end_time"`        // 结束时间
		StartUserId   string `json:"start_user_id"`   // 开始用户id
		StartUserName string `json:"start_user_name"` // 开始用户名
		IsFinished    int64  `json:"is_finished"`     // 是否完成
		TenantId      string `json:"tenant_id"`       // 租户ID
	}

	ProcinstData struct {
		Id            string        `json:"id"`              // 流程实例ID
		ProcType      string        `json:"proc_type"`       // 流程类型
		ProcdefName   string        `json:"procdef_name"`    // 流程名
		Title         string        `json:"title"`           // 标题
		StartTime     int64         `json:"start_time"`      // 开始时间
		EndTime       sql.NullInt64 `json:"end_time"`        // 结束时间
		StartUserId   string        `json:"start_user_id"`   // 开始用户id
		StartUserName string        `json:"start_user_name"` // 开始用户名
		IsFinished    int64         `json:"is_finished"`     // 是否完成
		TaskId        string        `json:"task_id"`         // 当前任务
		NodeId        string        `json:"node_id"`         // 当前步骤node_id
		NodeName      string        `json:"node_name"`       // 节点名字
		NodeType      string        `json:"node_type"`       // 节点类型
		Step          int64         `json:"step"`            // 第几步
		AssigneeId    string        `json:"assignee_id"`     // 审批人名
		AssigneeName  string        `json:"assignee_name"`   // 审批人名
	}
)

func NewProcinstModel(conn sqlx.SqlConn, c cache.CacheConf) ProcinstModel {
	return &defaultProcinstModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`procinst`",
	}
}

func (m *defaultProcinstModel) Insert(data *Procinst) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, procinstRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Id, data.ProcType, data.ProcdefName, data.Title, data.NodeId, data.TaskId, data.StartTime, data.EndTime, data.StartUserId, data.StartUserName, data.IsFinished, data.TenantId)

	return ret, err
}

func (m *defaultProcinstModel) TransInsert(ctx context.Context, session sqlx.Session, data *Procinst) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, procinstRowsExpectAutoSet)
	ret, err := session.ExecCtx(ctx, query, data.Id, data.ProcType, data.ProcdefName, data.Title, data.NodeId, data.TaskId, data.StartTime, data.EndTime, data.StartUserId, data.StartUserName, data.IsFinished, data.TenantId)

	return ret, err
}

func (m *defaultProcinstModel) FindOneData(id string) (*ProcinstData, error) {
	procinstIdKey := fmt.Sprintf("%s%v", cacheProcinstIdPrefix, id)
	var resp ProcinstData
	err := m.QueryRow(&resp, procinstIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("SELECT  " +
			"procinst.id, " +
			"procinst.proc_type, " +
			"procinst.procdef_name, " +
			"procinst.title, " +
			"procinst.start_time, " +
			"procinst.end_time, " +
			"procinst.start_user_id, " +
			"procinst.start_user_name, " +
			"procinst.is_finished, " +
			"procinst.task_id, " +
			"procinst.node_id, " +
			"task.node_name, " +
			"task.node_type, " +
			"task.step, " +
			"task.assignee_id, " +
			"task.assignee_name " +
			"FROM " +
			"procinst " +
			"LEFT JOIN " +
			"task " +
			"ON " +
			"procinst.task_id = task.id " +
			"WHERE  " +
			"procinst.id = ?  limit 1")
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

func (m *defaultProcinstModel) FindOne(id string) (*Procinst, error) {
	var resp Procinst
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", procinstRows, m.table)

	err := m.CachedConn.QueryRowNoCache(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}

}

func (m *defaultProcinstModel) Update(data *Procinst) error {
	procinstIdKey := fmt.Sprintf("%s%v", cacheProcinstIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, procinstRowsWithPlaceHolder)
		return conn.Exec(query, data.ProcType, data.ProcdefName, data.Title, data.NodeId, data.TaskId, data.StartTime, data.EndTime, data.StartUserId, data.StartUserName, data.IsFinished, data.TenantId, data.Id)
	}, procinstIdKey)
	return err
}

func (m *defaultProcinstModel) TransUpdate(ctx context.Context, session sqlx.Session, data *Procinst) error {
	procinstIdKey := fmt.Sprintf("%s%v", cacheProcinstIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, procinstRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.ProcType, data.ProcdefName, data.Title, data.NodeId, data.TaskId, data.StartTime, data.EndTime, data.StartUserId, data.StartUserName, data.IsFinished, data.TenantId, data.Id)
	}, procinstIdKey)
	return err
}

func (m *defaultProcinstModel) Delete(id string) error {

	procinstIdKey := fmt.Sprintf("%s%v", cacheProcinstIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, procinstIdKey)
	return err
}

func (m *defaultProcinstModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheProcinstIdPrefix, primary)
}

func (m *defaultProcinstModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", procinstRows, m.table)
	return conn.QueryRow(v, query, primary)
}

func (m *defaultProcinstModel) TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error {
	return m.Transact(func(s sqlx.Session) error {
		return fn(ctx, s)
	})
}

func (m *defaultProcinstModel) FindList(current, pageSize int64, procType, procdefName, title string, startTime,
	endTime int64, startUserName string, isFinished int64, tenantId string) (*[]ProcinstData, error) {
	var resp []ProcinstData
	var where string
	if len(tenantId) != 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "procinst.tenant_id", tenantId)
	}
	if len(procType) > 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "procinst.proc_type", "%"+procType+"%")
	}
	if len(procdefName) > 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "procinst.procdef_name", "%"+procdefName+"%")
	}
	if len(title) > 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "procinst.title", "%"+title+"%")
	}
	if startTime != 0 {
		where += fmt.Sprintf(" AND %s >= '%s'", "procinst.start_time", strconv.FormatInt(startTime, 10))
	}
	if endTime != 0 {
		where += fmt.Sprintf(" AND %s <= '%s'", "procinst.end_time", strconv.FormatInt(endTime, 10))
	}
	if len(startUserName) > 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "procinst.start_user_name", "%"+startUserName+"%")
	}

	if isFinished != 99 {
		where += fmt.Sprintf(" AND %s = '%s'", "procinst.is_finished", strconv.FormatInt(isFinished, 10))
	}

	query := fmt.Sprintf("SELECT  "+
		"procinst.id, "+
		"procinst.proc_type, "+
		"procinst.procdef_name, "+
		"procinst.title, "+
		"procinst.start_time, "+
		"procinst.end_time, "+
		"procinst.start_user_id, "+
		"procinst.start_user_name, "+
		"procinst.is_finished, "+
		"procinst.task_id, "+
		"procinst.node_id, "+
		"task.node_name, "+
		"task.node_type, "+
		"task.step, "+
		"task.assignee_id, "+
		"task.assignee_name "+
		"FROM "+
		"procinst "+
		"LEFT JOIN "+
		"task "+
		"ON "+
		"procinst.task_id = task.id "+
		"WHERE  1=1 "+
		"%s  limit ?, ? ", where)
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

func (m *defaultProcinstModel) Count(procType, procdefName, title string, startTime, endTime int64,
	startUserName string, isFinished int64, tenantId string) int64 {
	var count int64
	var where string
	if len(tenantId) != 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "tenant_id", tenantId)
	}
	if len(procType) > 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "proc_type", "%"+procType+"%")
	}
	if len(procdefName) > 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "procdef_name", "%"+procdefName+"%")
	}
	if len(title) > 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "title", "%"+title+"%")
	}
	if startTime != 0 {
		where += fmt.Sprintf(" AND %s >= '%s'", "start_time", strconv.FormatInt(startTime, 10))
	}
	if endTime != 0 {
		where += fmt.Sprintf(" AND %s <= '%s'", "end_time", strconv.FormatInt(endTime, 10))
	}
	if len(startUserName) > 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "start_user_name", "%"+startUserName+"%")
	}

	if isFinished != 99 {
		where += fmt.Sprintf(" AND %s = '%s'", "procinst.is_finished", strconv.FormatInt(isFinished, 10))
	}

	query := fmt.Sprintf("SELECT count(*) as count from %s where 1=1  %s", m.table, where)
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
