package domain

import "time"

type Task struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"user_id" gorm:"not null"`
	ColumnID    uint   `json:"column_id" gorm:"not null"`
	Title       string `json:"title" gorm:"not null" binding:"required"`
	Description string `json:"description" binding:"required"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskRepository interface {
	Create(task *Task) error
	FetchByColumnID(columnID uint, userID uint) ([]Task, error)
	MoveTask(taskID uint, newColumnID uint, userID uint) error
	Update(task *Task) error
	Delete(id uint, userID uint) error
}
