package model

import (
	"time"
)

// TaskStatus 任务状态枚举
type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "PENDING"
	TaskStatusRunning   TaskStatus = "RUNNING"
	TaskStatusCompleted TaskStatus = "COMPLETED"
	TaskStatusFailed    TaskStatus = "FAILED"
	TaskStatusCanceled  TaskStatus = "CANCELED"
)

// Task 任务模型
type Task struct {
	ID             string    `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Type           string    `json:"type"`
	Priority       int       `json:"priority"`
	Status         TaskStatus `json:"status"`
	Payload        string    `json:"payload"`
	Result         string    `json:"result"`
	Error          string    `json:"error"`
	MaxRetries     int       `json:"max_retries"`
	CurrentRetries int       `json:"current_retries"`
	Timeout        int       `json:"timeout"` // 秒
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	StartedAt      time.Time `json:"started_at"`
	CompletedAt    time.Time `json:"completed_at"`
}