package application

import (
	"log-sys-api/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.ApplicationRequest) (domain.ApplicationData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Store(input domain.ApplicationRequest) (domain.ApplicationData, error) {

	application := domain.ApplicationData{ApplicationRequest: input}

	err := r.db.Create(&application).Error

	return application, err
}
