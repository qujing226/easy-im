package main

import (
	"flag"
	"fmt"
	"github.com/peninsula12/easy-im/go-im/apps/im/ws/internal/config"
	"github.com/peninsula12/easy-im/go-im/apps/im/ws/internal/handler"
	"github.com/peninsula12/easy-im/go-im/apps/im/ws/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/im/ws/websocket"
	"github.com/peninsula12/easy-im/go-im/pkg/configserver"
	"time"
)

var configFile = flag.String("f", "etc/dev/im.yaml", "the config file")

//var configFile = flag.String("f", "apps/im/ws/etc/dev/im.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	var configs = "im-ws.yaml"
	err := configserver.NewConfigServer(*configFile, configserver.NewSail(&configserver.Config{
		ETCDEndpoints: "118.178.120.11:3379",
		ProjectKey:    "3c46a0407be60a1f00731ab8e9575df2",
		Namespace:     "im",
		Configs:       configs,
		//ConfigFilePath: "../etc/conf",
		// 本地测试使用以下配置
		ConfigFilePath: "./etc/conf",
		LogLevel:       "DEBUG",
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

	if err := c.SetUp(); err != nil {
		panic(err)
	}

	ctx := svc.NewServiceContext(c)
	srv := websocket.NewServer(c.ListenOn,
		websocket.WithAuthentication(handler.NewJwtAuth(ctx)),
		//websocket.WithServerMaxConnectionIdle(180*time.Second),
		websocket.WithServerMaxConnectionIdle(180*time.Minute),
		websocket.WithServerAck(websocket.OnlyAck),
	)
	defer srv.Stop()

	handler.RegisterHandlers(srv, ctx)

	fmt.Printf("starting server at %v ...\n", c.ListenOn)
	srv.Start()
}
