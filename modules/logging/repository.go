package logging

import (
	"log-sys-api/domain"
	"log-sys-api/utils"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.LogRequest) (domain.LogRequest, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Store(input domain.LogRequest) (domain.LogRequest, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		utils.LogInit(err.Error())
	}

	return input, err
}
