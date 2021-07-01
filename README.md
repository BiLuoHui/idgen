# idgen

基于 Twitter Snowflake 、gRPC 开发的全局唯一ID生成服务，支持使用Docker容器化部署，同时优化了Docker镜像，减少镜像文件大小。

## 常规部署

```bash
$ ./server MACHINE_ID PORT
```

如：

```bash
$ ./server 1 8080
```
