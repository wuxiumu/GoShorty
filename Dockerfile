# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

WORKDIR /app

# 添加依赖管理文件
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# 拷贝源代码
COPY . .

# 构建可执行文件
RUN go build -o goshorty ./cmd/main.go

# 开放 web 和 pprof+metrics 端口
EXPOSE 8080 6060

# 启动服务
CMD ["/app/goshorty"]