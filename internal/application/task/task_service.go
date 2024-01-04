// Package task provides
package task

import (
	"context"

	"github.com/tingchima/gogolook/internal/domain"
)

// 列出任務
func (s *Service) ListTasks(ctx context.Context, param domain.TaskParam) ([]domain.Task, error) {

	return s.postgresRepo.ListTasks(ctx, param)
}

// 建立任務
func (s *Service) CreateTask(ctx context.Context, param domain.Task) (*domain.Task, error) {

	return s.postgresRepo.CreateTask(ctx, param)
}

// 修改任務
func (s *Service) UpdateTask(ctx context.Context, param domain.Task) (*domain.Task, error) {

	// update task
	// if task is not exist, should return not found error

	return s.postgresRepo.UpdateTask(ctx, param)
}

// 透過ID刪除任務
func (s *Service) DeleteTaskByID(ctx context.Context, id int64) error {

	// delete task by id
	// if task is not exist, should return not found error

	return s.postgresRepo.DeleteTaskByID(ctx, id)
}
