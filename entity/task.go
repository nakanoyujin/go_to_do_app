package entity

import "time"

type TaskID int64

type TaskStatus string

const (
	TaskTodo       TaskStatus = "todo"
	TaskDoing      TaskStatus = "doing"
	TaskStatusDone TaskStatus = "done"
)

type Task struct {
	ID      TaskID     `json:"id"`
	Title   string     `json:"Title"`
	Status  TaskStatus `json:"Status"`
	Created time.Time  `json:"created"`
}

type Tasks []*Task
