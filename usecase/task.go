package usecase

import (
	"context"
	"time"

	"github.com/xorwise/golang-todo-api/domain"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
	timeout        time.Duration
}

func NewTaskUsecase(tr domain.TaskRepository, timeout time.Duration) domain.TaskUsecase {
	return &taskUsecase{
		taskRepository: tr,
		timeout:        timeout,
	}
}

func (tu *taskUsecase) Create(c context.Context, task *domain.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.timeout)
	defer cancel()
	return tu.taskRepository.Create(ctx, task)
}

func (tu *taskUsecase) Fetch(c context.Context, req domain.FetchTaskRequest) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.timeout)
	defer cancel()
	return tu.taskRepository.FetchByUserID(ctx, req.UserID, req.Offset, req.Limit)
}

func (tu *taskUsecase) GetByID(c context.Context, id uint) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.timeout)
	defer cancel()

	return tu.taskRepository.GetByID(ctx, id)
}

func (tu *taskUsecase) Update(c context.Context, task *domain.Task, req *domain.UpdateTaskRequest) error {
	ctx, cancel := context.WithTimeout(c, tu.timeout)
	defer cancel()

	if req.Title != "" {
		task.Title = req.Title
	}
	if req.Description != "" {
		task.Description = req.Description
	}
	if !req.Deadline.IsZero() {
		task.Deadline = req.Deadline
	}

	return tu.taskRepository.Update(ctx, task)
}

func (tu *taskUsecase) Delete(c context.Context, task *domain.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.timeout)
	defer cancel()

	return tu.taskRepository.Delete(ctx, task.ID)
}
