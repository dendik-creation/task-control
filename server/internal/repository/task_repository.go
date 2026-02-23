package repository

import (
	"github.com/dendik-creation/task-control/internal/domain"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) Create(task *domain.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) FetchByColumnID(columnID uint, userID uint) ([]domain.Task, error) {
	var tasks []domain.Task
	err := r.db.Where("column_id = ? AND user_id = ?", columnID, userID).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) Update(task *domain.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) MoveTask(taskID uint, newColumnID uint, userID uint) error {
	return r.db.Model(&domain.Task{}).
		Where("id = ? AND user_id = ?", taskID, userID).
		Update("column_id", newColumnID).Error
}

func (r *taskRepository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&domain.Task{}).Error
}
