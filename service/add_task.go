package service

import (
	"context"
	"fmt"

	"github.com/nakanoyujin/go_to_do_app/entity"
	"github.com/nakanoyujin/go_to_do_app/store"
)

type AddTask struct {
	DB   store.Execer
	Repo TaskAdder //Repo型をTaskAdderインターフェースに代入している
}

func (a *AddTask) AddTask(ctx context.Context, title string) (*entity.Task, error) {
	t := &entity.Task{
		Title:  title,
		Status: entity.TaskStatusTodo,
	}
	//ここでmoqのテストをするために
	err := a.Repo.AddTask(ctx, a.DB, t)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return t, nil
}
