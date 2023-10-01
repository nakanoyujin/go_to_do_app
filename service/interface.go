package service

import (
	"context"

	"github.com/nakanoyujin/go_to_do_app/entity"
	"github.com/nakanoyujin/go_to_do_app/store"
)

// handler -> storeの参照をしたくない handler -> service interface -> store
//
//go:generate go run github.com/matryer/moq -out moq_test.go . TaskAdder TaskLister
type TaskAdder interface {
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}
type TaskLister interface {
	ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}

type UserRegister interface {
	RegisterUser(ctx context.Context, db store.Execer, u *entity.User) error
}
