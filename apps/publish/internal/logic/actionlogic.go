package logic

import (
	"OutTiktok/dao"
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v6"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/jpeg"
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

	// 视频链接
	playUrl := "http://116.62.164.8:9000/videos/" + filename + ".mp4"

	// 获取封面
	coverData, err := readFrameAsJpeg(playUrl)
	if err != nil {
		return nil, err
	}

	//上传封面
	coverReader := bytes.NewReader(coverData)
	if _, err := l.svcCtx.Minio.PutObject(l.svcCtx.Config.Minio.CoverBucket, filename+".jpg", coverReader, coverReader.Size(), minio.PutObjectOptions{ContentType: "image/jpeg"}); err != nil {
		return &publish.ActionRes{Status: -1}, nil
	}

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

	if ttl, _ := l.svcCtx.Redis.Ttl(key2); ttl > 0 {
		_, _ = l.svcCtx.Redis.Sadd(key2, video.Id) // 0占位
	}

	_, _ = l.svcCtx.Redis.Zadd("feed", video.UploadTime, strconv.FormatInt(video.Id, 10))

	return &publish.ActionRes{}, nil
}

// ReadFrameAsJpeg
// 从视频流中截取一帧并返回 需要在本地环境中安装ffmpeg并将bin添加到环境变量
func readFrameAsJpeg(filePath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), err
}
