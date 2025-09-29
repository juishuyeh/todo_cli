package database

import (
    "database/sql"
    "time"

    _ "github.com/mattn/go-sqlite3"
    "github.com/juishuyeh/todo-cli/pkg/models"
)

type DB struct {
    conn *sql.DB
}

// New 建立資料庫連線
func New(dbPath string) (*DB, error) {
    conn, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, err
    }

    db := &DB{conn: conn}
    if err := db.createTable(); err != nil {
        return nil, err
    }

    return db, nil
}

// createTable 建立資料表
func (db *DB) createTable() error {
    query := `
    CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        done BOOLEAN NOT NULL DEFAULT 0,
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL
    );
    `
    _, err := db.conn.Exec(query)
    return err
}

// AddTask 新增待辦事項
func (db *DB) AddTask(task *models.Task) error {
    query := `INSERT INTO tasks (title, done, created_at, updated_at) 
              VALUES (?, ?, ?, ?)`
    result, err := db.conn.Exec(query, task.Title, task.Done, task.CreatedAt, task.UpdatedAt)
    if err != nil {
        return err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    task.ID = int(id)
    return nil
}

// ListTasks 列出所有待辦事項
func (db *DB) ListTasks() ([]*models.Task, error) {
    query := `SELECT id, title, done, created_at, updated_at FROM tasks ORDER BY id`
    rows, err := db.conn.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []*models.Task
    for rows.Next() {
        task := &models.Task{}
        err := rows.Scan(&task.ID, &task.Title, &task.Done, &task.CreatedAt, &task.UpdatedAt)
        if err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }
    return tasks, nil
}

// UpdateTask 更新待辦事項
func (db *DB) UpdateTask(task *models.Task) error {
    task.UpdatedAt = time.Now()
    query := `UPDATE tasks SET title = ?, done = ?, updated_at = ? WHERE id = ?`
    _, err := db.conn.Exec(query, task.Title, task.Done, task.UpdatedAt, task.ID)
    return err
}

// DeleteTask 刪除待辦事項
func (db *DB) DeleteTask(id int) error {
    query := `DELETE FROM tasks WHERE id = ?`
    _, err := db.conn.Exec(query, id)
    return err
}

// Close 關閉資料庫連線
func (db *DB) Close() error {
    return db.conn.Close()
}
