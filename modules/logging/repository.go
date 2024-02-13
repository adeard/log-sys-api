package logging

import (
	"log-sys-api/domain"
	"log-sys-api/utils"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.LogRequest) (domain.LogRequest, error)
	FindAll(input domain.LogFilterRequest) ([]domain.LogData, error)
	CountData(input domain.LogFilterRequest) (int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(input domain.LogFilterRequest) ([]domain.LogData, error) {
	var logs []domain.LogData

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

	offset := (input.Limit * (input.Page - 1))

	err := q.
		Limit(input.Limit).
		Offset(offset).
		Find(&logs).
		Error
	if err != nil {
		return nil, err
	}

	return logs, err
}

func (r *repository) Store(input domain.LogRequest) (domain.LogRequest, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		utils.LogInit(err.Error())
	}

	return input, err
}

func (r *repository) CountData(input domain.LogFilterRequest) (int64, error) {
	var logTotal int64

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

	q.Count(&logTotal)

	return logTotal, nil
}
