package logger

import (
	"io"
	"log"
	"os"
)

var (
	//LogWarn ...
	LogWarn *log.Logger
	//LogInfo ...
	LogInfo *log.Logger
	//LogError ...
	LogError *log.Logger
	//LogFile ...
	LogFile string = "/tmp/logs.txt"
)

//InitLogger ...
func InitLogger() (LogInfo, LogWarn, LogError *log.Logger) {
	file, err := os.OpenFile(LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)

	log.SetOutput(multiWriter)
	LogInfo = log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	LogWarn = log.New(multiWriter, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	LogError = log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return LogInfo, LogWarn, LogError
}
