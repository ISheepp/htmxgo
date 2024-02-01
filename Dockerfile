# 基础镜像，基于golang的alpine镜像构建--编译阶段
FROM golang:alpine AS builder
# 作者
LABEL org.opencontainers.image.authors="isheep"
# 全局工作目录
WORKDIR /go/isheep
# 把运行Dockerfile文件的当前目录所有文件复制到目标目录
COPY . /go/isheep
# 环境变量
ENV API_PORT=8080 \
    HTML_PORT=8081 \
    GIN_MODE=release
RUN go mod tidy
# 编译，关闭CGO，防止编译后的文件有动态链接，而alpine镜像里有些c库没有，直接没有文件的错误
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build main.go


# 使用alpine这个轻量级镜像为基础镜像--运行阶段
FROM alpine AS runner
# 全局工作目录
WORKDIR /go/isheep
# 复制编译阶段编译出来的运行文件到目标目录
COPY --from=builder /go/isheep .

# docker run命令触发的真实命令(相当于直接运行编译后的可运行文件)
CMD ["./main"]
