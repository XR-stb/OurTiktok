package logic

import (
	"OutTiktok/dao"
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v6"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"os"
	"strconv"
	"time"

	"OutTiktok/apps/publish/internal/svc"
	"OutTiktok/apps/publish/publish"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionLogic {
	return &ActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActionLogic) Action(in *publish.ActionReq) (*publish.ActionRes, error) {
	filename := l.svcCtx.Sf.New()

	// 上传视频
	reader := bytes.NewReader(in.Data)
	if _, err := l.svcCtx.Minio.PutObject(l.svcCtx.Config.Minio.VideoBucket, filename+".mp4", reader, reader.Size(), minio.PutObjectOptions{ContentType: "video/mp4"}); err != nil {
		return &publish.ActionRes{Status: -1}, nil
	}

	// 上传缩略图
	//reader2 := readFrameAsJpeg(l.svcCtx.Config.Minio.Host+"/videos/"+filename, 1)
	//if _, err := l.svcCtx.Minio.PutObject(l.svcCtx.Config.Minio.CoverBucket, filename+".jpg", reader2, reader.Size(), minio.PutObjectOptions{ContentType: "image/jpeg"}); err != nil {
	//	return &publish.ActionRes{Status: -1}, nil
	//}

	// 写入数据库
	video := dao.Video{
		AuthorId:   in.UserId,
		UploadTime: time.Now().UnixMilli(),
		PlayUrl:    "http://" + l.svcCtx.Config.Minio.Host + "/videos/" + filename + ".mp4",
		CoverUrl:   "http://" + l.svcCtx.Config.Minio.Host + "/covers/" + filename + ".jpg",
		Title:      in.Title,
	}
	if l.svcCtx.DB.Create(&video).Error != nil {
		return &publish.ActionRes{Status: -1}, nil
	}

	// 写入缓存
	key := fmt.Sprintf("vinfo_%d", video.Id)
	key2 := fmt.Sprintf("uv_%d", in.UserId)
	val := fmt.Sprintf("%d_%s_%s_%s", in.UserId, video.PlayUrl, video.CoverUrl, in.Title)
	_ = l.svcCtx.Redis.Setex(key, val, 86400)
	_, _ = l.svcCtx.Redis.Sadd(key2, 0, video.Id) // 0占位
	_ = l.svcCtx.Redis.Expire(key2, 86400)
	_, _ = l.svcCtx.Redis.Zadd("feed", video.UploadTime, strconv.FormatInt(video.Id, 10))

	return &publish.ActionRes{}, nil
}

func readFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}
