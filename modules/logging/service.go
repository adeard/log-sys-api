package logging

import (
	"log-sys-api/domain"
	"log-sys-api/utils"
)

type Service interface {
	Store(input domain.LogRequest) (domain.LogData, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Store(input domain.LogRequest) (domain.LogData, error) {

	input.CreatedAt = utils.GetCurrentDateTime()
	input.UpdatedAt = utils.GetCurrentDateTime()

	log, err := s.repository.Store(input)
	if err != nil {
		return domain.LogData{}, err
	}

	return log, err
}
