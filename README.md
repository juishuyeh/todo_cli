# Todo CLI

[![Go Version](https://img.shields.io/github/go-mod/go-version/yourname/todo-cli)](https://golang.org/)
[![License](https://img.shields.io/github/license/yourname/todo-cli)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/juishuyeh/todo_cli)](https://goreportcard.com/report/github.com/juishuyeh/todo_cli)

一個簡潔優雅的終端機待辦事項管理工具，使用 Go 和 SQLite 構建。

## ✨ 功能特點

- 🚀 快速啟動，無需配置
- 💾 本地 SQLite 儲存
- 🎨 美觀的 TUI 介面
- ⌨️ 鍵盤快捷鍵操作
- 📦 單一二進制檔案，易於分發

## 📦 安裝

### 使用 Go Install（推薦）
```bash
go install github.com/juishuyeh/todo_cli/cmd/todo@latest
```

### 從原始碼編譯
```bash
git clone https://github.com/juishuyeh/todo_cli.git
cd todo_cli
make build
# 二進制檔案位於 dist/todo
```

## 🚀 使用方式

### 基本命令

```bash
# 啟動互動模式（推薦）
todo

# 新增任務
todo add 買牛奶
todo a 完成專案文檔

# 列出所有任務
todo list
todo ls
todo l

# 切換任務完成狀態
todo done 1
todo do 2
todo d 3

# 刪除任務
todo delete 1
todo del 2
todo rm 3

# 顯示幫助
todo help
```

### 互動模式

執行 `todo` 進入互動模式，享受更流暢的操作體驗：

```
╔═══════════════════════════════════════════════════════╗
║         📝 Todo CLI - 待辦事項管理工具               ║
╚═══════════════════════════════════════════════════════╝

> [1] ☐ 完成專案文檔
  [2] ✓ 學習 Go 語言
  [3] ☐ 買牛奶

─────────────────────────────────────────────────────────
[a] 新增  [d] 刪除  [Space] 完成/未完成  [j/k] 上下移動  [q] 退出
```

#### 快捷鍵
- `a` - 新增任務
- `d` - 刪除選中的任務
- `Space` - 切換任務完成狀態
- `j` 或 `↓` - 向下移動
- `k` 或 `↑` - 向上移動
- `q` - 退出

## 💾 資料儲存

任務資料儲存在 `~/.todo-cli/todo.db`，使用 SQLite 資料庫。

## 🛠️ 開發

```bash
# 執行測試
make test

# 編譯
make build

# 安裝到 GOPATH
make install

# 清理編譯產物
make clean

# 跨平台編譯
make build-all
```

## 📝 授權

本專案採用 MIT 授權條款 - 詳見 [LICENSE](LICENSE) 檔案
