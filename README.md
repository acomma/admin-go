## Docker 指南

参考 [Build your Go image](https://docs.docker.com/guides/golang/build-images/) 创建 Dockerfile 文件。

构建 Docker 镜像
```bash
docker build -t admin-go .
```

查看 Docker 镜像内容
```bash
docker run -it --rm --entrypoint /bin/sh admin-go:latest
```

启动 Docker 容器
- Windows (PowerShell):
```bash
docker run -d --name admin-go -p 8080:8080 `
    -e "MYSQL_ADDR=host.docker.internal:3306" `
    -e "MYSQL_USER=root" `
    -e "MYSQL_PASSWD=123456" `
    -e "MYSQL_DBNAME=admin-go" `
    admin-go:latest
```
- Windows (Command Prompt):
```bash
docker run -d --name admin-go -p 8080:8080 ^
    -e "MYSQL_ADDR=host.docker.internal:3306" ^
    -e "MYSQL_USER=root" ^
    -e "MYSQL_PASSWD=123456" ^
    -e "MYSQL_DBNAME=admin-go" ^
    admin-go:latest
```
