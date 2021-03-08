package main

import (
	"fmt"
	"go-utils/cerrors"
	"go-utils/couts"
	"go-utils/logger"
	"log"
	"os"
)

var (
	logInfo, logWarn, logError *log.Logger
)

func init() {
	logInfo, logWarn, logError = logger.InitLogger()
}

var (
	err error
)

// Sample snippet to use Custom Error
func openFile() (Output string, out interface{}) {
	var file *os.File
	if file, err = os.Open("test.txt"); err != nil {
		generalError := &cerrors.GeneralError{Code: "500", Message: "Unable to open file", Err: err}
		return "Failed", generalError
	}

	generalOutput := &couts.GeneralOutput{Code: "200", Message: "File opened successfully", Out: file.Name()}

	return "Sucess", generalOutput
}

func main() {
	output, out := openFile()
	// Check error type
	if out, ok := out.(*cerrors.GeneralError); ok {
		finalError := output + ": GeneralError - " + out.Error()
		logError.Println(finalError)
	}

	if out, ok := out.(*couts.GeneralOutput); ok {
		finalOutput := output + ": GeneralOutput - " + fmt.Sprintf("%v", out)
		logInfo.Println(finalOutput)
	}
}
