package main

import (
	"easy-chat/pkg/configserver"
	"easy-chat/pkg/resultx"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"

	"easy-chat/apps/im/api/internal/config"
	"easy-chat/apps/im/api/internal/handler"
	"easy-chat/apps/im/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/dev/im.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	//conf.MustLoad(*configFile, &c)
	err := configserver.NewConfigServer(*configFile, configserver.NewSail(&configserver.Config{
		ETCDEndpoints:  "118.178.120.11:3379",
		ProjectKey:     "98c6f2c2287f4c73cea3d40ae7ec3ff2",
		Namespace:      "im",
		Configs:        "im-api.yaml",
		ConfigFilePath: "./etc/conf",
		LogLevel:       "DEBUG",
	})).MustLoad(&c)
	if err != nil {
		panic(err)
	}

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandlerCtx(resultx.ErrHandler(c.Name))
	httpx.SetOkHandler(resultx.OKHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
