package handler

import (
	"context"

	"github.com/nakanoyujin/go_to_do_app/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . ListTaskService AddTaskService
type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}
type ListTaskService interface {
	ListTasks(ctx context.Context) (entity.Tasks, error)
}
