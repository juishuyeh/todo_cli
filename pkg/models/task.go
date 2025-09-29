package models

import "time"

// Task 代表一個待辦事項
type Task struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Done      bool      `json:"done"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// NewTask 建立一個新的待辦事項
func NewTask(title string) *Task {
    now := time.Now()
    return &Task{
        Title:     title,
        Done:      false,
        CreatedAt: now,
        UpdatedAt: now,
    }
}

// MarkDone 標記為完成
func (t *Task) MarkDone() {
    t.Done = true
    t.UpdatedAt = time.Now()
}

// MarkUndone 標記為未完成
func (t *Task) MarkUndone() {
    t.Done = false
    t.UpdatedAt = time.Now()
}
