# 基础镜像
FROM golang:1.25.0

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

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./admin-go"]
