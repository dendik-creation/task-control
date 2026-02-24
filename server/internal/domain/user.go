package domain

import "time"

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null" binding:"required"`
	Email    string `json:"email" gorm:"unique;not null" binding:"required,email"`
	Password string `json:"-" gorm:"not null"`

	Columns []Column `json:"columns" gorm:"foreignKey:UserID"`
	Tasks   []Task   `json:"tasks" gorm:"foreignKey:UserID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
	GetByID(id uint) (*User, error)
}

type UserUsecase interface {
	Register(user *User) error
	Login(email string, password string) (string, error)
}
