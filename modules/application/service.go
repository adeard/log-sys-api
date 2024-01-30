package application

import (
	"log-sys-api/domain"
	"log-sys-api/utils"
)

type Service interface {
	Store(input domain.ApplicationRequest) (domain.ApplicationData, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Store(input domain.ApplicationRequest) (domain.ApplicationData, error) {

	input.CreatedAt = utils.GetCurrentDateTime()
	input.UpdatedAt = utils.GetCurrentDateTime()

	application, err := s.repository.Store(input)
	if err != nil {
		return domain.ApplicationData{}, err
	}

	return application, err
}
