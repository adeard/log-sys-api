package logging

import (
	"log-sys-api/domain"
	"log-sys-api/utils"
	"math"
	"strconv"
	"time"
)

type Service interface {
	Store(input domain.LogRequest) (domain.LogRequest, error)
	GetAll(input domain.LogFilterRequest) (domain.PagingResponse, error)
	GetDetail(input string) (domain.LogData, error)
	GetTopError(input domain.LogFilterRequest) ([]domain.LogTopErrorData, error)
	GetTotalByDate(input domain.LogFilterRequest) ([]domain.LogTotalData, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(logFilter domain.LogFilterRequest) (domain.PagingResponse, error) {
	var result domain.PagingResponse

	if logFilter.Limit == 0 {
		logFilter.Limit = 10
	}

	if logFilter.Page == 0 {
		logFilter.Page = 1
	}

	logs, err := s.repository.FindAll(logFilter)
	logsTotal, _ := s.repository.CountData(logFilter)

	for i, logData := range logs {

		logDateParse, _ := time.Parse("2006-01-02T15:04:05Z07:00", logData.CreatedAt)
		logDate := logDateParse.Format("2006-01-02 15:04:05")

		updateDateParse, _ := time.Parse("2006-01-02T15:04:05Z07:00", logData.UpdatedAt)
		logUpdateDate := updateDateParse.Format("2006-01-02 15:04:05")

		logs[i].CreatedAt = logDate
		logs[i].UpdatedAt = logUpdateDate
	}

	totalPage := float64(logsTotal) / float64(logFilter.Limit)

	result.Data = logs
	result.PerPage = logFilter.Limit
	result.Page = logFilter.Page
	result.TotalData = int(logsTotal)
	result.TotalPage = int(math.Ceil(totalPage))

	return result, err
}

func (s *service) GetDetail(input string) (domain.LogData, error) {
	var result domain.LogData

	logId, err := strconv.Atoi(input)
	if err != nil {
		return result, err
	}

	result, err = s.repository.FindById(logId)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *service) GetTopError(logFilter domain.LogFilterRequest) ([]domain.LogTopErrorData, error) {

	if logFilter.Limit == 0 {
		logFilter.Limit = 5
	}

	result, err := s.repository.GetTopError(logFilter)

	return result, err
}

func (s *service) GetTotalByDate(logFilter domain.LogFilterRequest) ([]domain.LogTotalData, error) {

	result, err := s.repository.CountByDate(logFilter)

	return result, err
}

func (s *service) Store(input domain.LogRequest) (domain.LogRequest, error) {

	if input.CreatedAt == "" {
		input.CreatedAt = utils.GetCurrentDateTime()
	}

	if input.UpdatedAt == "" {
		input.UpdatedAt = utils.GetCurrentDateTime()
	}

	// if input.StatusCode == 200 {
	// 	return domain.LogRequest{}, nil
	// }

	log, err := s.repository.Store(input)
	if err != nil {
		utils.LogInit("error", err.Error())

		return domain.LogRequest{}, err
	}

	return log, err
}
