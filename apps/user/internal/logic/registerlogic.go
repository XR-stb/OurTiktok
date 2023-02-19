package logic

import (
	"OutTiktok/dao"
	"context"
	"crypto/md5"
	"fmt"
	"math/rand"

	"OutTiktok/apps/user/internal/svc"
	"OutTiktok/apps/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

const avatarNum = 12
const backgroundImageNum = 12
const signatureNum = len(signatures)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterRes, error) {
	username := in.Username
	password := in.Password

	// 将用户写入数据库
	u := dao.User{
		Username:        username,
		Password:        fmt.Sprintf("%x", md5.Sum([]byte(password))), // md5加密
		Avatar:          fmt.Sprintf("http://%s/avatar/%d", l.svcCtx.Config.Minio.Host, rand.Intn(avatarNum)),
		BackgroundImage: fmt.Sprintf("http://%s/backgroundimage/%d", l.svcCtx.Config.Minio.Host, rand.Intn(backgroundImageNum)),
		Signature:       signatures[rand.Intn(signatureNum)],
	}
	if err := l.svcCtx.DB.Create(&u).Error; err != nil {
		return &user.RegisterRes{Status: -1}, nil
	}

	// 写入缓存
	key := fmt.Sprintf("uinfo_%d", u.Id)
	val := fmt.Sprintf("%s_%s_%s_%s", u.Username, u.Avatar, u.BackgroundImage, u.Signature)
	_ = l.svcCtx.Redis.Setex(key, val, 86400)

	return &user.RegisterRes{
		UserId: u.Id,
	}, nil
}

var signatures = [...]string{
	"从前从前有个人爱你很久，但偏偏风渐渐把距离吹得好远",
	"让爱渗透了地面，我要的只是你在我身边",
	"繁华如三千东流水，我只取一瓢爱了解",
	"心里的雨倾盆的下，也沾不湿她的发",
	"我一路向北，离开有你的季节，你说你好累，已无法再爱上谁",
	"能不能给我一首歌的时间，紧紧的把那拥抱变成永远",
	"最美的不是下雨天，是曾与你躲过雨的屋檐",
	"终有一天，我有属于我的天",
	"我想就这样牵着你的手不放开，爱能不能够永远单纯没有悲哀",
	"海鸟和鱼相爱，只是一场意外",
	"我顶着大太阳，只想为你撑伞",
	"而我已经分不清，你是友情，还是错过的爱情",
	"远方传来风笛，我只在意有你的消息，城堡为爱守着秘密，而我为你守着回忆",
	"你的脸没有化妆，我却疯狂爱上",
}
