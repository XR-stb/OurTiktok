package main

import (
	"OutTiktok/apps/relation/internal/config"
	"OutTiktok/apps/relation/internal/server"
	"OutTiktok/apps/relation/internal/svc"
	"OutTiktok/apps/relation/relation"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"time"
)

var configFile = flag.String("f", "etc/relation.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		relation.RegisterRelationServer(grpcServer, server.NewRelationServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	_ = consul.RegisterService(c.ListenOn, c.Consul)

	fmt.Printf("[%s] Starting rpc server at %s...\n", time.Now().Format("2006-01-02 15:04:05"), c.ListenOn)
	s.Start()
}
