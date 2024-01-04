// Package task provides
package task

import (
	"context"

	"github.com/tingchima/gogolook/internal/domain"
)

// Repository
//
//go:generate mockgen -destination mocks/repository.go -package=mocks . Repository
type Repository interface {
	TaskRepository

	// maybe other repositories
}

// TaskRepository .
type TaskRepository interface {
	// 列出任務
	ListTasks(ctx context.Context, param domain.TaskParam) ([]domain.Task, error)
	// 建立任務
	CreateTask(ctx context.Context, param domain.Task) (*domain.Task, error)
	// 修改任務
	UpdateTask(ctx context.Context, param domain.Task) (*domain.Task, error)
	// 透過ID刪除任務
	DeleteTaskByID(ctx context.Context, id int64) error
}
