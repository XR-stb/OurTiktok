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

	"OutTiktok/apps/comment/comment"
	"OutTiktok/apps/comment/internal/config"
	"OutTiktok/apps/comment/internal/server"
	"OutTiktok/apps/comment/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/comment.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	initComment(ctx.Redis, ctx.DB)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		comment.RegisterCommentServer(grpcServer, server.NewCommentServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	//将 rpc注册到consul
	if err := consul.RegisterService(c.ListenOn, c.Consul); err != nil {
		panic(err)
	}

	fmt.Printf("[%s] Starting rpc server at %s...\n", time.Now().Format("2006-01-02 15:04:05"), c.ListenOn)
	s.Start()
}

// 将评论ID写入缓存
func initComment(Redis *redis.Redis, DB *gorm.DB) {
	videoIds, err := Redis.Zrange("feed", 0, -1)
	if err != nil {
		panic(err)
	}
	for _, id := range videoIds {
		var comments []*dao.Comment
		DB.Table("comments").Where("video_id = ?", id).Find(&comments)
		if len(comments) == 0 {
			continue
		}
		key := fmt.Sprintf("cids_%s", id)
		pairs := make([]redis.Pair, len(comments))
		for i, c := range comments {
			pairs[i].Score = c.CreateTime
			pairs[i].Key = strconv.FormatInt(c.Id, 10)
		}
		_, _ = Redis.Zadds(key, pairs...)
	}
}
