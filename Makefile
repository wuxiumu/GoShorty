APP_NAME=goshorty
PKG=./cmd/main.go
PORT=8080
PPORT=6060

.PHONY: all run build clean docker docker-run lint test

all: build

## 👟 本地运行
run:
	go run $(PKG)

## 🧱 构建二进制
build:
	go build -o $(APP_NAME) $(PKG)

## 🧹 清理构建产物
clean:
	rm -f $(APP_NAME)

## 🐳 构建 Docker 镜像
docker:
	docker build -t $(APP_NAME):latest .

## 🐳 启动容器（含 pprometheus + pprof 暴露）
docker-run:
	docker run -p $(PORT):8080 -p $(PPORT):6060 $(APP_NAME):latest

## 🔍 静态代码检查
lint:
	golangci-lint run ./...

## ✅ 测试（可扩展 *_test.go 文件）
test:
	go test ./...