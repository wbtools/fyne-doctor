# Fyne Doctor 🩺

[![Go Version](https://img.shields.io/badge/Go-1.16%2B-blue.svg)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Fyne](https://img.shields.io/badge/Fyne-v2.x-orange.svg)](https://fyne.io/)

一个全面的 Fyne 开发环境诊断工具。Fyne Doctor 检查您系统中的所有必需依赖项，检测常见问题，并根据最新的 [Fyne GitHub Issues](https://github.com/fyne-io/fyne/issues) 和官方文档提供平台特定的安装指导。

> **注意**: Fyne 官方也提供了一个图形化的环境设置工具 [Fyne Setup](https://github.com/fyne-io/setup)。如果您更喜欢图形界面，可以安装并运行 `go install fyne.io/setup@latest` 然后执行 `setup` 命令。

## ✨ 功能特性

### 🔍 **全面的环境检测**
- **核心依赖项**: Go、Fyne CLI、C 编译器、构建工具、Fyne Setup
- **移动开发**: Android SDK/NDK、iOS 模拟器、Android Studio
- **Web 开发**: WebAssembly 支持、Node.js
- **性能检测**: GPU 加速、系统资源
- **兼容性**: Wayland 支持、显示服务器检测

### 🎯 **智能问题检测**
- **Wayland 兼容性**: 检测 Wayland 环境问题
- **移动开发**: 验证 Android/iOS 设置
- **性能问题**: 识别 GPU 加速问题
- **版本冲突**: 检查 Go 版本要求

### 📊 **分类输出**
- **核心依赖项**: Fyne 开发必需工具
- **移动开发**: Android 和 iOS 开发工具
- **Web 开发**: WebAssembly 和 Web 构建工具
- **性能**: 系统性能指标
- **兼容性**: 平台特定兼容性检查

### 🔗 **GitHub Issues 集成**
- 链接到相关的 GitHub Issues
- 显示最近的问题趋势
- 提供针对性解决方案

## 🚀 安装

### 先决条件
- Go 1.16 或更高版本
- Git

### 从源码安装
```bash
git clone https://github.com/wbtools/fyne-doctor.git
cd fyne-doctor
go install
```

### 通过 Go 安装
```bash
go install github.com/wbtools/fyne-doctor@latest
```

## 📖 使用方法

### 基本使用
```bash
fyne-doctor doctor
```

### 高级选项
```bash
# 详细输出和日志记录
fyne-doctor doctor --verbose

# 按特定类别过滤
fyne-doctor doctor --categories core,mobile

# 以 JSON 格式输出结果
fyne-doctor doctor --json

# 保存输出到文件
fyne-doctor doctor --output report.txt

# 设置命令执行的自定义超时时间
fyne-doctor doctor --timeout 30s
```

### 命令行选项
| 参数 | 简写 | 描述 | 默认值 |
|------|------|------|--------|
| `--verbose` | `-v` | 启用详细输出 | `false` |
| `--json` | `-j` | 以 JSON 格式输出结果 | `false` |
| `--categories` | `-c` | 按类别过滤 | `all` |
| `--timeout` | `-t` | 命令执行超时时间 | `10s` |
| `--output` | `-o` | 保存输出到文件 | `stdout` |

### 可用类别
- `core`: Fyne 开发必需工具
- `mobile`: Android 和 iOS 开发工具
- `web`: WebAssembly 和 Web 开发工具
- `performance`: 系统性能指标
- `compatibility`: 平台兼容性检查

## 📋 输出示例

```
          Fyne Doctor

# Fyne
Version | v2.4.0

# System
┌───────────────────────────┐
| OS           | darwin     |
| Architecture | arm64      |
| CPU          | Apple M1 @ 3200MHz |
| Memory       | 16.00 GB   |
└───────────────────────────┘

## Core Dependencies
┌─────────────────────────────┬────────────┬────────────┬─────────────┬─────────────────────────────────────┐
| Dependency                  | Optional   | Status     | Version     | Description                         |
├─────────────────────────────┼────────────┼────────────┼─────────────┼─────────────────────────────────────┤
| Go                          |            | Installed  | go1.21.0    | Go programming language (minimum... |
| Fyne CLI                    |            | Installed  | v2.4.0      | Fyne command line tools            |
| Xcode CLI Tools             |            | Installed  | /Library/...| Xcode command line tools            |
└─────────────────────────────┴────────────┴────────────┴─────────────┴─────────────────────────────────────┘

## Mobile Development
┌─────────────────────────────┬────────────┬────────────┬─────────────┬─────────────────────────────────────┐
| Dependency                  | Optional   | Status     | Version     | Description                         |
├─────────────────────────────┼────────────┼────────────┼─────────────┼─────────────────────────────────────┤
| Android SDK                 | *          | Missing    |             | Android SDK for mobile development |
| iOS Simulator               | *          | Installed  | Found       | iOS Simulator for iOS development  |
└─────────────────────────────┴────────────┴────────────┴─────────────┴─────────────────────────────────────┘

# Diagnosis
 SUCCESS  Your system is ready for Fyne development!

# Common Issues Check
✅ No common issues detected!

# GitHub Issues Summary
Based on recent Fyne GitHub Issues, common problems include:
• Mobile web builds: Paste functionality issues (#5916)
• Android: NewMultiLineEntry scrolling problems (#5915)
• Mobile: Grid container performance issues (#5914)
• Wayland: App crashes on some systems (#5908)
• X11: Display wake-up crashes (#5899)
• Windows: UI position issues after minimize/restore (#5898)

For more details, visit: https://github.com/fyne-io/fyne/issues
```

## 🛠️ 平台特定设置

### Windows
```bash
# 从 https://golang.org/dl/ 安装 Go
# 从 https://www.msys2.org/ 安装 MSYS2
# 打开 MSYS2 MinGW 64-bit 终端并运行:
pacman -Syu
pacman -S git mingw-w64-x86_64-toolchain
# 将 C:\msys64\mingw64\bin 添加到您的 PATH
# 安装 Fyne CLI:
go install fyne.io/fyne/v2/cmd/fyne@latest
# 安装 Fyne Setup (可选，图形化环境检查工具):
go install fyne.io/setup@latest
```

### macOS
```bash
# 从 https://golang.org/dl/ 安装 Go
# 从 Mac App Store 安装 Xcode
# 安装 Xcode CLI 工具:
xcode-select --install
# 安装 Fyne CLI:
go install fyne.io/fyne/v2/cmd/fyne@latest
# 安装 Fyne Setup (可选，图形化环境检查工具):
go install fyne.io/setup@latest
```

### Linux (Ubuntu/Debian)
```bash
sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev pkg-config
go install fyne.io/fyne/v2/cmd/fyne@latest
# 安装 Fyne Setup (可选，图形化环境检查工具):
go install fyne.io/setup@latest
```

### Linux (Fedora)
```bash
sudo dnf install golang gcc libXcursor-devel libXrandr-devel mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel
go install fyne.io/fyne/v2/cmd/fyne@latest
# 安装 Fyne Setup (可选，图形化环境检查工具):
go install fyne.io/setup@latest
```

### Linux (Arch)
```bash
sudo pacman -S go xorg-server-devel libxcursor libxrandr libxinerama libxi
go install fyne.io/fyne/v2/cmd/fyne@latest
# 安装 Fyne Setup (可选，图形化环境检查工具):
go install fyne.io/setup@latest
```

## 🔧 开发

### 项目结构
```
fyne-doctor/
├── main.go          # 主应用程序代码
├── go.mod           # Go 模块定义
├── go.sum           # Go 模块校验和
└── README.md        # 本文档
```

### 从源码构建
```bash
git clone https://github.com/wbtools/fyne-doctor.git
cd fyne-doctor
go build -o fyne-doctor main.go
```

### 运行测试
```bash
go test ./...
```

### 贡献
1. Fork 仓库
2. 创建功能分支
3. 进行更改
4. 如适用，添加测试
5. 提交 Pull Request

## 🐛 故障排除

### 常见问题

#### 命令超时
如果命令超时，请增加超时时间：
```bash
fyne-doctor doctor --timeout 30s
```

#### 缺少依赖项
按照输出中提供的平台特定安装说明进行操作。

#### Wayland 问题
如果您在 Wayland 上运行并遇到问题：
- 检查是否安装了 Wayland 支持
- 考虑切换到 X11 进行开发
- 关注 [GitHub Issue #5908](https://github.com/fyne-io/fyne/issues/5908)

#### 移动开发设置
Android 开发：
- 安装 Android Studio
- 设置 `ANDROID_HOME` 和 `ANDROID_NDK_HOME` 环境变量
- 确保安装了 NDK

iOS 开发（仅限 macOS）：
- 安装 Xcode
- 使用 Apple Developer 账户登录
- 安装 iOS 模拟器

## 📚 相关资源

- [Fyne 文档](https://developer.fyne.io/)
- [Fyne GitHub 仓库](https://github.com/fyne-io/fyne)
- [Fyne GitHub Issues](https://github.com/fyne-io/fyne/issues)
- [Fyne Setup (官方图形化工具)](https://github.com/fyne-io/setup)
- [Go 文档](https://golang.org/doc/)

## 📄 许可证

本项目采用 MIT 许可证 - 详情请参阅 [LICENSE](LICENSE) 文件。

## 🤝 贡献

欢迎贡献！请随时提交 Pull Request。对于重大更改，请先打开一个 issue 来讨论您想要更改的内容。

## 📞 支持

如果您遇到任何问题或有疑问：
1. 查看[故障排除部分](#-故障排除)
2. 搜索现有的 [GitHub Issues](https://github.com/wbtools/fyne-doctor/issues)
3. 创建包含详细信息的新 issue

## 🙏 致谢

- [Fyne 团队](https://fyne.io/) 提供的出色 GUI 工具包
- [Fyne Setup](https://github.com/fyne-io/setup) 官方环境设置工具的启发
- [Cobra](https://github.com/spf13/cobra) 提供的 CLI 框架
- [gopsutil](https://github.com/shirou/gopsutil) 提供的系统信息
- 所有帮助改进此工具的贡献者和用户

---

**为 Fyne 社区用 ❤️ 制作**
