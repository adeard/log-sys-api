package utils

import (
	"log"
	"os"
	"time"
)

func LogInit(message string) {

	now := time.Now().Format("2006-01-02")

	f, err := os.OpenFile("logs/"+now+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(message)
}
