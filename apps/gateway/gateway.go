package main

import (
	"OutTiktok/apps/gateway/internal/config"
	"OutTiktok/apps/gateway/internal/handler"
	"OutTiktok/apps/gateway/internal/svc"
	"flag"
	"fmt"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/gateway-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.RestConf.MaxBytes = 67108864
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
