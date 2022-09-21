package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero-workflow/workflow/api/internal/config"
	"zero-workflow/workflow/rpc/workflow"
)

type ServiceContext struct {
	Config      config.Config
	WorkflowRpc workflow.Workflow
}

func NewServiceContext(c config.Config) *ServiceContext {
	c.WorkflowRpc.Timeout = 500000
	return &ServiceContext{
		Config:      c,
		WorkflowRpc: workflow.NewWorkflow(zrpc.MustNewClient(c.WorkflowRpc)),
	}
}
