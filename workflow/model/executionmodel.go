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
	executionFieldNames          = builder.RawFieldNames(&Execution{})
	executionRows                = strings.Join(executionFieldNames, ",")
	executionRowsExpectAutoSet   = strings.Join(stringx.Remove(executionFieldNames, "`create_time`", "`update_time`"), ",")
	executionRowsWithPlaceHolder = strings.Join(stringx.Remove(executionFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheExecutionIdPrefix = "cache:execution:id:"
)

type (
	ExecutionModel interface {
		Insert(data *Execution) (sql.Result, error)
		TansInsert(ctx context.Context, session sqlx.Session, data *Execution) (sql.Result, error)
		FindOne(id string) (*Execution, error)
		FindOneByProcinstIdAndTenantId(procinstId string, tenantId string) (*Execution, error)
		Update(data *Execution) error
		Delete(id string) error
	}

	defaultExecutionModel struct {
		sqlc.CachedConn
		table string
	}

	Execution struct {
		Id          string `db:"id"`           // 执行ID
		ProcinstId  string `db:"procinst_id"`  // 实例ID
		ProcdefName string `db:"procdef_name"` // 流程名
		NodeInfos   string `db:"node_infos"`   // 节点信息
		StartTime   int64  `db:"start_time"`   // 开始时间
		TenantId    string `db:"tenant_id"`    // 租户ID
	}
)

func NewExecutionModel(conn sqlx.SqlConn, c cache.CacheConf) ExecutionModel {
	return &defaultExecutionModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`execution`",
	}
}

func (m *defaultExecutionModel) Insert(data *Execution) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, executionRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Id, data.ProcinstId, data.ProcdefName, data.NodeInfos, data.StartTime, data.TenantId)

	return ret, err
}

func (m *defaultExecutionModel) TansInsert(ctx context.Context, session sqlx.Session, data *Execution) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, executionRowsExpectAutoSet)
	ret, err := session.ExecCtx(ctx, query, data.Id, data.ProcinstId, data.ProcdefName, data.NodeInfos, data.StartTime, data.TenantId)

	return ret, err
}

func (m *defaultExecutionModel) FindOne(id string) (*Execution, error) {
	executionIdKey := fmt.Sprintf("%s%v", cacheExecutionIdPrefix, id)
	var resp Execution
	err := m.QueryRow(&resp, executionIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", executionRows, m.table)
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

func (m *defaultExecutionModel) Update(data *Execution) error {
	executionIdKey := fmt.Sprintf("%s%v", cacheExecutionIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, executionRowsWithPlaceHolder)
		return conn.Exec(query, data.ProcinstId, data.ProcdefName, data.NodeInfos, data.StartTime, data.TenantId, data.Id)
	}, executionIdKey)
	return err
}

func (m *defaultExecutionModel) Delete(id string) error {

	executionIdKey := fmt.Sprintf("%s%v", cacheExecutionIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, executionIdKey)
	return err
}

func (m *defaultExecutionModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheExecutionIdPrefix, primary)
}

func (m *defaultExecutionModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", executionRows, m.table)
	return conn.QueryRow(v, query, primary)
}

func (m *defaultExecutionModel) FindOneByProcinstIdAndTenantId(procinstId string, tenantId string) (*Execution, error) {
	var resp Execution
	var where string
	if len(tenantId) > 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "tenant_id", tenantId)
	}
	if len(procinstId) > 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "procinst_id", procinstId)
	}
	query := fmt.Sprintf("select %s from %s where 1=1  %s limit 1 ", executionRows, m.table, where)
	err := m.CachedConn.QueryRowNoCache(&resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
