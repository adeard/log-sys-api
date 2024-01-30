package logging

import "gorm.io/gorm"

func LoggingRegistry(db *gorm.DB) Service {
	logRepository := NewRepository(db)
	logService := NewService(logRepository)

	return logService
}
