package repository

import (
	"github.com/dendik-creation/task-control/internal/domain"
	"gorm.io/gorm"
)

type columnRepository struct {
	db *gorm.DB
}

func NewColumnRepository(db *gorm.DB) domain.ColumnRepository {
	return &columnRepository{db}
}

func (r *columnRepository) Create(column *domain.Column) error {
	return r.db.Create(column).Error
}

func (r *columnRepository) FetchByUserID(userID uint) ([]domain.Column, error) {
	var columns []domain.Column
	err := r.db.Preload("Tasks").Where("user_id = ?", userID).Order("position asc").Find(&columns).Error
	if err != nil {
		return nil, err
	}
	return columns, nil
}

func (r *columnRepository) Update(column *domain.Column) error {
	return r.db.Save(column).Error
}

func (r *columnRepository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&domain.Column{}).Error
}
