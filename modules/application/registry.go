package application

import "gorm.io/gorm"

func ApplicationRegistry(db *gorm.DB) Service {
	applicationRepository := NewRepository(db)
	applicationService := NewService(applicationRepository)

	return applicationService
}
