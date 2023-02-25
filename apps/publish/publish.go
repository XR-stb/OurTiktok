package main

import (
	"OutTiktok/dao"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"gorm.io/gorm"
	"strconv"
	"time"

	"OutTiktok/apps/publish/internal/config"
	"OutTiktok/apps/publish/internal/server"
	"OutTiktok/apps/publish/internal/svc"
	"OutTiktok/apps/publish/publish"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/publish.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	initFeed(ctx.Redis, ctx.DB)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		publish.RegisterPublishServer(grpcServer, server.NewPublishServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	//将 rpc注册到consul
	err := consul.RegisterService(c.ListenOn, c.Consul)
	if err != nil {
		panic(err)
	}
	defer s.Stop()

	fmt.Printf("[%s] Starting rpc server at %s...\n", time.Now().Format("2006-01-02 15:04:05"), c.ListenOn)
	s.Start()
}

// 将视频流缓存至Redis
func initFeed(redis *redis.Redis, DB *gorm.DB) {
	var videos []*dao.Video
	DB.Find(&videos)
	for _, v := range videos {
		_, err := redis.Zadd("feed", v.UploadTime, strconv.FormatInt(v.Id, 10))
		if err != nil {
			panic("Make sure Redis alive")
		}
	}
}
