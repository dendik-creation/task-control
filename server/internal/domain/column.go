package domain

import "time"

type Column struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserID   uint   `json:"user_id" gorm:"not null"`
	Name     string `json:"name" gorm:"not null"`
	Color    string `json:"color"`
	Position int    `json:"position"` // from left to right, starting at 0

	Tasks []Task `json:"tasks" gorm:"foreignKey:ColumnID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Kontrak Interface untuk Column
type ColumnRepository interface {
	Create(column *Column) error
	FetchByUserID(userID uint) ([]Column, error)
}
