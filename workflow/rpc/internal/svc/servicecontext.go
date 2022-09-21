package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-workflow/workflow/model"
	"zero-workflow/workflow/rpc/internal/config"
)

type ServiceContext struct {
	Config            config.Config
	ProcdefModel      model.ProcdefModel
	ProcinstModel     model.ProcinstModel
	TaskModel         model.TaskModel
	IdentitylinkModel model.IdentitylinkModel
	ExecutionModel    model.ExecutionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:            c,
		ProcdefModel:      model.NewProcdefModel(conn, c.CacheRedis),
		ProcinstModel:     model.NewProcinstModel(conn, c.CacheRedis),
		TaskModel:         model.NewTaskModel(conn, c.CacheRedis),
		IdentitylinkModel: model.NewIdentitylinkModel(conn, c.CacheRedis),
		ExecutionModel:    model.NewExecutionModel(conn, c.CacheRedis),
	}
}
