# 基础镜像
FROM golang:1.25.0 AS build

# 改变工作目录到 `app` 目录
WORKDIR /app

# 复制 go mod 文件
COPY go.mod go.sum ./

# 下载 Go 模块
RUN go mod download

# 复制源代码
COPY . ./

# 编译应用
RUN CGO_ENABLED=0 GOOS=linux go build -o admin-go .

# 基础镜像，使用 gcr.io/distroless/base-debian11 镜像更小，但缺少 /bin/sh 导致无法进入镜像内部查看内容
FROM debian:12.13 AS release

# 改变工作目录到 `app` 目录
WORKDIR /app

# 复制编译好的应用
COPY --from=build /app/admin-go /app/admin-go

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./admin-go"]
