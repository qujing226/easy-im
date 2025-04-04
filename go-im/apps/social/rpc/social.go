package main

import (
	"flag"
	"fmt"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/internal/config"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/internal/server"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/internal/svc"
	"github.com/peninsula12/easy-im/go-im/pkg/configserver"
	"github.com/peninsula12/easy-im/go-im/pkg/interceptor"
	"github.com/peninsula12/easy-im/go-im/pkg/interceptor/rpcserver"

	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/dev/social.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	var configs = "social-rpc.yaml"
	err := configserver.NewConfigServer(*configFile, configserver.NewSail(&configserver.Config{
		ETCDEndpoints:  "118.178.120.11:3379",
		ProjectKey:     "3c46a0407be60a1f00731ab8e9575df2",
		Namespace:      "social",
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
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		social.RegisterSocialServer(grpcServer, server.NewSocialServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	s.AddUnaryInterceptors(rpcserver.LogInterceptor, rpcserver.SyncLimiterInterceptor(10))
	s.AddUnaryInterceptors(interceptor.NewIdempotenceServer(interceptor.NewDefaultIdempotent(c.Cache[0].RedisConf)))

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
