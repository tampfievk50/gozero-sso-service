package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zeromicro/go-zero/core/logx"

	"gozero-sso-service/internal/config"
	"gozero-sso-service/internal/handler"
	"gozero-sso-service/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	_ "gozero-sso-service/internal/bootstrap"
)

var configFile = flag.String("f", "etc/Main.yaml", "the config file")

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

	fmt.Printf("Starting server at %s:%d...\n", c.Rest.Host, c.Rest.Port)
	server.Start()
}
