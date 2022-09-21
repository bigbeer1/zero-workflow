// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package workflowclient

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WorkflowClient is the client API for Workflow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkflowClient interface {
	//-------------------------添加流程定义----------------------
	ProcdefAdd(ctx context.Context, in *ProcdefAddReq, opts ...grpc.CallOption) (*CommonResp, error)
	ProcdefDelete(ctx context.Context, in *ProcdefDeleteReq, opts ...grpc.CallOption) (*CommonResp, error)
	ProcdefFindList(ctx context.Context, in *ProcdefFindListReq, opts ...grpc.CallOption) (*ProcdefFindListResp, error)
	ProcdefFindOne(ctx context.Context, in *ProcdefFindOneReq, opts ...grpc.CallOption) (*ProcdefFindOneResp, error)
	// 启动流程
	ProcessStart(ctx context.Context, in *ProcessStartReq, opts ...grpc.CallOption) (*CommonResp, error)
	// -----------------------流程实例-----------------------
	ProcinstFindOne(ctx context.Context, in *ProcinstFindOneReq, opts ...grpc.CallOption) (*ProcinstFindOneResp, error)
	ProcinstFindList(ctx context.Context, in *ProcinstFindListReq, opts ...grpc.CallOption) (*ProcinstFindListResp, error)
	ProcinstClose(ctx context.Context, in *ProcinstCloseReq, opts ...grpc.CallOption) (*CommonResp, error)
	// -----------------------任务--------------------------
	TaskComplete(ctx context.Context, in *TaskCompleteReq, opts ...grpc.CallOption) (*CommonResp, error)
	TaskFindListByUserId(ctx context.Context, in *TaskFindListByUserIdReq, opts ...grpc.CallOption) (*TaskFindListByUserIdResp, error)
	TaskFindListByProcinstId(ctx context.Context, in *TaskFindListByProcinstIdReq, opts ...grpc.CallOption) (*TaskFindListByProcinstIdResp, error)
	// -----------------------实例内容-----------------------------
	ExecutionFindOneByProcinstId(ctx context.Context, in *ExecutionFindOneByProcinstIdReq, opts ...grpc.CallOption) (*ExecutionFindOneByProcinstIdResp, error)
}

type workflowClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkflowClient(cc grpc.ClientConnInterface) WorkflowClient {
	return &workflowClient{cc}
}

func (c *workflowClient) ProcdefAdd(ctx context.Context, in *ProcdefAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, "/workflowclient.Workflow/ProcdefAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) ProcdefDelete(ctx context.Context, in *ProcdefDeleteReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, "/workflowclient.Workflow/ProcdefDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) ProcdefFindList(ctx context.Context, in *ProcdefFindListReq, opts ...grpc.CallOption) (*ProcdefFindListResp, error) {
	out := new(ProcdefFindListResp)
	err := c.cc.Invoke(ctx, "/workflowclient.Workflow/ProcdefFindList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) ProcdefFindOne(ctx context.Context, in *ProcdefFindOneReq, opts ...grpc.CallOption) (*ProcdefFindOneResp, error) {
	out := new(ProcdefFindOneResp)
	err := c.cc.Invoke(ctx, "/workflowclient.Workflow/ProcdefFindOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) ProcessStart(ctx context.Context, in *ProcessStartReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, "/workflowclient.Workflow/ProcessStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) ProcinstFindOne(ctx context.Context, in *ProcinstFindOneReq, opts ...grpc.CallOption) (*ProcinstFindOneResp, error) {
	out := new(ProcinstFindOneResp)
	err := c.cc.Invoke(ctx, "/workflowclient.Workflow/ProcinstFindOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) ProcinstFindList(ctx context.Context, in *ProcinstFindListReq, opts ...grpc.CallOption) (*ProcinstFindListResp, error) {
	out := new(ProcinstFindListResp)
	err := c.cc.Invoke(ctx, "/workflowclient.Workflow/ProcinstFindList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) ProcinstClose(ctx context.Context, in *ProcinstCloseReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, "/workflowclient.Workflow/ProcinstClose", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) TaskComplete(ctx context.Context, in *TaskCompleteReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, "/workflowclient.Workflow/TaskComplete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) TaskFindListByUserId(ctx context.Context, in *TaskFindListByUserIdReq, opts ...grpc.CallOption) (*TaskFindListByUserIdResp, error) {
	out := new(TaskFindListByUserIdResp)
	err := c.cc.Invoke(ctx, "/workflowclient.Workflow/TaskFindListByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) TaskFindListByProcinstId(ctx context.Context, in *TaskFindListByProcinstIdReq, opts ...grpc.CallOption) (*TaskFindListByProcinstIdResp, error) {
	out := new(TaskFindListByProcinstIdResp)
	err := c.cc.Invoke(ctx, "/workflowclient.Workflow/TaskFindListByProcinstId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) ExecutionFindOneByProcinstId(ctx context.Context, in *ExecutionFindOneByProcinstIdReq, opts ...grpc.CallOption) (*ExecutionFindOneByProcinstIdResp, error) {
	out := new(ExecutionFindOneByProcinstIdResp)
	err := c.cc.Invoke(ctx, "/workflowclient.Workflow/ExecutionFindOneByProcinstId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowServer is the server API for Workflow service.
// All implementations must embed UnimplementedWorkflowServer
// for forward compatibility
type WorkflowServer interface {
	//-------------------------添加流程定义----------------------
	ProcdefAdd(context.Context, *ProcdefAddReq) (*CommonResp, error)
	ProcdefDelete(context.Context, *ProcdefDeleteReq) (*CommonResp, error)
	ProcdefFindList(context.Context, *ProcdefFindListReq) (*ProcdefFindListResp, error)
	ProcdefFindOne(context.Context, *ProcdefFindOneReq) (*ProcdefFindOneResp, error)
	// 启动流程
	ProcessStart(context.Context, *ProcessStartReq) (*CommonResp, error)
	// -----------------------流程实例-----------------------
	ProcinstFindOne(context.Context, *ProcinstFindOneReq) (*ProcinstFindOneResp, error)
	ProcinstFindList(context.Context, *ProcinstFindListReq) (*ProcinstFindListResp, error)
	ProcinstClose(context.Context, *ProcinstCloseReq) (*CommonResp, error)
	// -----------------------任务--------------------------
	TaskComplete(context.Context, *TaskCompleteReq) (*CommonResp, error)
	TaskFindListByUserId(context.Context, *TaskFindListByUserIdReq) (*TaskFindListByUserIdResp, error)
	TaskFindListByProcinstId(context.Context, *TaskFindListByProcinstIdReq) (*TaskFindListByProcinstIdResp, error)
	// -----------------------实例内容-----------------------------
	ExecutionFindOneByProcinstId(context.Context, *ExecutionFindOneByProcinstIdReq) (*ExecutionFindOneByProcinstIdResp, error)
	mustEmbedUnimplementedWorkflowServer()
}

// UnimplementedWorkflowServer must be embedded to have forward compatible implementations.
type UnimplementedWorkflowServer struct {
}

func (UnimplementedWorkflowServer) ProcdefAdd(context.Context, *ProcdefAddReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcdefAdd not implemented")
}
func (UnimplementedWorkflowServer) ProcdefDelete(context.Context, *ProcdefDeleteReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcdefDelete not implemented")
}
func (UnimplementedWorkflowServer) ProcdefFindList(context.Context, *ProcdefFindListReq) (*ProcdefFindListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcdefFindList not implemented")
}
func (UnimplementedWorkflowServer) ProcdefFindOne(context.Context, *ProcdefFindOneReq) (*ProcdefFindOneResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcdefFindOne not implemented")
}
func (UnimplementedWorkflowServer) ProcessStart(context.Context, *ProcessStartReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessStart not implemented")
}
func (UnimplementedWorkflowServer) ProcinstFindOne(context.Context, *ProcinstFindOneReq) (*ProcinstFindOneResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcinstFindOne not implemented")
}
func (UnimplementedWorkflowServer) ProcinstFindList(context.Context, *ProcinstFindListReq) (*ProcinstFindListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcinstFindList not implemented")
}
func (UnimplementedWorkflowServer) ProcinstClose(context.Context, *ProcinstCloseReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcinstClose not implemented")
}
func (UnimplementedWorkflowServer) TaskComplete(context.Context, *TaskCompleteReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskComplete not implemented")
}
func (UnimplementedWorkflowServer) TaskFindListByUserId(context.Context, *TaskFindListByUserIdReq) (*TaskFindListByUserIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskFindListByUserId not implemented")
}
func (UnimplementedWorkflowServer) TaskFindListByProcinstId(context.Context, *TaskFindListByProcinstIdReq) (*TaskFindListByProcinstIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskFindListByProcinstId not implemented")
}
func (UnimplementedWorkflowServer) ExecutionFindOneByProcinstId(context.Context, *ExecutionFindOneByProcinstIdReq) (*ExecutionFindOneByProcinstIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecutionFindOneByProcinstId not implemented")
}
func (UnimplementedWorkflowServer) mustEmbedUnimplementedWorkflowServer() {}

// UnsafeWorkflowServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkflowServer will
// result in compilation errors.
type UnsafeWorkflowServer interface {
	mustEmbedUnimplementedWorkflowServer()
}

func RegisterWorkflowServer(s grpc.ServiceRegistrar, srv WorkflowServer) {
	s.RegisterService(&Workflow_ServiceDesc, srv)
}

func _Workflow_ProcdefAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcdefAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).ProcdefAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflowclient.Workflow/ProcdefAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).ProcdefAdd(ctx, req.(*ProcdefAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_ProcdefDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcdefDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).ProcdefDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflowclient.Workflow/ProcdefDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).ProcdefDelete(ctx, req.(*ProcdefDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_ProcdefFindList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcdefFindListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).ProcdefFindList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflowclient.Workflow/ProcdefFindList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).ProcdefFindList(ctx, req.(*ProcdefFindListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_ProcdefFindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcdefFindOneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).ProcdefFindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflowclient.Workflow/ProcdefFindOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).ProcdefFindOne(ctx, req.(*ProcdefFindOneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_ProcessStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessStartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).ProcessStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflowclient.Workflow/ProcessStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).ProcessStart(ctx, req.(*ProcessStartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_ProcinstFindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcinstFindOneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).ProcinstFindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflowclient.Workflow/ProcinstFindOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).ProcinstFindOne(ctx, req.(*ProcinstFindOneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_ProcinstFindList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcinstFindListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).ProcinstFindList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflowclient.Workflow/ProcinstFindList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).ProcinstFindList(ctx, req.(*ProcinstFindListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_ProcinstClose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcinstCloseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).ProcinstClose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflowclient.Workflow/ProcinstClose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).ProcinstClose(ctx, req.(*ProcinstCloseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_TaskComplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskCompleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).TaskComplete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflowclient.Workflow/TaskComplete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).TaskComplete(ctx, req.(*TaskCompleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_TaskFindListByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskFindListByUserIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).TaskFindListByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflowclient.Workflow/TaskFindListByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).TaskFindListByUserId(ctx, req.(*TaskFindListByUserIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_TaskFindListByProcinstId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskFindListByProcinstIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).TaskFindListByProcinstId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflowclient.Workflow/TaskFindListByProcinstId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).TaskFindListByProcinstId(ctx, req.(*TaskFindListByProcinstIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_ExecutionFindOneByProcinstId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutionFindOneByProcinstIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).ExecutionFindOneByProcinstId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflowclient.Workflow/ExecutionFindOneByProcinstId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).ExecutionFindOneByProcinstId(ctx, req.(*ExecutionFindOneByProcinstIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Workflow_ServiceDesc is the grpc.ServiceDesc for Workflow service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Workflow_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "workflowclient.Workflow",
	HandlerType: (*WorkflowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcdefAdd",
			Handler:    _Workflow_ProcdefAdd_Handler,
		},
		{
			MethodName: "ProcdefDelete",
			Handler:    _Workflow_ProcdefDelete_Handler,
		},
		{
			MethodName: "ProcdefFindList",
			Handler:    _Workflow_ProcdefFindList_Handler,
		},
		{
			MethodName: "ProcdefFindOne",
			Handler:    _Workflow_ProcdefFindOne_Handler,
		},
		{
			MethodName: "ProcessStart",
			Handler:    _Workflow_ProcessStart_Handler,
		},
		{
			MethodName: "ProcinstFindOne",
			Handler:    _Workflow_ProcinstFindOne_Handler,
		},
		{
			MethodName: "ProcinstFindList",
			Handler:    _Workflow_ProcinstFindList_Handler,
		},
		{
			MethodName: "ProcinstClose",
			Handler:    _Workflow_ProcinstClose_Handler,
		},
		{
			MethodName: "TaskComplete",
			Handler:    _Workflow_TaskComplete_Handler,
		},
		{
			MethodName: "TaskFindListByUserId",
			Handler:    _Workflow_TaskFindListByUserId_Handler,
		},
		{
			MethodName: "TaskFindListByProcinstId",
			Handler:    _Workflow_TaskFindListByProcinstId_Handler,
		},
		{
			MethodName: "ExecutionFindOneByProcinstId",
			Handler:    _Workflow_ExecutionFindOneByProcinstId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workflow.proto",
}
