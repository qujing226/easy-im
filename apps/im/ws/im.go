package main

import (
	"easy-chat/apps/im/ws/internal/config"
	"easy-chat/apps/im/ws/internal/handler"
	"easy-chat/apps/im/ws/internal/svc"
	"easy-chat/apps/im/ws/websocket"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"time"
)

var configFile = flag.String("f", "etc/dev/im.yaml", "the config file")

//var configFile = flag.String("f", "apps/im/ws/etc/dev/im.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

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
