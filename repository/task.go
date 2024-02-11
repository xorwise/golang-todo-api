package repository

import (
	"context"

	"github.com/xorwise/golang-todo-api/domain"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &taskRepository{db: db}
}

func (tr *taskRepository) Create(c context.Context, task *domain.Task) error {
	return tr.db.WithContext(c).Create(task).Error
}

func (tr *taskRepository) FetchByUserID(c context.Context, userID uint) ([]domain.Task, error) {
	var tasks []domain.Task
	if err := tr.db.WithContext(c).Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (tr *taskRepository) GetByID(c context.Context, id uint) (domain.Task, error) {
	var task domain.Task
	if err := tr.db.WithContext(c).Where("id = ?", id).First(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}

func (tr *taskRepository) Update(c context.Context, task *domain.Task) error {
	return tr.db.WithContext(c).Save(task).Error
}

func (tr *taskRepository) Delete(c context.Context, id uint) error {
	return tr.db.WithContext(c).Where("id = ?", id).Delete(&domain.Task{}).Error
}
