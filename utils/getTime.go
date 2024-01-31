package utils

import "time"

func GetCurrentDateTime() string {
	location, _ := time.LoadLocation("Local")

	return time.Now().In(location).Format("2006-01-02 15:04:05")
}
