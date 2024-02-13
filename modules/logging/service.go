package logging

import (
	"log-sys-api/domain"
	"log-sys-api/utils"
	"math"
)

type Service interface {
	Store(input domain.LogRequest) (domain.LogRequest, error)
	GetAll(input domain.LogFilterRequest) (domain.PagingResponse, error)
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

	totalPage := float64(logsTotal) / float64(logFilter.Limit)

	result.Data = logs
	result.PerPage = logFilter.Limit
	result.Page = logFilter.Page
	result.TotalData = int(logsTotal)
	result.TotalPage = int(math.Ceil(totalPage))

	return result, err
}

func (s *service) Store(input domain.LogRequest) (domain.LogRequest, error) {

	input.CreatedAt = utils.GetCurrentDateTime()
	input.UpdatedAt = utils.GetCurrentDateTime()

	log, err := s.repository.Store(input)
	if err != nil {
		utils.LogInit(err.Error())

		return domain.LogRequest{}, err
	}

	return log, err
}
