package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zeromicro/go-zero/core/logx"

	"gozero-sso-service/application/config"
	"gozero-sso-service/application/handler"
	"gozero-sso-service/application/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "application/etc/Main.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Rest.Log)
	logx.AddWriter(logx.NewWriter(os.Stdout))

	server := rest.MustNewServer(c.Rest.RestConf)
	defer server.Stop()

	serviceCtx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, serviceCtx)
	defer serviceCtx.Close()

	fmt.Printf("Starting server at %s:%d...\n", c.Rest.Host, c.Rest.Port)
	server.Start()
}
