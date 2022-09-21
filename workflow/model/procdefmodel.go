package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	procdefFieldNames          = builder.RawFieldNames(&Procdef{})
	procdefRows                = strings.Join(procdefFieldNames, ",")
	procdefRowsExpectAutoSet   = strings.Join(stringx.Remove(procdefFieldNames, "`create_time`", "`update_time`"), ",")
	procdefRowsWithPlaceHolder = strings.Join(stringx.Remove(procdefFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheProcdefIdPrefix = "cache:procdef:id:"
)

type (
	ProcdefModel interface {
		Insert(data *Procdef) (sql.Result, error)
		TxInsert(tx *sql.Tx, data *Procdef) (sql.Result, error)
		FindOne(id string) (*Procdef, error)
		Update(data *Procdef) error
		Delete(id string) error
		FindList(current, pageSize int64, name string, procType string, tenantId string) (*[]Procdef, error)
		Count(name string, procType string, tenantId string) int64
	}

	defaultProcdefModel struct {
		sqlc.CachedConn
		table string
	}

	Procdef struct {
		Id          string         `db:"id"`           // 流程ID
		CreatedAt   time.Time      `db:"created_at"`   // 创建时间
		CreatedName string         `db:"created_name"` // 创建人
		Name        string         `db:"name"`         // 流程名称
		Data        sql.NullString `db:"data"`         // 数据
		ProcType    string         `db:"proc_type"`    // 流程类型
		Resource    string         `db:"resource"`     // 流程内容
		TenantId    string         `db:"tenant_id"`    // 租户ID
	}
)

func NewProcdefModel(conn sqlx.SqlConn, c cache.CacheConf) ProcdefModel {
	return &defaultProcdefModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`procdef`",
	}
}

func (m *defaultProcdefModel) TxInsert(tx *sql.Tx, data *Procdef) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, procdefRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Id, data.CreatedAt, data.CreatedName, data.Name, data.Data, data.ProcType, data.Resource, data.TenantId)

	return ret, err
}

func (m *defaultProcdefModel) Insert(data *Procdef) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, procdefRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Id, data.CreatedAt, data.CreatedName, data.Name, data.Data, data.ProcType, data.Resource, data.TenantId)

	return ret, err
}

func (m *defaultProcdefModel) FindOne(id string) (*Procdef, error) {
	procdefIdKey := fmt.Sprintf("%s%v", cacheProcdefIdPrefix, id)
	var resp Procdef
	err := m.QueryRow(&resp, procdefIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", procdefRows, m.table)
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

func (m *defaultProcdefModel) Update(data *Procdef) error {
	procdefIdKey := fmt.Sprintf("%s%v", cacheProcdefIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, procdefRowsWithPlaceHolder)
		return conn.Exec(query, data.CreatedAt, data.CreatedName, data.Name, data.Data, data.ProcType, data.Resource, data.TenantId, data.Id)
	}, procdefIdKey)
	return err
}

func (m *defaultProcdefModel) Delete(id string) error {

	procdefIdKey := fmt.Sprintf("%s%v", cacheProcdefIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, procdefIdKey)
	return err
}

func (m *defaultProcdefModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheProcdefIdPrefix, primary)
}

func (m *defaultProcdefModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", procdefRows, m.table)
	return conn.QueryRow(v, query, primary)
}

func (m *defaultProcdefModel) FindList(current, pageSize int64, procType string,
	name string, tenantId string) (*[]Procdef, error) {
	var resp []Procdef
	var where string
	if len(tenantId) > 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "tenant_id", tenantId)
	}
	if len(name) > 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "name", "%"+name+"%")
	}
	if len(procType) > 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "proc_type", "%"+procType+"%")
	}
	query := fmt.Sprintf("select %s from %s where  1=1 %s  ORDER BY created_at DESC limit ?,? ", procdefRows, m.table, where)
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

func (m *defaultProcdefModel) Count(name string, procType string, tenantId string) int64 {
	var count int64
	var where string
	if len(tenantId) > 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "tenant_id", tenantId)
	}
	if len(name) > 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "name", "%"+name+"%")
	}
	if len(procType) > 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "proc_type", "%"+procType+"%")
	}
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
