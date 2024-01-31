package logging

import (
	"log-sys-api/domain"
	"log-sys-api/utils"
)

type Service interface {
	Store(input domain.LogRequest) (domain.LogRequest, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
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