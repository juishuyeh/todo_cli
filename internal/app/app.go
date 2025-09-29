package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/juishuyeh/todo_cli/internal/database"
	"github.com/juishuyeh/todo_cli/pkg/models"
)

type App struct {
	db *database.DB
}

// New 建立應用實例
func New() (*App, error) {
	// 取得使用者主目錄
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	// 建立設定目錄
	configDir := filepath.Join(home, ".todo-cli")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, err
	}

	// 開啟資料庫
	dbPath := filepath.Join(configDir, "todo.db")
	db, err := database.New(dbPath)
	if err != nil {
		return nil, err
	}

	return &App{db: db}, nil
}

// AddTask 新增待辦事項
func (a *App) AddTask(title string) error {
	task := models.NewTask(title)
	return a.db.AddTask(task)
}

// ListTasks 列出所有待辦事項
func (a *App) ListTasks() ([]*models.Task, error) {
	return a.db.ListTasks()
}

// ToggleTask 切換完成狀態
func (a *App) ToggleTask(id int) error {
	tasks, err := a.db.ListTasks()
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if task.ID == id {
			if task.Done {
				task.MarkUndone()
			} else {
				task.MarkDone()
			}
			return a.db.UpdateTask(task)
		}
	}
	return fmt.Errorf("task not found: %d", id)
}

// DeleteTask 刪除待辦事項
func (a *App) DeleteTask(id int) error {
	return a.db.DeleteTask(id)
}

// Close 關閉應用
func (a *App) Close() error {
	return a.db.Close()
}
