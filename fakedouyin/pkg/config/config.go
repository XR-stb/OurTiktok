package config

import (
	_ "fakedouyin/pkg/logger"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Server   Server
	Mysql    Mysql
	Redis    Redis
	User     User
	Favorite Favorite
}

type Server struct {
	IP   string
	Port string
}

type Mysql struct {
	Dsn string
}

type Redis struct {
	IP   string
	Port string
}

type User struct {
	IP   string
	Port string
}

type Favorite struct {
	IP   string
	Port string
}

var C Config

func init() {
	C = Config{}
	_, err := toml.DecodeFile("./config/config.toml", &C)
	if err != nil {
		logrus.Fatalln(err)
	}
}
