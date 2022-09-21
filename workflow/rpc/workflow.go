package main

import (
	"flag"
	"fmt"
	zredis "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"zero-workflow/workflow/rpc/internal/config"
	"zero-workflow/workflow/rpc/internal/server"
	"zero-workflow/workflow/rpc/internal/svc"
	"zero-workflow/workflow/rpc/workflowclient"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/workflow.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewWorkflowServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		workflowclient.RegisterWorkflowServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	// 设置日志输出 接口慢时间  rpc
	zrpc.SetServerSlowThreshold(time.Second * 2000)
	// redis
	zredis.SetSlowThreshold(time.Second * 2000)
	// sqlx
	sqlx.SetSlowThreshold(time.Second * 2000)

	defer s.Stop()

	fmt.Println(c.Name, c.ListenOn)

	s.Start()

}
