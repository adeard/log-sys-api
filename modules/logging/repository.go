package logging

import (
	"log-sys-api/domain"
	"log-sys-api/utils"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.LogRequest) (domain.LogRequest, error)
	FindAll(input domain.LogFilterRequest) ([]domain.LogData, error)
	FindById(input int) (domain.LogData, error)
	CountData(input domain.LogFilterRequest) (int64, error)
	CountByDate(input domain.LogFilterRequest) ([]domain.LogTotalData, error)
	GetTopError(input domain.LogFilterRequest) ([]domain.LogTopErrorData, error)
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

	if input.OrderBy != "" {
		sort := "asc"
		order := strings.ToUpper(input.OrderBy)

		if input.SortBy != "" {
			sort = input.SortBy
		}

		q = q.Order(order + " " + sort)
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

func (r *repository) FindById(input int) (domain.LogData, error) {
	var log domain.LogData

	err := r.db.Debug().Table("logs").
		Where("id = ? ", input).
		First(&log).
		Error

	if err != nil {
		return domain.LogData{}, err
	}

	return log, err
}

func (r *repository) Store(input domain.LogRequest) (domain.LogRequest, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		utils.LogInit("error", err.Error())
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
		q = q.Where("source LIKE ?", "%"+input.Source+"%")
	}

	if input.StartDate != "" && input.EndDate != "" {
		q = q.Where("created_at between ? and ?", input.StartDate, input.EndDate)
	}

	q.Count(&logTotal)

	return logTotal, nil
}

func (r *repository) CountByDate(input domain.LogFilterRequest) ([]domain.LogTotalData, error) {
	var logTotalData []domain.LogTotalData

	q := r.db.Table("logs").
		Select(
			`CAST(created_at AS DATE) as log_date`,
			`SUM(1) as log_total`,
		)

	if input.Request != "" {
		q = q.Where("request LIKE %" + input.Request + "%")
	}
	if input.Response != "" {
		q = q.Where("response LIKE %" + input.Response + "%")
	}
	if input.StatusCode > 0 {
		q = q.Where("status_code = " + string(input.StatusCode))
	}
	if input.Source != "" {
		q = q.Where("source LIKE %" + input.Source + "%")
	}
	if input.StartDate != "" && input.EndDate != "" {
		q = q.Where(`created_at between '` + input.StartDate + ` 00:00:01' and '` + input.EndDate + ` 23:59:59'`)
	}

	err := q.Group("CAST(created_at AS DATE)").
		Order("CAST(created_at AS DATE) ASC").
		Find(&logTotalData).
		Error

	if err != nil {
		return []domain.LogTotalData{}, err
	}

	return logTotalData, nil
}

func (r *repository) GetTopError(input domain.LogFilterRequest) ([]domain.LogTopErrorData, error) {
	var logTopErrorData []domain.LogTopErrorData

	q := `
		SELECT TOP (` + strconv.Itoa(input.Limit) + `)
			b.log_total, 
			b.source, 
			a.response,
			b.last_created_at,
			b.last_id           
        FROM
            logs a 
			inner join (
				SELECT 
				   count(a.id) log_total,
				   max(a.id) last_id,
				   a.source,
				  MAX(a.created_at) last_created_at
			   FROM logs a
			   group by a.source
			) b on b.last_id=a.id `

	if input.StartDate != "" && input.EndDate != "" {
		q = q + `WHERE last_created_at between '` + input.StartDate + ` 00:00:01' and '` + input.EndDate + ` 23:59:59'`
	}

	if input.Source != "" {
		q = q + `WHERE source like '%` + input.Source + `%'`
	}

	r.db.Raw(q + ` 
	ORDER BY 
		log_total DESC
	`).Scan(&logTopErrorData)

	return logTopErrorData, nil
}
