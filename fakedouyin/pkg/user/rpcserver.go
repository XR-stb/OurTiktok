package user

import (
	"context"
	"fakedouyin/pkg/config"
	"fakedouyin/pkg/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

var s *rpcserver

type rpcserver struct {
	service.UnimplementedUserServiceServer
}

func init() {
	s = &rpcserver{}
	go s.serve()
}

func (server *rpcserver) serve() {
	l, err := net.Listen("tcp", config.C.User.Port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	service.RegisterUserServiceServer(s, server)
	s.Serve(l)
}

func (server *rpcserver) GetUserInfo(ctx context.Context, req *service.GetUserInfoReq) (*service.GetUserInfoRes, error) {
	var users []User
	db.Where("id IN ?", req.UsersId).Find(&users)
	usersinfo := make([]*service.UserInfo, 0, len(users))
	rep := &service.GetUserInfoRes{
		UsersInfo: usersinfo,
	}
	for _, v := range users {
		rep.UsersInfo = append(rep.UsersInfo, &service.UserInfo{
			UserId: int64(v.Id),
			Name:   v.Username,
		})
	}
	return rep, nil
}
