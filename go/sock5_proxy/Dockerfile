# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.18 as builder

# 设置工作目录
WORKDIR /app

# 复制当前目录的内容到工作目录
COPY . .

# 编译 Go 应用
RUN go build -o myapp .

# 使用 Alpine 作为运行阶段的基础镜像
FROM alpine:latest

# 设置时区环境变量
ENV TZ=Asia/Shanghai

# 安装 tzdata 包以支持时区
ARG INSTALL_EXTRA_PACKAGES=false

RUN apk update && apk add --no-cache tzdata && \
    if [ "$INSTALL_EXTRA_PACKAGES" = "true" ]; then \
        apk add --no-cache some-extra-package; \
    fi

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译后的 Go 应用
COPY --from=builder /app/myapp /app/myapp

# 运行 Go 应用
CMD ["./myapp"]