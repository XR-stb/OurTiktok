package publish

import (
	"fakedouyin/pkg/config"
	"fakedouyin/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var userclient service.UserServiceClient
var favoriteclient service.FavoriteServiceClient

func init() {
	conn, err := grpc.Dial(config.C.User.IP+config.C.User.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	userclient = service.NewUserServiceClient(conn)

	conn, err = grpc.Dial(config.C.Favorite.IP+config.C.Favorite.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	favoriteclient = service.NewFavoriteServiceClient(conn)
}
