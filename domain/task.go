package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string
	Description string
	Completed   bool
	Deadline    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      uint
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID uint, offset int, limit int) ([]Task, error)
	GetByID(c context.Context, id uint) (Task, error)
	Update(c context.Context, task *Task) error
	Delete(c context.Context, id uint) error
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	Fetch(c context.Context, req FetchTaskRequest) ([]Task, error)
	GetByID(c context.Context, id uint) (Task, error)
	Update(c context.Context, task *Task, req *UpdateTaskRequest) error
}
