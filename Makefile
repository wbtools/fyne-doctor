# Fyne Doctor Makefile

.PHONY: build test clean install lint help

# 变量
BINARY_NAME=fyne-doctor
VERSION=$(shell git describe --tags --always --dirty)
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

# 默认目标
all: build

# 构建二进制文件
build:
	@echo "正在构建 $(BINARY_NAME)..."
	go build $(LDFLAGS) -o $(BINARY_NAME) main.go

# 为多个平台构建
build-all:
	@echo "正在为多个平台构建..."
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY_NAME)-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY_NAME)-windows-amd64.exe main.go
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY_NAME)-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BINARY_NAME)-darwin-arm64 main.go

# 运行测试
test:
	@echo "正在运行测试..."
	go test -v ./...

# 运行带覆盖率的测试
test-coverage:
	@echo "正在运行带覆盖率的测试..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# 安装二进制文件
install:
	@echo "正在安装 $(BINARY_NAME)..."
	go install $(LDFLAGS) .

# 运行代码检查
lint:
	@echo "正在运行代码检查..."
	golangci-lint run

# 清理构建产物
clean:
	@echo "正在清理..."
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME)-*
	rm -f coverage.out coverage.html

# 运行 doctor 命令
doctor:
	@echo "正在运行 fyne-doctor..."
	go run main.go doctor

# 使用详细输出运行
doctor-verbose:
	@echo "正在使用详细输出运行 fyne-doctor..."
	go run main.go doctor --verbose

# 使用 JSON 输出运行
doctor-json:
	@echo "正在使用 JSON 输出运行 fyne-doctor..."
	go run main.go doctor --json

# 仅运行核心依赖项检查
doctor-core:
	@echo "正在运行 fyne-doctor 仅检查核心依赖项..."
	go run main.go doctor --categories core

# 显示帮助
help:
	@echo "可用目标:"
	@echo "  build          - 构建二进制文件"
	@echo "  build-all      - 为多个平台构建"
	@echo "  test           - 运行测试"
	@echo "  test-coverage  - 运行带覆盖率的测试"
	@echo "  install        - 安装二进制文件"
	@echo "  lint           - 运行代码检查"
	@echo "  clean          - 清理构建产物"
	@echo "  doctor         - 运行 fyne-doctor"
	@echo "  doctor-verbose - 使用详细输出运行"
	@echo "  doctor-json    - 使用 JSON 输出运行"
	@echo "  doctor-core    - 仅检查核心依赖项"
	@echo "  help           - 显示此帮助"
