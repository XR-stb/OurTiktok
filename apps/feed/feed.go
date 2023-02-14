package main

import (
	"flag"
	"fmt"
	"time"

	"OutTiktok/apps/feed/feed"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"

	"OutTiktok/apps/feed/internal/config"
	"OutTiktok/apps/feed/internal/server"
	"OutTiktok/apps/feed/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/feed.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		feed.RegisterFeedServer(grpcServer, server.NewFeedServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	//将 rpc注册到consul
	if err := consul.RegisterService(c.ListenOn, c.Consul); err != nil {
		panic(err)
	}

	fmt.Printf("[%s] Starting rpc server at %s...\n", time.Now().Format("2006-01-02 15:04:05"), c.ListenOn)
	s.Start()
}
