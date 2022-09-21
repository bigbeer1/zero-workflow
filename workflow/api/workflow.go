package main

import (
	"flag"
	"fmt"
	zero_handler "github.com/zeromicro/go-zero/rest/handler"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
	"net/http"
	"time"
	"zero-workflow/common"
	"zero-workflow/workflow/api/internal/config"
	"zero-workflow/workflow/api/internal/handler"
	"zero-workflow/workflow/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/workflow-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *common.CodeError:
			return http.StatusOK, e.Result()
		default:
			return http.StatusInternalServerError, e.Error()
		}
	})
	// 设置日志输出 接口慢时间
	zrpc.SetClientSlowThreshold(time.Second * 2000)
	zero_handler.SetSlowThreshold(time.Second * 2000)

	handler.RegisterHandlers(server, ctx)

	fmt.Println(fmt.Sprintf("Starting %s at %s:%d", c.Name, c.Host, c.Port))
	server.Start()
}
