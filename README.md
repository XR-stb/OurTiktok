## 字节跳动青训营项目——极简版抖音

第五届字节跳动青训营大项目，实现极简版抖音后台

## 技术栈

该项目使用Go语言进行开发，Go-zero作为微服务框架

1. 服务注册与发现：Consul
2. 服务网关：Go-zero生成
3. 服务调用：gRPC
4. 数据库交互：Gorm
5. 服务监控：Prometheus
6. 链路追踪：Jaeger
7. 配置中心：Nacos
8. 对象存储：Minio

## 项目功能

基础功能：视频Feed流、视频投稿、用户中心

互动功能：点赞、评论

社交功能：关注、收发消息

## 整体架构

![架构图](https://github.com/XR-stb/OurTiktok/blob/main/docs/架构图.jpg)

项目整体可分为八个模块：

1. 网关Gateway：验证请求、鉴权、缓存、数据转换、服务调用、负载均衡
2. 视频流Feed：获取视频流
3. 发布Publish：视频投稿、发布列表
4. 用户User：用户注册、登陆、信息
5. 点赞Favorite：点赞、取消点赞、点赞列表
6. 评论Comment：发布评论，删除评论
7. 关注Follow：关注用户、取消关注、关注列表
8. 消息Message：发送消息、接收消息

## 部署项目-Docker

1. 拉取项目到本地

```shell
git clone https://github.com/XR-stb/OurTiktok.git
```

2. 配置apps/publish/etc/publish.yaml中Minio--Expose为本机IP

3. 编译所有项目代码

```shell
make
```

4. 确保已经安装Docker、Docker-Compose，在项目根目录运行

```shell
docker-compose up
```

5. 项目会自动部署并运行，访问`127.0.0.1:3000`，点击上方的status-targets，可以看到服务的上线情况

![架构图](https://github.com/XR-stb/OurTiktok/blob/main/docs/服务上线.png)

## TODO
1. 使用Nacos作为配置中心，简化配置步骤
2. 使用Prometheus作为服务监控✅
3. 使用Jaeger作为链路追踪
4. 开发后台管理系统（待定）
5. Docker部署✅