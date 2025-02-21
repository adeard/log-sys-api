package utils

import (
	"log"
	"log-sys-api/domain"
	"net/url"
	"os"
	"strings"
	"time"
)

func LogInit(host string, message string) {

	now := time.Now().Format("2006-01-02")

	f, err := os.OpenFile("./logs/"+host+"-"+now+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(message)
}

func LogInit2(host string, message domain.LogRequest) {
	appName := "unnamed"
	now := time.Now().Format("2006-01-02")

	if !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
		host = "http://" + host // Add default protocol
	}

	parsedURL, err := url.Parse(host)
	if err != nil {
		LogInit("error", err.Error())
	}

	split := strings.Split(parsedURL.Path, "/")

	if len(split) > 1 {
		appName = split[1]
	}

	hostRes := strings.ReplaceAll(parsedURL.Host, ":", "@")
	filename := hostRes + "_" + appName + "_" + now + ".log"

	f, err := os.OpenFile("./logs/"+filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(message)
}
