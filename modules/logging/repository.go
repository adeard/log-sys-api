package logging

import (
	"log-sys-api/domain"
	"log-sys-api/utils"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.LogRequest) (domain.LogRequest, error)
	FindAll(input domain.LogFilterRequest) ([]domain.LogData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(input domain.LogFilterRequest) ([]domain.LogData, error) {
	var mob []domain.LogData

	q := r.db.Debug().Table("logs")

	if input.Request != "" {
		q = q.Where("request LIKE ?", "%"+input.Request+"%")
	}

	if input.Response != "" {
		q = q.Where("response LIKE ?", "%"+input.Response+"%")
	}

	if input.StatusCode > 0 {
		q = q.Where("status_code = ?", input.StatusCode)
	}

	if input.Source != "" {
		q = q.Where("source = ?", input.Source)
	}

	if input.StartDate != "" && input.EndDate != "" {
		q = q.Where("created_at between ? and ?", input.StartDate, input.EndDate)
	}

	err := q.Find(&mob).Error

	return mob, err
}

func (r *repository) Store(input domain.LogRequest) (domain.LogRequest, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		utils.LogInit(err.Error())
	}

	return input, err
}
