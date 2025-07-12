# 第一阶段：构建阶段
FROM golang:1.24 AS builder

WORKDIR /app

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

# 优先复制 mod 文件以提高构建缓存命中率
COPY go.mod go.sum ./
RUN go mod download

# 再复制全部源码（分开是为了缓存优化）
COPY . .

# 构建 Go 应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o devhub-server \
    main.go models.go handlers.go

# 第二阶段：生产镜像（最小化 + 支持时区）
FROM alpine:latest

# 安装系统 CA + 时区数据
RUN apk --no-cache add ca-certificates tzdata

# 设置默认时区为 Asia/Shanghai（也方便日志）
ENV TZ=Asia/Shanghai

# 设置工作目录
WORKDIR /root/

# 拷贝构建好的二进制
COPY --from=builder /app/devhub-server .

# 暴露端口
EXPOSE 8080

# 启动应用
CMD ["./devhub-server"]
