package main

import (
	"fmt"
	"go-utils/cerrors"
	"os"
)

// Sample snippet to use Custom Error
func openFile() (output string, err error) {
	var file *os.File
	if file, err = os.Open("test.txt"); err != nil {
		return "Failed", &cerrors.GeneralError{Code: "123", Message: "Unable to open file", Err: err}
	}
	fmt.Println(file.Name())
	return "Sucess", nil
}

func main() {
	a, err := openFile()
	// Check error type
	if _, ok := err.(*cerrors.GeneralError); ok {
		fmt.Println("GeneralError - ", err)
	}
	fmt.Println(a)
	fmt.Println(err)
}
