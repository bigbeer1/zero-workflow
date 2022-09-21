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
	identitylinkFieldNames          = builder.RawFieldNames(&Identitylink{})
	identitylinkRows                = strings.Join(identitylinkFieldNames, ",")
	identitylinkRowsExpectAutoSet   = strings.Join(stringx.Remove(identitylinkFieldNames, "`create_time`", "`update_time`"), ",")
	identitylinkRowsWithPlaceHolder = strings.Join(stringx.Remove(identitylinkFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheIdentitylinkIdPrefix = "cache:identitylink:id:"
)

type (
	IdentitylinkModel interface {
		Insert(data *Identitylink) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *Identitylink) (sql.Result, error)
		FindOne(id string) (*Identitylink, error)
		FindOneByUserIdAndTaskIdAndTenantId(userId, taskId, tenantId string) (*Identitylink, error)
		FindOneByProcInstIDAndGroupAndTenantId(procinstId, tenantId string) (*Identitylink, error)
		Update(data *Identitylink) error
		TransUpdate(ctx context.Context, session sqlx.Session, data *Identitylink) error
		Delete(id string) error
		TransDelete(ctx context.Context, session sqlx.Session, id string) error
	}

	defaultIdentitylinkModel struct {
		sqlc.CachedConn
		table string
	}

	Identitylink struct {
		Id         string         `db:"id"`          // 身份ID
		UserId     string         `db:"user_id"`     // 用户id
		UserName   string         `db:"user_name"`   // 用户昵称
		TaskId     string         `db:"task_id"`     // 任务ID
		Step       int            `db:"step"`        // 第几步
		IsAgree    int64          `db:"is_agree"`    // 是否同意
		ProcinstId string         `db:"procinst_id"` // 流程实例ID
		Comment    sql.NullString `db:"comment"`     // 评论
		TenantId   string         `db:"tenant_id"`   // 租户ID
	}
)

func NewIdentitylinkModel(conn sqlx.SqlConn, c cache.CacheConf) IdentitylinkModel {
	return &defaultIdentitylinkModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`identitylink`",
	}
}

func (m *defaultIdentitylinkModel) Insert(data *Identitylink) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, identitylinkRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Id, data.UserId, data.UserName, data.TaskId, data.Step, data.IsAgree, data.ProcinstId, data.Comment, data.TenantId)

	return ret, err
}

func (m *defaultIdentitylinkModel) TransInsert(ctx context.Context, session sqlx.Session, data *Identitylink) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, identitylinkRowsExpectAutoSet)
	ret, err := session.ExecCtx(ctx, query, data.Id, data.UserId, data.UserName, data.TaskId, data.Step, data.IsAgree, data.ProcinstId, data.Comment, data.TenantId)

	return ret, err
}

func (m *defaultIdentitylinkModel) FindOne(id string) (*Identitylink, error) {
	identitylinkIdKey := fmt.Sprintf("%s%v", cacheIdentitylinkIdPrefix, id)
	var resp Identitylink
	err := m.QueryRow(&resp, identitylinkIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", identitylinkRows, m.table)
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

func (m *defaultIdentitylinkModel) Update(data *Identitylink) error {
	identitylinkIdKey := fmt.Sprintf("%s%v", cacheIdentitylinkIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, identitylinkRowsWithPlaceHolder)
		return conn.Exec(query, data.UserId, data.UserName, data.TaskId, data.Step, data.IsAgree, data.ProcinstId, data.Comment, data.TenantId, data.Id)
	}, identitylinkIdKey)
	return err
}

func (m *defaultIdentitylinkModel) TransUpdate(ctx context.Context, session sqlx.Session, data *Identitylink) error {
	identitylinkIdKey := fmt.Sprintf("%s%v", cacheIdentitylinkIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, identitylinkRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.UserId, data.UserName, data.TaskId, data.Step, data.IsAgree, data.ProcinstId, data.Comment, data.TenantId, data.Id)
	}, identitylinkIdKey)
	return err
}

func (m *defaultIdentitylinkModel) Delete(id string) error {

	identitylinkIdKey := fmt.Sprintf("%s%v", cacheIdentitylinkIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, identitylinkIdKey)
	return err
}

func (m *defaultIdentitylinkModel) TransDelete(ctx context.Context, session sqlx.Session, id string) error {

	identitylinkIdKey := fmt.Sprintf("%s%v", cacheIdentitylinkIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return session.ExecCtx(ctx, query, id)
	}, identitylinkIdKey)
	return err
}

func (m *defaultIdentitylinkModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheIdentitylinkIdPrefix, primary)
}

func (m *defaultIdentitylinkModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", identitylinkRows, m.table)
	return conn.QueryRow(v, query, primary)
}

func (m *defaultIdentitylinkModel) FindOneByProcInstIDAndGroupAndTenantId(procinstId, tenantId string) (*Identitylink, error) {
	var resp Identitylink
	var where string
	if len(tenantId) > 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "tenant_id", tenantId)
	}
	if len(procinstId) > 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "procinst_id", procinstId)
	}
	query := fmt.Sprintf("select %s from %s where 1=1  %s limit 1 ", identitylinkRows, m.table, where)
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

func (m *defaultIdentitylinkModel) FindOneByUserIdAndTaskIdAndTenantId(userId, taskId, tenantId string) (*Identitylink, error) {
	var resp Identitylink
	var where string
	if len(tenantId) > 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "tenant_id", tenantId)
	}
	if len(userId) > 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "user_id", userId)
	}
	if len(taskId) > 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "task_id", taskId)
	}
	query := fmt.Sprintf("select %s from %s where 1=1  %s limit 1 ", identitylinkRows, m.table, where)
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
