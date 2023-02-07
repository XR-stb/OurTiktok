package favorite

import (
	"context"
	"fakedouyin/pkg/config"
	"fakedouyin/pkg/service"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"google.golang.org/grpc"
	"log"
	"net"
)

var s *rpcserver

type rpcserver struct {
	service.UnimplementedFavoriteServiceServer
}

func init() {
	s = &rpcserver{}
	go s.serve()
}

func (server *rpcserver) serve() {
	l, err := net.Listen("tcp", config.C.Favorite.Port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	service.RegisterFavoriteServiceServer(s, server)
	s.Serve(l)
}

func (server *rpcserver) GetFavorite(ctx context.Context, req *service.GetFavoriteReq) (*service.GetFavoriteRes, error) {
	userId := req.UsersId
	rep := &service.GetFavoriteRes{
		Favorites: make([]*service.VideoFavorite, 0, len(req.VideosId)),
	}
	for _, vid := range req.VideosId {
		// 查缓存
		key := fmt.Sprintf("f_%x", vid)
		favorite := Favorite{}
		count, err := redis.Int64(r.Do("get", key))
		if err != nil {
			// 查数据库
			db.Model(&favorite).Where("video_id=?", vid).Count(&count)
			r.Do("setex", key, 86400, count)
		}
		var isFavorite int64
		if userId != 0 {
			db.Model(favorite).Where("video_id=? AND user_id=? AND status=?", vid, userId, 1).Count(&isFavorite)
		}
		rep.Favorites = append(rep.Favorites, &service.VideoFavorite{
			VideoId:    vid,
			Count:      count,
			IsFavorite: isfavorite(isFavorite),
		})
	}
	return rep, nil
}

func isfavorite(i int64) bool {
	if i == 1 {
		return true
	}
	return false
}
