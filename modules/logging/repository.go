package logging

import (
	"log-sys-api/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.LogRequest) (domain.LogData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Store(input domain.LogRequest) (domain.LogData, error) {

	log := domain.LogData{LogRequest: input}

	err := r.db.Create(&log).Error

	return log, err
}
