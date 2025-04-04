package main

import (
	"flag"
	"fmt"
	"github.com/peninsula12/easy-im/go-im/pkg/configserver"
	"github.com/peninsula12/easy-im/go-im/pkg/resultx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/peninsula12/easy-im/go-im/apps/im/api/internal/config"
	"github.com/peninsula12/easy-im/go-im/apps/im/api/internal/handler"
	"github.com/peninsula12/easy-im/go-im/apps/im/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/dev/im.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	//conf.MustLoad(*configFile, &c)
	var configs = "im-api.yaml"
	err := configserver.NewConfigServer(*configFile, configserver.NewSail(&configserver.Config{
		ETCDEndpoints:  "118.178.120.11:3379",
		ProjectKey:     "3c46a0407be60a1f00731ab8e9575df2",
		Namespace:      "im",
		Configs:        configs,
		ConfigFilePath: "../etc/conf",
		// 本地测试使用以下配置
		//ConfigFilePath: "./etc/conf",
		LogLevel: "DEBUG",
	})).MustLoad(&c, func(bytes []byte) error {
		var c config.Config
		err := configserver.LoadFromJsonBytes(bytes, &c)
		if err != nil {
			fmt.Println("config read err :", err)
			return nil
		}
		fmt.Printf(configs, "config has changed :%+v \n", c)
		return nil
	})
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
