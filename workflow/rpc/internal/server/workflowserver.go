// Code generated by goctl. DO NOT EDIT!
// Source: workflow.proto

package server

import (
	"context"
	"zero-workflow/workflow/rpc/internal/logic"
	"zero-workflow/workflow/rpc/internal/svc"
	"zero-workflow/workflow/rpc/workflowclient"
)

type WorkflowServer struct {
	svcCtx *svc.ServiceContext
	workflowclient.UnimplementedWorkflowServer
}

func NewWorkflowServer(svcCtx *svc.ServiceContext) *WorkflowServer {
	return &WorkflowServer{
		svcCtx: svcCtx,
	}
}

// -------------------------添加流程定义----------------------
func (s *WorkflowServer) ProcdefAdd(ctx context.Context, in *workflowclient.ProcdefAddReq) (*workflowclient.CommonResp, error) {
	l := logic.NewProcdefAddLogic(ctx, s.svcCtx)
	return l.ProcdefAdd(in)
}

func (s *WorkflowServer) ProcdefDelete(ctx context.Context, in *workflowclient.ProcdefDeleteReq) (*workflowclient.CommonResp, error) {
	l := logic.NewProcdefDeleteLogic(ctx, s.svcCtx)
	return l.ProcdefDelete(in)
}

func (s *WorkflowServer) ProcdefFindList(ctx context.Context, in *workflowclient.ProcdefFindListReq) (*workflowclient.ProcdefFindListResp, error) {
	l := logic.NewProcdefFindListLogic(ctx, s.svcCtx)
	return l.ProcdefFindList(in)
}

func (s *WorkflowServer) ProcdefFindOne(ctx context.Context, in *workflowclient.ProcdefFindOneReq) (*workflowclient.ProcdefFindOneResp, error) {
	l := logic.NewProcdefFindOneLogic(ctx, s.svcCtx)
	return l.ProcdefFindOne(in)
}

//  启动流程
func (s *WorkflowServer) ProcessStart(ctx context.Context, in *workflowclient.ProcessStartReq) (*workflowclient.CommonResp, error) {
	l := logic.NewProcessStartLogic(ctx, s.svcCtx)
	return l.ProcessStart(in)
}

//  -----------------------流程实例-----------------------
func (s *WorkflowServer) ProcinstFindOne(ctx context.Context, in *workflowclient.ProcinstFindOneReq) (*workflowclient.ProcinstFindOneResp, error) {
	l := logic.NewProcinstFindOneLogic(ctx, s.svcCtx)
	return l.ProcinstFindOne(in)
}

func (s *WorkflowServer) ProcinstFindList(ctx context.Context, in *workflowclient.ProcinstFindListReq) (*workflowclient.ProcinstFindListResp, error) {
	l := logic.NewProcinstFindListLogic(ctx, s.svcCtx)
	return l.ProcinstFindList(in)
}

func (s *WorkflowServer) ProcinstClose(ctx context.Context, in *workflowclient.ProcinstCloseReq) (*workflowclient.CommonResp, error) {
	l := logic.NewProcinstCloseLogic(ctx, s.svcCtx)
	return l.ProcinstClose(in)
}

//  -----------------------任务--------------------------
func (s *WorkflowServer) TaskComplete(ctx context.Context, in *workflowclient.TaskCompleteReq) (*workflowclient.CommonResp, error) {
	l := logic.NewTaskCompleteLogic(ctx, s.svcCtx)
	return l.TaskComplete(in)
}

func (s *WorkflowServer) TaskFindListByUserId(ctx context.Context, in *workflowclient.TaskFindListByUserIdReq) (*workflowclient.TaskFindListByUserIdResp, error) {
	l := logic.NewTaskFindListByUserIdLogic(ctx, s.svcCtx)
	return l.TaskFindListByUserId(in)
}

func (s *WorkflowServer) TaskFindListByProcinstId(ctx context.Context, in *workflowclient.TaskFindListByProcinstIdReq) (*workflowclient.TaskFindListByProcinstIdResp, error) {
	l := logic.NewTaskFindListByProcinstIdLogic(ctx, s.svcCtx)
	return l.TaskFindListByProcinstId(in)
}

//  -----------------------实例内容-----------------------------
func (s *WorkflowServer) ExecutionFindOneByProcinstId(ctx context.Context, in *workflowclient.ExecutionFindOneByProcinstIdReq) (*workflowclient.ExecutionFindOneByProcinstIdResp, error) {
	l := logic.NewExecutionFindOneByProcinstIdLogic(ctx, s.svcCtx)
	return l.ExecutionFindOneByProcinstId(in)
}