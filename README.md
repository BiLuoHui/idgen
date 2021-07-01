# idgen

基于 Twitter Snowflake 、gRPC 开发的全局唯一ID生成服务，支持使用Docker容器化部署，同时优化了Docker镜像，减少镜像文件大小。

## 常规部署

```bash
$ cd cmd/server
$ go build -o server
$ ./server MACHINE_ID PORT
```

如：

```bash
$ ./server 1 8080
```

验证

```bash
$ cd cmd/client
$ go build -o client
$ ./clent 127.0.0.1:8080
```

## Docker 部署

构建镜像
```bash
$ sudo docker build -t idgen:v1.0 .
```

查看构建好的本地镜像
```bash
$ sudo docker image ls | grep idgen
```

运行容器
```bash
$ sudo docker run idgen:v1.0
```