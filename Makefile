APP_NAME=goshorty
PKG=./cmd/main.go
PORT=8080

.PHONY: all run build clean docker docker-run compose lint test

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

## 🐳 启动容器。指标默认在容器内 127.0.0.1:6060，不映射到宿主机。
docker-run:
	docker run --rm -p $(PORT):8080 \
		-e HTTP_ADDR=:8080 \
		-e PUBLIC_BASE_URL=http://localhost:$(PORT) \
		$(APP_NAME):latest

## 🐳 应用与 Prometheus（6060 仅在 compose 网络内）
compose:
	docker compose up --build

## 🔍 静态代码检查
lint:
	golangci-lint run ./...

## ✅ 测试（可扩展 *_test.go 文件）
test:
	go test ./...